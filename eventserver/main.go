package main

import (
	"net/http"
	"log"
	"runtime"

	"code.google.com/p/go.net/websocket"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

type command int

const (
	COMMAND_PLAY command = iota
	COMMAND_PAUSE
	COMMAND_VOLUME_UP
	COMMAND_VOLUME_DOWN
)

// This seems pretty weird, aw well
var socketChannels []chan command

func YtfdServer(ws *websocket.Conn) {
	myChan := make(chan command)
	socketChannels = append(socketChannels, myChan)
	for {
		cmd := <-myChan
		cmdData := map[string]command{
			"command": cmd,
		}
		websocket.JSON.Send(ws, &cmdData)
	}
	myChan = nil
}

func main() {
	runtime.GOMAXPROCS(2)

	socketChannels = make([]chan command, 0)
	go func() {
		X, _ := xgbutil.NewConn()
		keybind.Initialize(X)		

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending PLAY...")
			for i := range socketChannels {
				if socketChannels[i] != nil {
					socketChannels[i] <- COMMAND_PLAY
				}
			}
		}).Connect(X, X.RootWin(), "Mod1-Insert", true)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending PAUSE...")
			for i := range socketChannels {
				if socketChannels[i] != nil {
					socketChannels[i] <- COMMAND_PAUSE
				}
			}
		}).Connect(X, X.RootWin(), "Mod1-Home", true)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending VOLUME UP...")
			for i := range socketChannels {
				if socketChannels[i] != nil {
					socketChannels[i] <- COMMAND_VOLUME_UP
				}
			}
		}).Connect(X, X.RootWin(), "Mod1-Prior", true)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending VOLUME DOWN...")
			for i := range socketChannels {
				if socketChannels[i] != nil {
					socketChannels[i] <- COMMAND_VOLUME_DOWN
				}
			}
		}).Connect(X, X.RootWin(), "Mod1-Next", true)

		xevent.Main(X)


		/*
		X, _ := xgb.NewConn()
		screen := xproto.Setup(X).DefaultScreen(X)
		wid := screen.Root
		xproto.GrabKey(X, false, wid, xproto.ModMaskControl|xproto.ModMask1|xproto.ModMask2, 118, xproto.GrabModeAsync, xproto.GrabModeAsync)
		*/
	}()

	http.Handle("/ws", websocket.Handler(YtfdServer))
	err := http.ListenAndServe(":42050", nil)
	if err != nil {
		log.Println(err)
	}
}

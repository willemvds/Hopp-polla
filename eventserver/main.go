package main

import (
	"net/http"
	"log"
	"runtime"

	"code.google.com/p/go.net/websocket"
)

type command int

const (
	COMMAND_PLAY command = iota
	COMMAND_PAUSE
	COMMAND_VOLUME_UP
	COMMAND_VOLUME_DOWN
)

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
	commandChannel := make(chan command)

	go func() {
		for {
			cmd := <- commandChannel
			for i := range socketChannels {
				socketChannels[i] <- cmd
			}
		}
	}()
	StartListener(commandChannel)

	http.Handle("/ws", websocket.Handler(YtfdServer))
	err := http.ListenAndServe(":42050", nil)
	if err != nil {
		log.Println(err)
	}
}

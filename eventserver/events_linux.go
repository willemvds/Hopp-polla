package main

import (
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

func StartListener(cmdChan chan command) {
	go func() {
		X, _ := xgbutil.NewConn()
		keybind.Initialize(X)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending PLAY...")
			cmdChan <- COMMAND_PLAY
		}).Connect(X, X.RootWin(), "Mod1-Insert", true)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending PAUSE...")
			cmdChan <- COMMAND_PAUSE
		}).Connect(X, X.RootWin(), "Mod1-Home", true)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending VOLUME UP...")
			cmdChan <- COMMAND_VOLUME_UP
		}).Connect(X, X.RootWin(), "Mod1-Prior", true)

		keybind.KeyPressFun(func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("Sending VOLUME DOWN...")
			cmdChan <- COMMAND_VOLUME_DOWN
		}).Connect(X, X.RootWin(), "Mod1-Next", true)

		xevent.Main(X)
	}()
}

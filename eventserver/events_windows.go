package main

import (
	"log"

	"github.com/willemvds/w32"
)

// 0x2D // Insert
// 0x24 // Home
// 0x21 // PGUP
// 0x22 // PGDOWN

func StartListener(cmdChan chan command) {
	go func() {
		var msg w32.MSG
		w32.RegisterHotKey(0, 42030, 0x0001|0x0002, 0x2D)
		w32.RegisterHotKey(0, 42031, 0x0001|0x0002, 0x24)
		w32.RegisterHotKey(0, 42032, 0x0001|0x0002, 0x21)
		w32.RegisterHotKey(0, 42033, 0x0001|0x0002, 0x22)
		for {
			ok := w32.GetMessage(&msg, 0, 0, 0)
			if ok != 1 {
				continue
			}
			switch msg.WParam {
			case 42030:
				log.Println("Sending PLAY...")
				cmdChan <- COMMAND_PLAY
			case 42031:
				log.Println("Sending PAUSE...")
				cmdChan <- COMMAND_PAUSE
			case 42032:
				log.Println("Sending VOLUME UP...")
				cmdChan <- COMMAND_VOLUME_UP
			case 42033:
				log.Println("Sending VOLUME DOWN...")
				cmdChan <- COMMAND_VOLUME_DOWN
			}
		}
	}()
}

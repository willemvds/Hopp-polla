package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -leventloop -framework Carbon
#include "eventloop.h"

CGEventRef MyEventTapCallBack_cgo(CGEventTapProxy proxy, CGEventType type, CGEventRef ref, void *udref);
*/
import "C"

import (
	"log"
	"unsafe"
)

const maskCtrl uint64 = 262401
const maskShift uint64 = 131330

var gCmdChan chan command

//export KeyPressedCallback
func KeyPressedCallback(keyCode uint32, eventFlags uint64) {
	if keyCode == 40 && ((eventFlags & maskCtrl) == maskCtrl) && ((eventFlags & maskShift) == maskShift) {
		log.Println("Sending PLAY...")
		gCmdChan <- COMMAND_PLAY
	} else if keyCode == 37 && ((eventFlags & maskCtrl) == maskCtrl) && ((eventFlags & maskShift) == maskShift) {
		log.Println("Sending PAUSE...")
		gCmdChan <- COMMAND_PAUSE
	}
}

func StartListener(cmdChan chan command) {
	gCmdChan = cmdChan
	go func() {
		C.startEventLoop((C.MyEventTapCallBack_fcn)(unsafe.Pointer(C.MyEventTapCallBack_cgo)))
	}()
}

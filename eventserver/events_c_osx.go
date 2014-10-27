package main

/*
#include <stdio.h>
#include <ApplicationServices/ApplicationServices.h>

CGEventRef MyEventTapCallBack_cgo(CGEventTapProxy proxy, CGEventType type, CGEventRef ref, void *udref)
{
	CGKeyCode keyCode = (CGKeyCode) CGEventGetIntegerValueField(ref, kCGKeyboardEventKeycode);
	CGEventFlags eventFlags = CGEventGetFlags(ref);
	KeyPressedCallback(keyCode, eventFlags);
	return NULL;
}
*/
import "C"

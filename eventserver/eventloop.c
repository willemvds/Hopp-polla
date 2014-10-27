#include <stdio.h>
#include "eventloop.h"

void startEventLoop(MyEventTapCallBack_fcn callback)
{
    int userData = 42;
    CGEventMask mask = CGEventMaskBit(kCGEventKeyUp);
    CFMachPortRef port = CGEventTapCreate(kCGHIDEventTap, kCGHeadInsertEventTap, kCGEventTapOptionListenOnly, mask, callback, &userData);

    if (port != NULL) {
	    CFRunLoopSourceRef loopref = CFMachPortCreateRunLoopSource(NULL, port, 0);
	    CFRunLoopAddSource(CFRunLoopGetCurrent(), loopref, kCFRunLoopDefaultMode);
	    CFRunLoopRun();
    }
}

#ifndef EVENTLOOP_H
#define EVENTLOOP_H

#include <ApplicationServices/ApplicationServices.h>

typedef CGEventRef(*MyEventTapCallBack_fcn)(CGEventTapProxy proxy, CGEventType type, CGEventRef ref, void *udref);
void startEventLoop(MyEventTapCallBack_fcn);

#endif //EVENTLOOP_H

// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Event interface {
	GetType() EventType
}

func (e *WindowEvent) GetType() EventType {
	return e.Type
}

func (e *KeyboardEvent) GetType() EventType {
	return e.Type
}

func (e *TextEditingEvent) GetType() EventType {
	return e.Type
}

func (e *TextInputEvent) GetType() EventType {
	return e.Type
}

func (e *MouseMotionEvent) GetType() EventType {
	return e.Type
}

func (e *MouseButtonEvent) GetType() EventType {
	return e.Type
}

func (e *MouseWheelEvent) GetType() EventType {
	return e.Type
}

func (e *JoyAxisEvent) GetType() EventType {
	return e.Type
}

func (e *JoyBallEvent) GetType() EventType {
	return e.Type
}

func (e *JoyHatEvent) GetType() EventType {
	return e.Type
}

func (e *JoyButtonEvent) GetType() EventType {
	return e.Type
}

func (e *JoyDeviceEvent) GetType() EventType {
	return e.Type
}

func (e *ControllerAxisEvent) GetType() EventType {
	return e.Type
}

func (e *ControllerButtonEvent) GetType() EventType {
	return e.Type
}

func (e *ControllerDeviceEvent) GetType() EventType {
	return e.Type
}

func (e *TouchFingerEvent) GetType() EventType {
	return e.Type
}

func (e *MultiGestureEvent) GetType() EventType {
	return e.Type
}

func (e *DollarGestureEvent) GetType() EventType {
	return e.Type
}

func (e *DropEvent) GetType() EventType {
	return e.Type
}

// File returns the file name.
func (e *DropEvent) File() string {
	return C.GoString(e.file)
}

// SetFile sets the file name.  SetFile calls FreeFile before setting the new
// file name. 
func (e *DropEvent) SetFile(file string) {
	e.FreeFile()
	cFile := C.CString(file)
	e.file = cFile
}

// FreeFile frees the file name.  This must be called for each DropEvent.
func (e *DropEvent) FreeFile() {
	C.SDL_free(unsafe.Pointer(e.file))
	e.file = nil
}

func (e *QuitEvent) GetType() EventType {
	return e.Type
}

func (e *UserEvent) GetType() EventType {
	return e.Type
}

func (e *SysWMEvent) GetType() EventType {
	return e.Type
}

func (e *EventUnion) GetType() EventType {
	return e.Type
}

// Convert converts event form an EventUnion into one of the *Event types.
// Use it to gain access to all the fields of event.
//
// If Convert can not figure out what type to convert event into then Convert
// will return event without converting it.  This should only happen if this
// package has fallen out of sync with SDL2 and should be reported if
// possible.
func (event *EventUnion) Convert() Event {
	if !validEventType(event.Type) {
		return event
	}

	var e Event
	switch event.Type {
	case QUIT:
		e = (*QuitEvent)(unsafe.Pointer(event))
	case WINDOWEVENT:
		e = (*WindowEvent)(unsafe.Pointer(event))
	case SYSWMEVENT:
		e = (*SysWMEvent)(unsafe.Pointer(event))
	case KEYDOWN, KEYUP:
		e = (*KeyboardEvent)(unsafe.Pointer(event))
	case TEXTEDITING:
		e = (*TextEditingEvent)(unsafe.Pointer(event))
	case TEXTINPUT:
		e = (*TextInputEvent)(unsafe.Pointer(event))
	case MOUSEMOTION:
		e = (*MouseMotionEvent)(unsafe.Pointer(event))
	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		e = (*MouseButtonEvent)(unsafe.Pointer(event))
	case MOUSEWHEEL:
		e = (*MouseWheelEvent)(unsafe.Pointer(event))
	case JOYAXISMOTION:
		e = (*JoyAxisEvent)(unsafe.Pointer(event))
	case JOYBALLMOTION:
		e = (*JoyBallEvent)(unsafe.Pointer(event))
	case JOYHATMOTION:
		e = (*JoyHatEvent)(unsafe.Pointer(event))
	case JOYBUTTONDOWN, JOYBUTTONUP:
		e = (*JoyButtonEvent)(unsafe.Pointer(event))
	case JOYDEVICEADDED, JOYDEVICEREMOVED:
		e = (*JoyDeviceEvent)(unsafe.Pointer(event))
	case CONTROLLERAXISMOTION:
		e = (*ControllerAxisEvent)(unsafe.Pointer(event))
	case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		e = (*ControllerButtonEvent)(unsafe.Pointer(event))
	case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		e = (*ControllerDeviceEvent)(unsafe.Pointer(event))
	case FINGERMOTION, FINGERDOWN, FINGERUP:
		e = (*TouchFingerEvent)(unsafe.Pointer(event))
	case MULTIGESTURE:
		e = (*MultiGestureEvent)(unsafe.Pointer(event))
	case DOLLARGESTURE:
		e = (*DollarGestureEvent)(unsafe.Pointer(event))
	case DROPFILE:
		e = (*DropEvent)(unsafe.Pointer(event))
	default:
		e = (*UserEvent)(unsafe.Pointer(event))
	}

	return e
}

func CopyEventToEventUnion(ev Event, evu *EventUnion) error {
	if !validEventType(ev.GetType()) {
		return fmt.Errorf("Invalid EventType: %s", ev.GetType())
	}

	dst := []byte{}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
	sh.Data = uintptr(unsafe.Pointer(evu))
	sh.Len = int(unsafe.Sizeof(*evu))
	sh.Cap = sh.Len

	src := []byte{}
	sh = (*reflect.SliceHeader)(unsafe.Pointer(&src))

	switch t := ev.(type) {
	case *QuitEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *WindowEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *SysWMEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *KeyboardEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *TextEditingEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *TextInputEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *MouseMotionEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *MouseButtonEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *MouseWheelEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *JoyAxisEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *JoyBallEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *JoyHatEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *JoyButtonEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *JoyDeviceEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *ControllerAxisEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *ControllerButtonEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *ControllerDeviceEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *TouchFingerEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *MultiGestureEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *DollarGestureEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *DropEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	case *UserEvent:
		sh.Data = uintptr(unsafe.Pointer(t))
		sh.Len = int(unsafe.Sizeof(*t))
	default:
		return fmt.Errorf("Unknown type: %T", t)
	}
	sh.Cap = sh.Len

	copy(dst, src)
	return nil
}

const (
	RELEASED = C.SDL_RELEASED
	PRESSED  = C.SDL_PRESSED
)

type EventType uint32

const (
	FIRSTEVENT EventType = C.SDL_FIRSTEVENT

	QUIT EventType = C.SDL_QUIT

	WINDOWEVENT EventType = C.SDL_WINDOWEVENT
	SYSWMEVENT  EventType = C.SDL_SYSWMEVENT

	KEYDOWN     EventType = C.SDL_KEYDOWN
	KEYUP       EventType = C.SDL_KEYUP
	TEXTEDITING EventType = C.SDL_TEXTEDITING
	TEXTINPUT   EventType = C.SDL_TEXTINPUT

	MOUSEMOTION     EventType = C.SDL_MOUSEMOTION
	MOUSEBUTTONDOWN EventType = C.SDL_MOUSEBUTTONDOWN
	MOUSEBUTTONUP   EventType = C.SDL_MOUSEBUTTONUP
	MOUSEWHEEL      EventType = C.SDL_MOUSEWHEEL

	JOYAXISMOTION    EventType = C.SDL_JOYAXISMOTION
	JOYBALLMOTION    EventType = C.SDL_JOYBALLMOTION
	JOYHATMOTION     EventType = C.SDL_JOYHATMOTION
	JOYBUTTONDOWN    EventType = C.SDL_JOYBUTTONDOWN
	JOYBUTTONUP      EventType = C.SDL_JOYBUTTONUP
	JOYDEVICEADDED   EventType = C.SDL_JOYDEVICEADDED
	JOYDEVICEREMOVED EventType = C.SDL_JOYDEVICEREMOVED

	CONTROLLERAXISMOTION     EventType = C.SDL_CONTROLLERAXISMOTION
	CONTROLLERBUTTONDOWN     EventType = C.SDL_CONTROLLERBUTTONDOWN
	CONTROLLERBUTTONUP       EventType = C.SDL_CONTROLLERBUTTONUP
	CONTROLLERDEVICEADDED    EventType = C.SDL_CONTROLLERDEVICEADDED
	CONTROLLERDEVICEREMOVED  EventType = C.SDL_CONTROLLERDEVICEREMOVED
	CONTROLLERDEVICEREMAPPED EventType = C.SDL_CONTROLLERDEVICEREMAPPED

	FINGERDOWN   EventType = C.SDL_FINGERDOWN
	FINGERUP     EventType = C.SDL_FINGERUP
	FINGERMOTION EventType = C.SDL_FINGERMOTION

	DOLLARGESTURE EventType = C.SDL_DOLLARGESTURE
	DOLLARRECORD  EventType = C.SDL_DOLLARRECORD
	MULTIGESTURE  EventType = C.SDL_MULTIGESTURE

	CLIPBOARDUPDATE EventType = C.SDL_CLIPBOARDUPDATE

	DROPFILE EventType = C.SDL_DROPFILE

	USEREVENT EventType = C.SDL_USEREVENT

	LASTEVENT EventType = C.SDL_LASTEVENT
)

var eventTypeStrings = map[EventType]string{
	FIRSTEVENT:               "FIRSTEVENT",
	QUIT:                     "QUIT",
	WINDOWEVENT:              "WINDOWEVENT",
	SYSWMEVENT:               "SYSWMEVENT",
	KEYDOWN:                  "KEYDOWN",
	KEYUP:                    "KEYUP",
	TEXTEDITING:              "TEXTEDITING",
	TEXTINPUT:                "TEXTINPUT",
	MOUSEMOTION:              "MOUSEMOTION",
	MOUSEBUTTONDOWN:          "MOUSEBUTTONDOWN",
	MOUSEBUTTONUP:            "MOUSEBUTTONUP",
	MOUSEWHEEL:               "MOUSEWHEEL",
	JOYAXISMOTION:            "JOYAXISMOTION",
	JOYBALLMOTION:            "JOYBALLMOTION",
	JOYHATMOTION:             "JOYHATMOTION",
	JOYBUTTONDOWN:            "JOYBUTTONDOWN",
	JOYBUTTONUP:              "JOYBUTTONUP",
	JOYDEVICEADDED:           "JOYDEVICEADDED",
	JOYDEVICEREMOVED:         "JOYDEVICEREMOVED",
	CONTROLLERAXISMOTION:     "CONTROLLERAXISMOTION",
	CONTROLLERBUTTONDOWN:     "CONTROLLERBUTTONDOWN",
	CONTROLLERBUTTONUP:       "CONTROLLERBUTTONUP",
	CONTROLLERDEVICEADDED:    "CONTROLLERDEVICEADDED",
	CONTROLLERDEVICEREMOVED:  "CONTROLLERDEVICEREMOVED",
	CONTROLLERDEVICEREMAPPED: "CONTROLLERDEVICEREMAPPED",
	FINGERDOWN:               "FINGERDOWN",
	FINGERUP:                 "FINGERUP",
	FINGERMOTION:             "FINGERMOTION",
	DOLLARGESTURE:            "DOLLARGESTURE",
	DOLLARRECORD:             "DOLLARRECORD",
	MULTIGESTURE:             "MULTIGESTURE",
	CLIPBOARDUPDATE:          "CLIPBOARDUPDATE",
	DROPFILE:                 "DROPFILE",
	USEREVENT:                "USEREVENT",
	LASTEVENT:                "LASTEVENT",
}

func (et EventType) String() string {
	if et >= USEREVENT && et < LASTEVENT {
		return "USEREVENT"
	}

	str, ok := eventTypeStrings[et]
	if !ok {
		return fmt.Sprintf("Unknown (%d)", et)
	}
	return str
}

func validEventType(t EventType) bool {
	switch {
	case t == QUIT:
		return true
	case t == WINDOWEVENT:
		return true
	case t == SYSWMEVENT:
		return true
	case t == KEYDOWN:
		return true
	case t == KEYUP:
		return true
	case t == TEXTEDITING:
		return true
	case t == TEXTINPUT:
		return true
	case t == MOUSEMOTION:
		return true
	case t == MOUSEBUTTONDOWN:
		return true
	case t == MOUSEBUTTONUP:
		return true
	case t == MOUSEWHEEL:
		return true
	case t == JOYAXISMOTION:
		return true
	case t == JOYBALLMOTION:
		return true
	case t == JOYHATMOTION:
		return true
	case t == JOYBUTTONDOWN:
		return true
	case t == JOYBUTTONUP:
		return true
	case t == JOYDEVICEADDED:
		return true
	case t == JOYDEVICEREMOVED:
		return true
	case t == CONTROLLERAXISMOTION:
		return true
	case t == CONTROLLERBUTTONDOWN:
		return true
	case t == CONTROLLERBUTTONUP:
		return true
	case t == CONTROLLERDEVICEADDED:
		return true
	case t == CONTROLLERDEVICEREMOVED:
		return true
	case t == CONTROLLERDEVICEREMAPPED:
		return true
	case t == FINGERDOWN:
		return true
	case t == FINGERUP:
		return true
	case t == FINGERMOTION:
		return true
	case t == DOLLARGESTURE:
		return true
	case t == DOLLARRECORD:
		return false
	case t == MULTIGESTURE:
		return true
	case t == CLIPBOARDUPDATE:
		return false
	case t == DROPFILE:
		return true
	case t >= USEREVENT && t <= LASTEVENT:
		return true
	}
	return false
}

// PumpEvents pumps the event loop, gathering events from the input devices.
// It updates the event queue and internal input device state.
//
// On windows this will need to be run in each OS thread that creates a
// windows.  For this reason you will likely want to call
// runtime.LockOSThread() before creating a window.
func PumpEvents() {
	C.SDL_PumpEvents()
}

type EventAction uint32

const (
	ADDEVENT  EventAction = C.SDL_ADDEVENT
	PEEKEVENT EventAction = C.SDL_PEEKEVENT
	GETEVENT  EventAction = C.SDL_GETEVENT
)

// PeepEvents checks the event queue for messages and optionally returns them.
// The length of the EventUnion slice determines how many events are added or
// possibly returned from the event queue.
//
// If action is ADDEVENT, events will be added to the back of the event queue.
//
// If action is PEEKEVENT, events at the front of the event queue, within the
// specified minimum and maximum type, will be returned and will not be removed
// from the queue.
//
// If action is GETEVENT, events at the front of the event queue, within the
// specified minimum and maximum type, will be returned and will be removed
// from the queue.
func PeepEvents(events []EventUnion, action EventAction, minType, maxType EventType) (int, error) {
	var ptr *C.SDL_Event
	if len(events) > 0 {
		ptr = (*C.SDL_Event)(unsafe.Pointer(&events[0]))
	}

	r := int(C.SDL_PeepEvents(ptr, C.int(len(events)),
		C.SDL_eventaction(action),
		C.Uint32(minType), C.Uint32(maxType)))

	if r == -1 {
		return r, sdlError(r)
	}

	return r, nil
}

// HasEvent checks to see if any events of type typ are in the event queue.
func HasEvent(typ EventType) bool {
	if b := C.SDL_HasEvent(C.Uint32(typ)); b == C.SDL_TRUE {
		return true
	}
	return false
}

// HasEvents checks if events between minType and maxType are in the event
// queue.
func HasEvents(minType, maxType EventType) bool {
	if b := C.SDL_HasEvents(C.Uint32(minType), C.Uint32(maxType)); b == C.SDL_TRUE {
		return true
	}
	return false
}

// FlushEvent clears events of type typ from the event queue.
func FlushEvent(typ EventType) {
	C.SDL_FlushEvent(C.Uint32(typ))
}

// FlushEvents clears events between minType and maxType from the event queue.
func FlustEvents(minType, maxType EventType) {
	C.SDL_FlushEvents(C.Uint32(minType), C.Uint32(maxType))
}

// PollEvent polls for currently pending events.  It returns 1 if there are
// any pending events, or 0 if there are none available.  If event is nil
// the next event will not be removed from the queue.
func PollEvent(event *EventUnion) int {
	return int(C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(event))))
}

// WaitEvent waits indefinitely for the next available event.  If event is nil
// the next event will not be removed from the queue.
func WaitEvent(event *EventUnion) error {
	r := int(C.SDL_WaitEvent((*C.SDL_Event)(unsafe.Pointer(event))))

	if r == 0 {
		return sdlError(r)
	}
	return nil
}

// WaitEventTimeout waits until the specified timeout (in milliseconds) for
// the next available event.  If event is nil the next event will not be
// removed from the queue.
func WaitEventTimeout(event *EventUnion, timeout int) error {
	r := int(C.SDL_WaitEventTimeout((*C.SDL_Event)(unsafe.Pointer(event)),
		C.int(timeout)))

	if r == 0 {
		return sdlError(r)
	}
	return nil
}

// PushEvent adds an event to the event queue. It returns 1 on succes or 0 if
// the event was filtered.
func PushEvent(event *EventUnion) (int, error) {
	r := int(C.SDL_PushEvent((*C.SDL_Event)(unsafe.Pointer(event))))

	if r == -1 {
		return r, sdlError(r)
	}
	return r, nil
}

/*
Skipping
SDL_SetEventFilter()
SDL_GetEventFilter()
The comment for SetEventFilter says the function can run in another thread,
as far as I know this would cause the program to crash as go would not know
about the thread.
SDL_AddEventWatch()
SDL_DelEventWatch()
SDL_FilterEvents()
I don't know if these can run in a different thread like SetEventFilter can.
*/

const (
	QUERY   = C.SDL_QUERY
	IGNORE  = C.SDL_IGNORE
	DISABLE = C.SDL_DISABLE
	ENABLE  = C.SDL_ENABLE
)

// EventState allows you to set the state of processing certain events.  If
// state is IGNORE, that event will be automatically dropped from the queue
// and will not even be filtered.  If state is ENABLE, that event will be
// processed normally.  If state is QUERY EventState will return the current
// processing state of the specified event.
func EventState(typ EventType, state int) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(typ), C.int(state)))
}

// GetEventState returns current processing state for events of type typ.
func GetEventState(typ EventType) uint8 {
	return EventState(typ, QUERY)
}

// RegisterEvents allocates a set of user-defined events, and returns the
// beginning event number for that set of events.  It there aren't enough
// user-defined events left, it returns (Uint32)-1.
func RegisterEvents(numevents int) EventType {
	return EventType(C.SDL_RegisterEvents(C.int(numevents)))
}

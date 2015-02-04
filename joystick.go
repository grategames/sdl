// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

// Joystick is used to identify an SDL joystick.
type Joystick struct {
	ptr *C.SDL_Joystick
}

type JoystickID int32

// NumJoysticks counts the number of joysticks attached to the system right
// now.
func NumJoysticks() int {
	return int(C.SDL_NumJoysticks())
}

// JoystickNameForIndex gets the implementation dependent name of a joystick.
// This can be called before any joysticks are opened.  If no name can be
// found, this function returns and empty string.
func JoystickNameForIndex(device_index int) string {
	return C.GoString(C.SDL_JoystickNameForIndex(C.int(device_index)))
}

// JoystickOpen opens a joystick for use.  device_index refers to the N'th
// joystick on the system.  device_index is the value which will identify this
// joystick in future joystick events.
func JoystickOpen(device_index int) (Joystick, error) {
	ptr := C.SDL_JoystickOpen(C.int(device_index))
	if ptr == nil {
		return Joystick{}, sdlError(0)
	}
	return Joystick{ptr}, nil
}

// Name gets the name for j.  If no name can be found, Name returns an empty
// string.
func (j Joystick) Name() string {
	return C.GoString(C.SDL_JoystickName(j.ptr))
}

// JoystickGetDeviceGUID returns the GUID for the joystick at device_index.
func JoystickGetDeviceGUID(device_index int) JoystickGUID {
	r := C.SDL_JoystickGetDeviceGUID(C.int(device_index))
	return *(*JoystickGUID)(unsafe.Pointer(&r))
}

// GetGUID returns the GUID for j.
func (j Joystick) GetGUID() JoystickGUID {
	r := C.SDL_JoystickGetGUID(j.ptr)
	return *(*JoystickGUID)(unsafe.Pointer(&r))
}

// JoystickGetGUIDString returns a string representation for guid.
func JoystickGetGUIDString(guid JoystickGUID) string {
	// SDL_joystick.h says the string must be at least 33 bytes.
	const size = 33

	buf := make([]byte, size)
	C.SDL_JoystickGetGUIDString(
		*(*C.SDL_JoystickGUID)(unsafe.Pointer(&guid)),
		(*C.char)(unsafe.Pointer(&buf[0])), size)

	return string(buf)
}

// JoystickGetGUIDFromString converts a string into a joystick formatted guid.
func JoystickGetGUIDFromString(pchGUID string) JoystickGUID {
	cstr := C.CString(pchGUID)
	defer C.free(unsafe.Pointer(cstr))
	r := C.SDL_JoystickGetGUIDFromString(cstr)
	return *(*JoystickGUID)(unsafe.Pointer(&r))
}

// GetAttached returns true if the joystick has been opened and is currently
// connected, or false otherwise.
func (j Joystick) GetAttached() bool {
	r := C.SDL_JoystickGetAttached(j.ptr)
	if r == C.SDL_TRUE {
		return true
	}
	return false
}

// InstanceID gets the device index of j.
func (j Joystick) InstanceID() JoystickID {
	return JoystickID(C.SDL_JoystickInstanceID(j.ptr))
}

// NumAxes gets the number of general axis controls on j.
func (j Joystick) NumAxes() int {
	return int(C.SDL_JoystickNumAxes(j.ptr))
}

// NumBalls gets the number of trackballs on j.
//
// Joystick trackballs have only relative motion events associated with them
// and their state cannot be polled.
func (j Joystick) NumBalls() int {
	return int(C.SDL_JoystickNumBalls(j.ptr))
}

// NumHats gets the number of POV hats on j.
func (j Joystick) NumHats() int {
	return int(C.SDL_JoystickNumHats(j.ptr))
}

// NumButtons gets the number of buttons on j.
func (j Joystick) NumButtons() int {
	return int(C.SDL_JoystickNumButtons(j.ptr))
}

// JoystickUpdate updates the current state of the open joysticks.
//
// This is called automatically by the event loop if any joystick events are
// enabled.
func JoystickUpdate() {
	C.SDL_JoystickUpdate()
}

// JoystickEventState enables/disables joystick event polling.
//
// If joystick events are disabled, you must call JoystickUpdate yourself and
// check the state of the joystick when you want joystick information.
//
// The state can be one of QUERY, ENABLE, or IGNORE.
func JoystickEventState(state int) int {
	return int(C.SDL_JoystickEventState(C.int(state)))
}

// GetAxis gets the current state of an axis control on j.
//
// The state is a value ranging from -32768 to 32767.
//
// The axis indices start at index 0.
func (j Joystick) GetAxis(axis int) int16 {
	return int16(C.SDL_JoystickGetAxis(j.ptr, C.int(axis)))
}

type HatPosition uint8

const (
	HAT_CENTERED  HatPosition = C.SDL_HAT_CENTERED
	HAT_UP        HatPosition = C.SDL_HAT_UP
	HAT_RIGHT     HatPosition = C.SDL_HAT_RIGHT
	HAT_DOWN      HatPosition = C.SDL_HAT_DOWN
	HAT_LEFT      HatPosition = C.SDL_HAT_LEFT
	HAT_RIGHTUP   HatPosition = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN HatPosition = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    HatPosition = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  HatPosition = C.SDL_HAT_LEFTDOWN
)

// GetHat gets the current state of a POV hat on j.
//
// The hat indices start at index 0.
func (j Joystick) GetHat(hat int) HatPosition {
	return HatPosition(C.SDL_JoystickGetHat(j.ptr, C.int(hat)))
}

// GetBall gets the ball axis change since the last poll.
//
// The ball indices start at index 0.
func (j Joystick) GetBall(ball int) (dx, dy, err error) {
	r := C.SDL_JoystickGetBall(j.ptr, C.int(ball),
		(*C.int)(unsafe.Pointer(&dx)), (*C.int)(unsafe.Pointer(&dy)))
	if r != 0 {
		err = sdlError(int(r))
	}
	return
}

// GetButton gets the current state of a button on j.
//
// The button indices start at index 0.
func (j Joystick) GetButton(button int) uint8 {
	return uint8(C.SDL_JoystickGetButton(j.ptr, C.int(button)))
}

// Close closes j.
func (j Joystick) Close() {
	C.SDL_JoystickClose(j.ptr)
}

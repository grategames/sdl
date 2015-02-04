// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

// #include "SDL.h"
// #include "gamecontroller.h"
import "C"
import "fmt"
import "unsafe"

type GameController struct {
	ptr *C.SDL_GameController
}

type ControllerBindtype int32

const (
	CONTROLLER_BINDTYPE_NONE   ControllerBindtype = C.SDL_CONTROLLER_BINDTYPE_NONE
	CONTROLLER_BINDTYPE_BUTTON ControllerBindtype = C.SDL_CONTROLLER_BINDTYPE_BUTTON
	CONTROLLER_BINDTYPE_AXIS   ControllerBindtype = C.SDL_CONTROLLER_BINDTYPE_AXIS
	CONTROLLER_BINDTYPE_HAT    ControllerBindtype = C.SDL_CONTROLLER_BINDTYPE_HAT
)

type GameControllerButtonBind struct {
	BindType ControllerBindtype
	Button   int32
}

type GameControllerAxisBind struct {
	BindType ControllerBindtype
	Axis     int32
}

type GameControllerHatBind struct {
	BindType ControllerBindtype
	Hat      int32
	HatMask  int32
}

func buildBind(typ ControllerBindtype, v1, v2 int32) interface{} {
	switch typ {
	case C.SDL_CONTROLLER_BINDTYPE_BUTTON:
		return GameControllerButtonBind{typ, v1}
	case C.SDL_CONTROLLER_BINDTYPE_AXIS:
		return GameControllerAxisBind{typ, v1}
	case C.SDL_CONTROLLER_BINDTYPE_HAT:
		return GameControllerHatBind{typ, v1, v2}
	default:
		str := fmt.Sprintf("Unhandled ControllerBindtype: %v", typ)
		panic(str)
	}
	panic("Unreachable")
}

// GameControllerAddMapping adds or updates an existing mapping configuration.
func GameControllerAddMapping(mappingString string) error {
	cstr := C.CString(mappingString)
	defer C.free(unsafe.Pointer(cstr))

	r := C.SDL_GameControllerAddMapping(cstr)
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// GameControllerMapping returns a mapping string for guid.
func (guid JoystickGUID) GameControllerMapping() (string, error) {
	cstr := C.SDL_GameControllerMappingForGUID(
		*(*C.SDL_JoystickGUID)(unsafe.Pointer(&guid)))
	if cstr == nil {
		return "", sdlError(0)
	}

	str := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))

	return str, nil
}

// Mapping returns a mapping string for c.
func (c GameController) Mapping() (string, error) {
	cstr := C.SDL_GameControllerMapping(c.ptr)
	if cstr == nil {
		return "", sdlError(0)
	}

	str := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))

	return str, nil
}

// IsGameController checks if the joystick on this index is supported by the
// game controller interface.
func IsGameController(joystick_index int) bool {
	r := C.SDL_IsGameController(C.int(joystick_index))
	if r == 1 {
		return true
	}
	return false
}

// GameControllerNameForIndex gets the implementation dependent name of a game
// controller.  This can be called before any controllers are opened.  If no
// name can be found, this function returns an empty string.
func GameControllerNameForIndex(joystick_index int) string {
	str := C.SDL_GameControllerNameForIndex(C.int(joystick_index))
	return C.GoString(str)
}

// GameControllerOpen opens a game controller for use.  The index passed as an
// argument refres to the N'th game controller on the system.  This index is
// the value which will identify this controller in future controller events.
func GameControllerOpen(joystick_index int) (GameController, error) {
	ptr := C.SDL_GameControllerOpen(C.int(joystick_index))
	if ptr == nil {
		return GameController{}, sdlError(0)
	}
	return GameController{ptr}, nil
}

// Name returns the name for c.
func (c GameController) Name() string {
	return C.GoString(C.SDL_GameControllerName(c.ptr))
}

// GetAttached return true if the controller has been opened and is currently
// connected, or false otherwise.
func (c GameController) GetAttached() bool {
	r := C.SDL_GameControllerGetAttached(c.ptr)
	if r == 1 {
		return true
	}
	return false
}

// GetJoystick gets the underlying joystick object used by c
func (c GameController) GetJoystick() Joystick {
	ptr := C.SDL_GameControllerGetJoystick(c.ptr)
	return Joystick{ptr}
}

// GameControllerEventState enables or disables controller event polling.
//
// If controller events are disabled, you must call GameControllerLUpdate()
// yourself and check the state of the controller when you want controller
// information.
//
// The state can be one of QUERY, ENABLE, OR IGNORE.
func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

// GameControllerUpdate updates the current state of the open game controllers.
//
// This is called automatically by the event loop if any game controller events
// are enabled.
func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

type ControllerAxis int32

const (
	CONTROLLER_AXIS_INVALID      ControllerAxis = C.SDL_CONTROLLER_AXIS_INVALID
	CONTROLLER_AXIS_LEFTX        ControllerAxis = C.SDL_CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY        ControllerAxis = C.SDL_CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX       ControllerAxis = C.SDL_CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY       ControllerAxis = C.SDL_CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT  ControllerAxis = C.SDL_CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT ControllerAxis = C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT
	CONTROLLER_AXIS_MAX          ControllerAxis = C.SDL_CONTROLLER_AXIS_MAX
)

// GameControllerGetAxisFromString turns pchString into an axis mapping.
func GameControllerGetAxisFromString(pchString string) ControllerAxis {
	cstr := C.CString(pchString)
	defer C.free(unsafe.Pointer(cstr))

	r := C.SDL_GameControllerGetAxisFromString(cstr)
	return ControllerAxis(r)
}

// GameControllerGetStringForAxis returns a string mapping for axis.
func GameControllerGetStringForAxis(axis ControllerAxis) string {
	return C.GoString(C.SDL_GameControllerGetStringForAxis(C.SDL_GameControllerAxis(axis)))
}

// GetBindForAxis gets the joystick layer binding for this controller axis
// mapping.  This returns a GameControllerButtonBind, GameControllerAxisBind
// or GameControllerHatBind.
func (c GameController) GetBindForAxis(axis ControllerAxis) interface{} {
	r := C.SDL_GameControllerGetBindForAxis(c.ptr, C.SDL_GameControllerAxis(axis))

	var typ C.SDL_GameControllerBindType
	var v1, v2 C.int
	C.extractBind(r, &typ, &v1, &v2)
	return buildBind(ControllerBindtype(typ), int32(v1), int32(v2))
}

// GetAxis gets the current state of an axis control on c.  The state is a
// a value ranging from -32768 to 32767.
func (c GameController) GetAxis(axis ControllerAxis) int16 {
	r := C.SDL_GameControllerGetAxis(c.ptr, C.SDL_GameControllerAxis(axis))
	return int16(r)
}

type ControllerButton int32

const (
	CONTROLLER_BUTTON_INVALID       ControllerButton = C.SDL_CONTROLLER_BUTTON_INVALID
	CONTROLLER_BUTTON_A             ControllerButton = C.SDL_CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B             ControllerButton = C.SDL_CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X             ControllerButton = C.SDL_CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y             ControllerButton = C.SDL_CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK          ControllerButton = C.SDL_CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE         ControllerButton = C.SDL_CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START         ControllerButton = C.SDL_CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK     ControllerButton = C.SDL_CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK    ControllerButton = C.SDL_CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER  ControllerButton = C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER ControllerButton = C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP       ControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN     ControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT     ControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT    ControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX           ControllerButton = C.SDL_CONTROLLER_BUTTON_MAX
)

// GameControllerGetButtonFromString turns pchString into a button mapping.
func GameControllerGetButtonFromString(pchString string) ControllerButton {
	cstr := C.CString(pchString)
	defer C.free(unsafe.Pointer(cstr))

	r := C.SDL_GameControllerGetButtonFromString(cstr)
	return ControllerButton(r)
}

// GameControllerGetStringForButton gets a string mapping for button.
func GameControllerGetStringForButton(button ControllerButton) string {
	return C.GoString(C.SDL_GameControllerGetStringForButton(C.SDL_GameControllerButton(button)))
}

// GetBindForButton gets the sdl joystick layer binding for this controller
// button mapping.  This returns a GameControllerButtonBind,
// GameControllerAxisBind or GameControllerHatBind.
func (c GameController) GetBindForButton(button ControllerButton) interface{} {
	r := C.SDL_GameControllerGetBindForButton(c.ptr, C.SDL_GameControllerButton(button))

	var typ C.SDL_GameControllerBindType
	var v1, v2 C.int
	C.extractBind(r, &typ, &v1, &v2)
	return buildBind(ControllerBindtype(typ), int32(v1), int32(v2))
}

// GetButton gets the current state of a button on c.
func (c GameController) GetButton(button ControllerButton) uint8 {
	r := C.SDL_GameControllerGetButton(c.ptr, C.SDL_GameControllerButton(button))
	return uint8(r)
}

// Close closes c.
func (c GameController) Close() {
	C.SDL_GameControllerClose(c.ptr)
	c.ptr = nil
}

// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

// GetKeyboardFocus gets the window which currently has keyboard focus.  focus
// is false if no window has focus.
func GetKeyboardFocus() (window Window, focus bool) {
	window.ptr = C.SDL_GetKeyboardFocus()
	if window.ptr != nil {
		focus = true
	}
	return
}

// GetKeyboardState gets a snapshot of the current state of the keyboard.
// Indexes into the returned slice are obtained by use Scancode values.
func GetKeyboardState() []uint8 {
	var numkeys C.int
	start := C.SDL_GetKeyboardState(&numkeys)
	sh := reflect.SliceHeader{}
	sh.Len = int(numkeys)
	sh.Cap = int(numkeys)
	sh.Data = uintptr(unsafe.Pointer(start))
	return *(*[]uint8)(unsafe.Pointer(&sh))
}

// GetModState gets the current modifier state for the keyboard.
func GetModState() Keymod {
	return Keymod(C.SDL_GetModState())
}

// SetModState set the current key modifier state for the keyboard.  This does
// not change the keyboard state, only the modifier flags.
func SetModState(modstate Keymod) {
	C.SDL_SetModState(C.SDL_Keymod(modstate))
}

// GetKeyFromScancode gets the key code corresponding to the given scancode
// according to the current keyboard layout.
func GetKeyFromScancode(scancode Scancode) Keycode {
	return Keycode(C.SDL_GetKeyFromScancode(C.SDL_Scancode(scancode)))
}

// GetScancodeFromKey gets teh scancode corresponding to the given key code
// according to the current keyboard layout.
func GetScancodeFromKey(key Keycode) Scancode {
	return Scancode(C.SDL_GetScancodeFromKey(C.SDL_Keycode(key)))
}

// GetScancodeName gets a human-readable name for a scancode.  If scancode
// does not have a name an empty string is returned.
func GetScancodeName(scancode Scancode) string {
	return C.GoString(C.SDL_GetScancodeName(C.SDL_Scancode(scancode)))
}

// GetScancodeFromName gets a scancode from a human-readable name.  If the
// name is not recognized SCANCODE_UNKNOWN is returned.
func GetScancodeFromName(name string) Scancode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return Scancode(C.SDL_GetScancodeFromName(cstr))
}

// GetKeyName gets a human-readable name for key.  If key does not have a name
// an empty string is returned.
func GetKeyName(key Keycode) string {
	return C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(key)))
}

// GetKeyFromName gets a key code from a human-readable name.  If name is not
// recognized K_UNKNOWN is returned.
func GetKeyFromName(name string) Keycode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return Keycode(C.SDL_GetKeyFromName(cstr))
}

// StartTextInput enables text input events.  It will show the on-screen
// keyboard if supported.
func StartTextInput() {
	C.SDL_StartTextInput()
}

// IsTextInputActive returns whether or not Unicode text input events are
// enabled.
func IsTextInputActive() bool {
	if C.SDL_IsTextInputActive() == C.SDL_TRUE {
		return true
	}
	return false
}

// StopTextInput disables text input events.  It will hide the on-screen
// keyboard if supported.
func StopTextInput() {
	C.SDL_StopTextInput()
}

// SetTextInputRect sets the rectangle used to type Unicode text inputs.
func SetTextInputRect(rect *Rect) {
	C.SDL_SetTextInputRect((*C.SDL_Rect)(unsafe.Pointer(rect)))
}

// HasScreenKeyboardSupport returns whether the platform has some screen
// keyboard support.
//
// Note: Not all screen keyboard functions are supported on all platforms.
func HasScreenKeyboardSupport() bool {
	if C.SDL_HasScreenKeyboardSupport() == C.SDL_TRUE {
		return true
	}
	return false
}

// IsScreenKeyboardShown returns whether the screen keyboard is shown for
// window.
//
// Note: May always return false on some platforms (not implemented there).
func (window Window) IsScreenKeyboardShown() bool {
	r := C.SDL_IsScreenKeyboardShown(window.ptr)
	if r == C.SDL_TRUE {
		return true
	}
	return false
}

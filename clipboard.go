// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

// SetClipboardText puts text into the clipboard.
func SetClipboardText(text string) int {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.SDL_SetClipboardText(cstr))
}

// GetClipboardText gets UTF-8 text from the clipboard.
func GetClipboardText() string {
	cstr := C.SDL_GetClipboardText()
	defer C.SDL_free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// HasClipboardText returns whether or not the clipboard exists and contains
// a text string that is non-empty.
func HasClipboardText() bool {
	if b := C.SDL_HasClipboardText(); b == C.SDL_TRUE {
		return true
	}
	return false
}

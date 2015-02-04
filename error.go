// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
static void SetError(const char *msg) { SDL_SetError("%s", msg); }
*/
import "C"

import "unsafe"

// SetError sets the SDL error message to msg.  SetError will replace any
// previous error message.
func SetError(msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))
	C.SetError(cmsg)
}

// GetError gets the SDL error message.  The returned string will contain
// information about the last error that occurred, or be empty if no error
// has occurred since the last call to ClearError.
//
// You should not have to call this in most cases as SDL functions that can
// return an error should return a SDLError, which already contains the SDL
// error message.
func GetError() string {
	return C.GoString(C.SDL_GetError())
}

// ClearError clears the SDL error message.
func ClearError() {
	C.SDL_ClearError()
}

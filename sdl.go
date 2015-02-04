// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package go-sdl2/sdl provides access to the SDL2 library.
package sdl

/*
#cgo pkg-config: sdl2
#include "SDL.h"
*/
import "C"

// SDLError is returned when sdl returns an error.  Msg is the message from
// GetError() and Value is the error code or 0 for nil pointers.
type SDLError struct {
	Msg   string
	Value int
}

func (e SDLError) Error() string {
	return e.Msg
}

// sdlError creates a new SDLError.
func sdlError(i int) SDLError {
	return SDLError{GetError(), i}
}

type InitFlags uint32

const (
	INIT_TIMER          InitFlags = C.SDL_INIT_TIMER
	INIT_AUDIO          InitFlags = C.SDL_INIT_AUDIO
	INIT_VIDEO          InitFlags = C.SDL_INIT_VIDEO
	INIT_JOYSTICK       InitFlags = C.SDL_INIT_JOYSTICK
	INIT_GAMECONTROLLER InitFlags = C.SDL_INIT_GAMECONTROLLER // Turn on game controller also implicitly does INIT_JOYSTICK
	INIT_NOPARACHUTE    InitFlags = C.SDL_INIT_NOPARACHUTE    // Don't catch fatal signals
	INIT_EVERYTHING     InitFlags = C.SDL_INIT_EVERYTHING
)

// Init initializes the subsystems specified by flags. Unless the
// INIT_NOPARACHUTE flag is set, Init will install cleanup signal handlers
// for some commonly ignored fatel signals (like SIGSEGV).
func Init(flags InitFlags) error {
	if r := int(C.SDL_Init(C.Uint32(flags))); r != 0 {
		return sdlError(r)
	}
	return nil
}

// InitSubSystem initializes specific SDL subsystems.
func InitSubSystem(flags InitFlags) error {
	if r := int(C.SDL_InitSubSystem(C.Uint32(flags))); r != 0 {
		return sdlError(r)
	}
	return nil
}

// QuitSubSystem cleans up a specific SDL subsystems.
func QuitSubSystem(flags InitFlags) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

// WasInit returns a mask of the specified subsystems which have previously
// been initialized.  If flags is 0 a mask of all initialized subsystems will
// be returned.
func WasInit(flags InitFlags) InitFlags {
	return InitFlags(C.SDL_WasInit(C.Uint32(flags)))
}

// Quit cleans up all initialized subsystems.  You should call it upon all
// exit conditions.
func Quit() {
	C.SDL_Quit()
}

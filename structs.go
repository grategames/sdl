// +build ignore

// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Use this file to help define structs "go tool cgo -godefs structs.go > structs_386.go"

package sdl

/*
#cgo pkg-config: sdl2
#include "SDL.h"
#include "SDL_syswm.h"
#include "SDL_haptic.h"
*/
import "C"

//SDL_events.h
type WindowEvent C.SDL_WindowEvent
type KeyboardEvent C.SDL_KeyboardEvent
type TextEditingEvent C.SDL_TextEditingEvent
type TextInputEvent C.SDL_TextInputEvent
type MouseMotionEvent C.SDL_MouseMotionEvent
type MouseButtonEvent C.SDL_MouseButtonEvent
type MouseWheelEvent C.SDL_MouseWheelEvent
type JoyAxisEvent C.SDL_JoyAxisEvent
type JoyBallEvent C.SDL_JoyBallEvent
type JoyHatEvent C.SDL_JoyHatEvent
type JoyButtonEvent C.SDL_JoyButtonEvent
type JoyDeviceEvent C.SDL_JoyDeviceEvent
type ControllerAxisEvent C.SDL_ControllerAxisEvent
type ControllerButtonEvent C.SDL_ControllerButtonEvent
type ControllerDeviceEvent C.SDL_ControllerDeviceEvent
type TouchFingerEvent C.SDL_TouchFingerEvent
type MultiGestureEvent C.SDL_MultiGestureEvent
type DollarGestureEvent C.SDL_DollarGestureEvent
type DropEvent C.SDL_DropEvent
type QuitEvent C.SDL_QuitEvent
type UserEvent C.SDL_UserEvent
type SysWMEvent C.SDL_SysWMEvent
type EventUnion C.SDL_Event

//SDL_joystick.h
type JoystickGUID C.SDL_JoystickGUID

//SDL_keyboard.h
type Keysym C.SDL_Keysym

//SDL_pixels.h
type Color C.SDL_Color
type Palette C.SDL_Palette
type PixelFormat C.SDL_PixelFormat

//SDL_rect.h
type Point C.SDL_Point
type Rect C.SDL_Rect

//SDL_surface.h
type Surface C.SDL_Surface

//SDL_version.h
type Version C.SDL_version

//SDL_video.h
type DisplayMode C.SDL_DisplayMode

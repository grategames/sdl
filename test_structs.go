// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// These types are used to test if the structs have the same memory layout in
// C and Go.

package sdl

/*
#include "SDL.h"
#include "SDL_syswm.h"
#include "SDL_haptic.h"
*/
import "C"

//SDL_events.h
type testWindowEvent C.SDL_WindowEvent
type testKeyboardEvent C.SDL_KeyboardEvent
type testTextEditingEvent C.SDL_TextEditingEvent
type testTextInputEvent C.SDL_TextInputEvent
type testMouseMotionEvent C.SDL_MouseMotionEvent
type testMouseButtonEvent C.SDL_MouseButtonEvent
type testMouseWheelEvent C.SDL_MouseWheelEvent
type testJoyAxisEvent C.SDL_JoyAxisEvent
type testJoyBallEvent C.SDL_JoyBallEvent
type testJoyHatEvent C.SDL_JoyHatEvent
type testJoyButtonEvent C.SDL_JoyButtonEvent
type testJoyDeviceEvent C.SDL_JoyDeviceEvent
type testControllerAxisEvent C.SDL_ControllerAxisEvent
type testControllerButtonEvent C.SDL_ControllerButtonEvent
type testControllerDeviceEvent C.SDL_ControllerDeviceEvent
type testTouchFingerEvent C.SDL_TouchFingerEvent
type testMultiGestureEvent C.SDL_MultiGestureEvent
type testDollarGestureEvent C.SDL_DollarGestureEvent
type testDropEvent C.SDL_DropEvent
type testQuitEvent C.SDL_QuitEvent
type testUserEvent C.SDL_UserEvent
type testSysWMEvent C.SDL_SysWMEvent
type testEventUnion C.SDL_Event

//SDL_joystick.h
type testJoystickGUID C.SDL_JoystickGUID

//SDL_keyboard.h
type testKeysym C.SDL_Keysym

//SDL_pixels.h
type testColor C.SDL_Color
type testPalette C.SDL_Palette
type testPixelFormat C.SDL_PixelFormat

//SDL_rect.h
type testPoint C.SDL_Point
type testRect C.SDL_Rect

//SDL_render.h
type testRendererInfo C.SDL_RendererInfo

//SDL_surface.h
type testSurface C.SDL_Surface

//SDL_version.h
type testVersion C.SDL_version

//SDL_video.h
type testDisplayMode C.SDL_DisplayMode

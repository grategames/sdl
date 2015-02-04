// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

// These come from SDL_hints.h. Use them as the name argument in the Hint functions.
const (
	// HINT_FRAMEBUFFER_ACCELERATION controls how 3D acceleration is used to accelerate
	// the screen surface.
	//
	// HINT_FRAMEBUFFER_ACCELERATION can be set to the following values:
	//  "0" - Disable 3D acceleration
	//  "1" - Enable 3D acceleration, using the default renderer.
	//  "X" - Enable 3D acceleration, using X where x is one of the valid drivers.
	//  (e.g. "direct3d", "opengl", etc.)
	//
	// By default SDL tries to make a best guess for each platform whether to use
	// acceleration or not.
	HINT_FRAMEBUFFER_ACCELERATION = C.SDL_HINT_FRAMEBUFFER_ACCELERATION

	// HINT_RENDER_DRIVER specifies which render driver to use.
	//
	// If the application doesn't pick a specific renderer to use, HINT_RENDER_DRIVER
	// specifies the name of the preferred renderer. If the preferred renderer can't
	// be initialized, the normal default renderer is used.
	//
	// HINT_RENDER_DRIVER is case insensitive and can be set to the following values:
	//  "direct3d"
	//  "opengl"
	//  "opengles2"
	//  "opengles"
	//  "software"
	//
	// The default varies by platform, but it's the first one in the list that is
	// available on the current platform.
	HINT_RENDER_DRIVER = C.SDL_HINT_RENDER_DRIVER

	// HINT_RENDER_OPENGL_SHADERS controls whether the OpenGL render driver uses shaders
	// if they are available.
	//
	// HINT_RENDER_OPENGL_SHADERS can be set to the following values:
	//  "0" - Disable shaders
	//  "1" - Enable shaders
	//
	// By default shaders are used if OpenGL supports them.
	HINT_RENDER_OPENGL_SHADERS = C.SDL_HINT_RENDER_OPENGL_SHADERS

	// HINT_RENDER_SCALE_QUALITY controles the scaling quality.
	//
	// HINT_RENDER_SCALE_QUALITY can be set to the following values:
	//  "0" or "nearest" - Nearest pixel sampling
	//  "1" or "linear"  - Linear filtering (supported by OpenGL and Direct3D)
	//  "2" or "best"    - Anisotropic filtering (supported by Direct3D)
	//
	// By default nearest pixel sampling is used.
	HINT_RENDER_SCALE_QUALITY = C.SDL_HINT_RENDER_SCALE_QUALITY

	// HINT_RENDER_VSYNC controls whether updates to the screen surface should be
	// synchronized with the vertical refresh, to avoid tearing.
	//
	// HINT_RENDER_VSYNC can be set to the following values:
	//  "0" - Disable vsync
	//  "1" - Enable vsync
	//
	// By default SDL does not sync screen surface updates with vertical refresh.
	HINT_RENDER_VSYNC = C.SDL_HINT_RENDER_VSYNC

	// HINT_VIDEO_X11_XVIDMODE controls whether the X11 VidMode extension should
	// be used
	//
	// HINT_VIDEO_X11_XVIDMODE can be set to the following values:
	//  "0" - Disable XVidMode
	//  "1" - Enable XVidMode
	//
	// By default SDL will use XVidMode if it is available.
	HINT_VIDEO_X11_XVIDMODE = C.SDL_HINT_VIDEO_X11_XVIDMODE

	// HINT_VIDEO_X11_XINERAMA controls whether the X11 Xinerama extension should
	// be used
	//
	// HINT_VIDEO_X11_XINERAMA can be set to the following values:
	//  "0" - Disable Xinerama
	//  "1" - Enable Xinerama
	//
	// By default SDL will use Xinerama if it is available.
	HINT_VIDEO_X11_XINERAMA = C.SDL_HINT_VIDEO_X11_XINERAMA

	// HINT_VIDEO_X11_XRANDR controls whether the X11 XRandR extension should
	// be used
	//
	// HINT_VIDEO_X11_XRANDR can be set to the following values:
	//  "0" - Disable XRandR
	//  "1" - Enable XRandR
	//
	// By default SDL will not use XRandR because of window manager issues.
	HINT_VIDEO_X11_XRANDR = C.SDL_HINT_VIDEO_X11_XRANDR

	// HINT_GRAB_KEYBOARD controls whether grabbing input grabs the keyboard.
	//
	// HINT_GRAB_KEYBOARD can be set to the following values:
	//  "0" - Grab will affect only the mouse
	//  "1" - Grab will affect mouse and keyboard
	//
	// By default SDL will not grab the keyboard so system shortcuts still work.
	HINT_GRAB_KEYBOARD = C.SDL_HINT_GRAB_KEYBOARD

	// HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS controls if Windows will minimize if they
	// lose key focus when in Fullscreen mode.
	//
	// Defaults to true.
	HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS = C.SDL_HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS

	// HINT_IDLE_TIMER_DISABLED controls whether the idle timer is disabled on iOS.
	//
	// When an iOS app does not receive touches for some time, the screen is dimmed
	// automatically. For games where the accelerometer is the only input this is
	// problematic. This functionality can be disabled by setting this hint.
	//
	// HINT_IDLE_TIMER_DISABLED can be set to the following values:
	//  "0" - Enable idle timer
	//  "1" - Disable idle timer
	HINT_IDLE_TIMER_DISABLED = C.SDL_HINT_IDLE_TIMER_DISABLED

	// HINT_ORIENTATIONS controls which orientations are allowed on iOS.
	//
	// In some circumstances it is necessary to be able to explicitly control which UI
	// orientations are allowed.
	//
	// HINT-ORIENTATIONS is a space delimited list of the following values:
	//  "LandscapeLeft", "LandscapeRight", "Portrait", "PortraitUpsideDown"
	HINT_ORIENTATIONS = C.SDL_HINT_ORIENTATIONS

	// HINT_XINPUT_ENABLED lets you disable the detection and use of Xinput gamepad
	// devices.
	//
	// HINT_XINPUT_ENABLED can be set to the following values:
	//  "0" - Disable XInput timer (only uses direct input)
	//  "1" - Enable XInput timer (the default)
	HINT_XINPUT_ENABLED = C.SDL_HINT_XINPUT_ENABLED

	// HINT_GAMECONTROLLERCONFIG lets you manually hint extra gamecontroller db entries.
	//
	// The variable expected newline delimited rows of gamecontroller config data.
	HINT_GAMECONTROLLERCONFIG = C.SDL_HINT_GAMECONTROLLERCONFIG

	// If HINT_ALLOW_TOPMOST is set to 0 then never set the top most bit on a SDL Window,
	// even if the video mode expects it.
	//
	// This is a debugging aid for developers and not expected to be used by end users.
	// The default is "1".
	//
	// HINT_ALLOW_TOPMOST can be set to the following values:
	//  "0" - don't allow topmost
	//  "1" - allow topmost
	HINT_ALLOW_TOPMOST = C.SDL_HINT_ALLOW_TOPMOST
)

type HintPriority uint32

const (
	HINT_DEFAULT  HintPriority = C.SDL_HINT_DEFAULT
	HINT_NORMAL   HintPriority = C.SDL_HINT_NORMAL
	HINT_OVERRIDE HintPriority = C.SDL_HINT_OVERRIDE
)

// SetHintWithPriority sets a hint with a specific priority.
//
// The priority controls the behavior when setting a hint that already
// has a value.  Hints will replace existing hints of their priority and
// lower.  Environment variables are considered to have override priority.
//
// SetHintWithPriority returns true if the hint was set or false otherwise.
func SetHintWithPriority(name, value string, priority HintPriority) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	result := C.SDL_SetHintWithPriority(cname, cvalue, C.SDL_HintPriority(priority))
	if result == C.SDL_TRUE {
		return true
	}
	return false
}

// SetHint sets a hint with normal priority.  It returns true if the hint was set
// or false otherwise.
func SetHint(name, value string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	if b := C.SDL_SetHint(cname, cvalue); b == C.SDL_TRUE {
		return true
	}
	return false
}

// GetHint returns the value of a hint variable.
func GetHint(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.SDL_GetHint(cname))
}

// ClearHints clears all hints.  It is called during Quit() to free stored hints.
func ClearHints() {
	C.SDL_ClearHints()
}

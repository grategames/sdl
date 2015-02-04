// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

type Cursor struct {
	ptr *C.SDL_Cursor
}

type SystemCursor int32

const (
	SYSTEM_CURSOR_ARROW     SystemCursor = C.SDL_SYSTEM_CURSOR_ARROW     // Arrow
	SYSTEM_CURSOR_IBEAM     SystemCursor = C.SDL_SYSTEM_CURSOR_IBEAM     // I-beam
	SYSTEM_CURSOR_WAIT      SystemCursor = C.SDL_SYSTEM_CURSOR_WAIT      // Wait
	SYSTEM_CURSOR_CROSSHAIR SystemCursor = C.SDL_SYSTEM_CURSOR_CROSSHAIR // Crosshair
	SYSTEM_CURSOR_WAITARROW SystemCursor = C.SDL_SYSTEM_CURSOR_WAITARROW // Small wait cursor (or Wait if not available)
	SYSTEM_CURSOR_SIZENWSE  SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENWSE  // Double arrow pointing northwest and southeast
	SYSTEM_CURSOR_SIZENESW  SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENESW  // Double arrow pointing northeast and southwest
	SYSTEM_CURSOR_SIZEWE    SystemCursor = C.SDL_SYSTEM_CURSOR_SIZEWE    // Double arrow pointing west and east
	SYSTEM_CURSOR_SIZENS    SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENS    // Double arrow pointing north and south
	SYSTEM_CURSOR_SIZEALL   SystemCursor = C.SDL_SYSTEM_CURSOR_SIZEALL   // Four pointed arrow pointing north, south, east, and west
	SYSTEM_CURSOR_NO        SystemCursor = C.SDL_SYSTEM_CURSOR_NO        // Slashed circle or crossbones
	SYSTEM_CURSOR_HAND      SystemCursor = C.SDL_SYSTEM_CURSOR_HAND      // Hand
)

// GetMouseFocus gets the window which currently has mouse focus.  focus is
// false if no window has focus.
func GetMouseFocus() (window Window, focus bool) {
	window.ptr = C.SDL_GetMouseFocus()
	if window.ptr != nil {
		focus = true
	}
	return
}

// GetMouseState gets the current state of the mouse.
//
// The current button state is returned as a button bitmask, which can be
// tested using the Button function, and x and y are the mouse cursor
// position relative to the focus window for the currently selected mouse.
func GetMouseState() (x, y int, state uint32) {
	state = uint32(C.SDL_GetMouseState((*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y))))
	return
}

// GetRelativeMouseState get relative state of the mouse.
//
// The current button state is returned as a bitmask, which can be tested
// using the Button function, and x and y are the mouse deltas since the last
// call to GetRelativeMouseState.
func GetRelativeMouseState() (x, y int, state uint32) {
	state = uint32(C.SDL_GetRelativeMouseState((*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y))))
	return
}

// WarpMouse moves the mouse to the given position within the window.
//
// Note: This function generates a mouse motion event.
func (window Window) WarpMouse(x, y int) {
	C.SDL_WarpMouseInWindow(window.ptr, C.int(x), C.int(y))
}

// SetRelativeMouseMode set the relative mouse mode.  While the mouse is in
// relative mode, the cursor is hidden, and the driver will try to report
// continuous motion in the current window.  Only relative motion events will
// be delivered, the mouse position will not change.
//
// Note: This function will flush any pending mouse motion
func SetRelativeMouseMode(enabled bool) error {
	var b C.SDL_bool = C.SDL_FALSE
	if enabled == true {
		b = C.SDL_TRUE
	}
	r := int(C.SDL_SetRelativeMouseMode(b))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetRelativeMouseMode returns whether relative mouse mode is enabled.
func GetRelativeMouseMode() bool {
	if b := C.SDL_GetRelativeMouseMode(); b == C.SDL_TRUE {
		return true
	}
	return false
}

// CreateCursor creates a cursor, using the specified bitmap data and mask
// (in MSB format).  The cursor width must be a multiple of 8 bits.
//
// The cursor is created in black and white according to the following:
//  data mask resulting pixel on screen
//   0    1   White
//   1    1   Black
//   0    0   Transparent
//   1    0   Inverted color if possible, black if not
func CreateCursor(data, mask []uint8, w, h, hot_x, hot_y int) Cursor {
	var cData *C.Uint8
	if len(data) > 0 {
		cData = (*C.Uint8)(&data[0])
	}
	var cMask *C.Uint8
	if len(mask) > 0 {
		cMask = (*C.Uint8)(&mask[0])
	}

	cur := Cursor{}
	cur.ptr = C.SDL_CreateCursor(cData, cMask, C.int(w), C.int(h),
		C.int(hot_x), C.int(hot_y))
	return cur
}

// CreateColorCursor creates a color cursor.
func CreateColorCursor(surface *Surface, hot_x, hot_y int) Cursor {
	cur := Cursor{}
	cur.ptr = C.SDL_CreateColorCursor((*C.SDL_Surface)(unsafe.Pointer(surface)),
		C.int(hot_x), C.int(hot_y))
	return cur
}

// CreateSystemCursor creates a system cursor.
func CreateSystemCursor(id SystemCursor) Cursor {
	cur := Cursor{}
	cur.ptr = C.SDL_CreateSystemCursor(C.SDL_SystemCursor(id))
	return cur
}

// Set sets cursor to be the active cursor.
func (cursor Cursor) Set() {
	C.SDL_SetCursor(cursor.ptr)
}

// GetCursor get the active cursor.
func GetCursor() Cursor {
	cur := Cursor{}
	cur.ptr = C.SDL_GetCursor()
	return cur
}

// Free frees the cursor.
func (cursor Cursor) Free() {
	C.SDL_FreeCursor(cursor.ptr)
	cursor.ptr = nil
}

// ShowCursor toggles whether or not the cursor is shown.  If toggle is 1 this
// will show the cursor, if toggle is 0 this will hide the cursor, and if
// toggle is -1 this will return the current state of the cursor.  Returns
// 1 if the cursor is shown, or 0 if the cursor is hidden.
func ShowCursor(toggle int) int {
	return int(C.SDL_ShowCursor(C.int(toggle)))
}

const (
	BUTTON_LEFT   = C.SDL_BUTTON_LEFT
	BUTTON_MIDDLE = C.SDL_BUTTON_MIDDLE
	BUTTON_RIGHT  = C.SDL_BUTTON_RIGHT
	BUTTON_X1     = C.SDL_BUTTON_X1
	BUTTON_X2     = C.SDL_BUTTON_X2
	BUTTON_LMASK  = C.SDL_BUTTON_LMASK
	BUTTON_MMASK  = C.SDL_BUTTON_MMASK
	BUTTON_RMASK  = C.SDL_BUTTON_RMASK
	BUTTON_X1MASK = C.SDL_BUTTON_X1MASK
	BUTTON_X2MASK = C.SDL_BUTTON_X2MASK
)

// Button is used to check the state returned by GetMouseState or
// GetRelativeMouseState.
func Button(button uint32) uint32 {
	return 1 << (button - 1)
}

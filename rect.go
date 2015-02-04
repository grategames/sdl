// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

// Empty checks if rect has no area.
func (rect *Rect) Empty() bool {
	if rect == nil {
		return true
	}
	return rect.W <= 0 || rect.H <= 0
}

// HasIntersection determines whether a and b intersect.
func (a *Rect) HasIntersection(b *Rect) bool {
	if r := C.SDL_HasIntersection((*C.SDL_Rect)(unsafe.Pointer(a)),
		(*C.SDL_Rect)(unsafe.Pointer(b))); r == C.SDL_TRUE {
		return true
	}
	return false
}

// IntersectRect calculates the intersection of a and b.
func (a *Rect) IntersectRect(b *Rect) (rect *Rect, intersect bool) {
	rect = &Rect{}
	r := C.SDL_IntersectRect((*C.SDL_Rect)(unsafe.Pointer(a)),
		(*C.SDL_Rect)(unsafe.Pointer(b)), (*C.SDL_Rect)(unsafe.Pointer(rect)))
	if r == C.SDL_TRUE {
		intersect = true
	}
	return
}

// UnionRect calculates the union of a and b.
func (a *Rect) UnionRect(b *Rect) *Rect {
	result := new(Rect)
	C.SDL_UnionRect((*C.SDL_Rect)(unsafe.Pointer(a)), (*C.SDL_Rect)(unsafe.Pointer(b)),
		(*C.SDL_Rect)(unsafe.Pointer(result)))
	return result
}

// EclosedPoints calculates a minimal rectangle enclosing a set of points.
func EnclosePoints(points []Point, clip *Rect) (rect *Rect, result bool) {
	var ptr *C.SDL_Point
	if len(points) > 0 {
		ptr = (*C.SDL_Point)(unsafe.Pointer(&points[0]))
	}

	rect = &Rect{}
	r := C.SDL_EnclosePoints(ptr, C.int(len(points)),
		(*C.SDL_Rect)(unsafe.Pointer(clip)),
		(*C.SDL_Rect)(unsafe.Pointer(rect)))

	if r == C.SDL_TRUE {
		result = true
	}
	return
}

// IntersectRectAndLine calculates the intersection of a rectangle and line
// segment.
func IntersectRectAndLine(rect *Rect, x1, y1, x2, y2 *int) bool {
	if r := C.SDL_IntersectRectAndLine((*C.SDL_Rect)(unsafe.Pointer(rect)),
		(*C.int)(unsafe.Pointer(x1)), (*C.int)(unsafe.Pointer(y1)),
		(*C.int)(unsafe.Pointer(x2)), (*C.int)(unsafe.Pointer(y2))); r == C.SDL_TRUE {
		return true
	}
	return false
}

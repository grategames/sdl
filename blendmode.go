// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

type BlendMode uint32

const (
	// No blending
	BLENDMODE_NONE BlendMode = C.SDL_BLENDMODE_NONE
	// dst = (src * A) + (dst * (1-A))
	BLENDMODE_BLEND BlendMode = C.SDL_BLENDMODE_BLEND
	// dst = (src * A) + dst
	BLENDMODE_ADD BlendMode = C.SDL_BLENDMODE_ADD
	// dst = src * dst
	BLENDMODE_MOD BlendMode = C.SDL_BLENDMODE_MOD
)

// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

const (
	MAJOR_VERSION = C.SDL_MAJOR_VERSION
	MINOR_VERSION = C.SDL_MINOR_VERSION
	PATCHLEVEL    = C.SDL_PATCHLEVEL
)

// VERSION sets v to the version of SDL compiled against.  It is based on the
// header the compiler used.
func VERSION(v *Version) {
	v.Major = MAJOR_VERSION
	v.Minor = MINOR_VERSION
	v.Patch = PATCHLEVEL
}

var COMPILEDVERSION int = VersionNum(MAJOR_VERSION, MINOR_VERSION, PATCHLEVEL)

// VersionNum turns the version numbers into a numeric value.
//
//  i := VersionNum(1, 2, 3)
//  i == 1203
func VersionNum(x, y, z int) int {
	return (x*1000 + y*100 + z)
}

// VersionAtleast returns true if the COMPILEDVERSION is >= X.Y.Z
func VersionAtleast(x, y, z int) bool {
	return COMPILEDVERSION >= VersionNum(x, y, z)
}

// GetVersion returns the version of SDL that is linked against your program.
//
// GetVersion may be called safely at any time, even before Init().
func GetVersion() *Version {
	v := new(Version)
	C.SDL_GetVersion((*C.SDL_version)(unsafe.Pointer(v)))
	return v
}

// GetRvision returns an arbitrary string (a hash value) uniquely identifying
// the exact revision of the SDL library in use.  It is only useful in comparing
// against other revisions. It is NOT an incrementing number.
func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}

// GetRevisionNumber returns the number uniquely identifying the exact revision
// of the SDL library in use.  It is an incrementing number based on commits to
// hg.libsdl.org.
func GetRevisionNumber() int {
	return int(C.SDL_GetRevisionNumber())
}

// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
#include "SDL_syswm.h"
#include "syswm.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

type SysWMType uint32

const (
	SYSWM_UNKNOWN  SysWMType = C.SDL_SYSWM_UNKNOWN
	SYSWM_WINDOWS  SysWMType = C.SDL_SYSWM_WINDOWS
	SYSWM_X11      SysWMType = C.SDL_SYSWM_X11
	SYSWM_DIRECTFB SysWMType = C.SDL_SYSWM_DIRECTFB
	SYSWM_COCOA    SysWMType = C.SDL_SYSWM_COCOA
	SYSWM_UIKIT    SysWMType = C.SDL_SYSWM_UIKIT
)

// SysWMmsg contains system-dependent window manager messages.
//
// The types of the fields in SysWMmsg are:
//  win.hwnd   HWND
//  win.msg    UINT
//  win.wParam WPARAM
//  win.lParam LPARAM
//
//  x11.event XEvent
//
//  dfb.event DFBEvent
type SysWMmsg C.SDL_SysWMmsg

// Subsystem returns the windowing system type of msg.
func (msg *SysWMmsg) Subsystem() SysWMType {
	return SysWMType(msg.subsystem)
}

// WinHwnd returns a pointer to msg.win.hwnd. If msg.Subsystem is not
// SYSWM_WINDOWS, WinHwnd will return an error.
func (msg *SysWMmsg) WinHwnd() (uintptr, error) {
	var ptr unsafe.Pointer
	if msg.subsystem == C.SDL_SYSWM_WINDOWS {
		ptr = C.win_hwnd((*C.SDL_SysWMmsg)(msg))
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("msg.Subsystem is not SYSWM_WINDOWS")
}

// WinMsg returns a pointer to msg.win.msg. If msg.Subsystem is not
// SYSWM_WINDOWS, WinMsg will return an error.
func (msg *SysWMmsg) WinMsg() (uintptr, error) {
	var ptr unsafe.Pointer
	if msg.subsystem == C.SDL_SYSWM_WINDOWS {
		ptr = C.win_msg((*C.SDL_SysWMmsg)(msg))
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("msg.Subsystem is not SYSWM_WINDOWS")
}

// WinWParam returns a pointer to msg.win.wParam. If msg.Subsystem is not
// SYSWM_WINDOWS, WinWParam will return an error.
func (msg *SysWMmsg) WinWParam() (uintptr, error) {
	var ptr unsafe.Pointer
	if msg.subsystem == C.SDL_SYSWM_WINDOWS {
		ptr = C.win_wParam((*C.SDL_SysWMmsg)(msg))
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("msg.Subsystem is not SYSWM_WINDOWS")
}

// WinLParam returns a pointer to msg.win.lParam. If msg.Subsystem is not
// SYSWM_WINDOWS, WinLParam will return an error.
func (msg *SysWMmsg) WinLParam() (uintptr, error) {
	var ptr unsafe.Pointer
	if msg.subsystem == C.SDL_SYSWM_WINDOWS {
		ptr = C.win_lParam((*C.SDL_SysWMmsg)(msg))
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("msg.Subsystem is not SYSWM_WINDOWS")
}

// X11Event returns a pointer to msg.x11.event. If msg.Subsystem is not
// SYSWM_X11, X11Event will return an error.
func (msg *SysWMmsg) X11Event() (uintptr, error) {
	var ptr unsafe.Pointer
	if msg.subsystem == C.SDL_SYSWM_X11 {
		ptr = C.x11_event((*C.SDL_SysWMmsg)(msg))
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("msg.Subsystem is not SYSWM_X11")
}

// DFBEvent returns a pointer to msg.dfb.event. If msg.Subsystem is not
// SYSWM_DIRECTFB, DFBEvent will return an error.
func (msg *SysWMmsg) DFBEvent() (uintptr, error) {
	var ptr unsafe.Pointer
	if msg.subsystem == C.SDL_SYSWM_DIRECTFB {
		ptr = C.dfb_event((*C.SDL_SysWMmsg)(msg))
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("msg.Subsystem is not SYSWM_DIRECTFB")
}

// SysWMinfo contains system-dependent information about a window.
//
// The types of the fields in SysWMinfo are:
//  win.window HWND
//
//  x11.display *Display
//  x11.window  Window
//
//  dfb.dfb     *IDirectFB
//  dfb.window  *IDirectFBWindow
//  dfb.surface *IDirectFBSurface
//
//  cocoa.window *NSWindow
//
//  uikit.window *UIWindow
type SysWMinfo struct {
	info *C.SDL_SysWMinfo
}

// Subsystem return the windowing system type of info.
func (info SysWMinfo) Subsystem() SysWMType {
	return SysWMType(info.info.subsystem)
}

// WinWindow returns a pointer to info.win.window.  If info.Subsystem is not
// SYSWM_WINDOWS, WinWindow will return an error.
func (info SysWMinfo) WinWindow() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_WINDOWS {
		ptr = C.win_window(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_WINDOWS")
}

// X11Display returns a pointer to info.x11.display.  If info.Subsystem is not
// SYSWM_X11, X11Display will return an error.
func (info SysWMinfo) X11Display() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_X11 {
		ptr := C.x11_display(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_X11")
}

// X11Window returns a pointer to info.x11.window.  If info.Subsystem is not
// SYSWM_X11, X11Window will return an error.
func (info SysWMinfo) X11Window() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_X11 {
		ptr := C.x11_window(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_X11")
}

// DFBdfb returns a pointer to info.dfb.dfb.  If info.Subsystem is not
// SYSWM_DIRECTFB, DFBdfb will return an error.
func (info SysWMinfo) DFBdfb() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_DIRECTFB {
		ptr := C.dfb_dfb(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_DIRECTFB")
}

// DFBWindow returns a pointer to info.dfb.window.  If info.Subsystem is not
// SYSWM_DIRECTFB, DFBWindow will return an error.
func (info SysWMinfo) DFBWindow() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_DIRECTFB {
		ptr := C.dfb_window(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_DIRECTFB")
}

// DFBSurface returns a pointer to info.dfb.surface.  If info.Subsystem is not
// SYSWM_DIRECTFB, DFBSurface will return an error.
func (info SysWMinfo) DFBSurface() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_DIRECTFB {
		ptr := C.dfb_surface(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_DIRECTFB")
}

// CocoaWindow returns a pointer to info.cocoa.window.  If info.Subsystem is
// not SYSWM_COCOA, CocoaWindow will return an error.
func (info SysWMinfo) CocoaWindow() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_COCOA {
		ptr := C.cocoa_window(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_COCOA")
}

// UIKITWindow returns a pointer to info.uikit.window.  If info.Subsystem is
// not SYSWM_UIKIT, UIKITWindow will return an error.
func (info SysWMinfo) UIKitWindow() (uintptr, error) {
	var ptr unsafe.Pointer
	if info.info.subsystem == C.SDL_SYSWM_UIKIT {
		ptr := C.uikit_window(info.info)
		return uintptr(ptr), nil
	}
	return uintptr(ptr), errors.New("info.Subsystem is not SYSWM_UIKIT")
}

// GetWindowWMInfo allows access to driver-depended window information.
func (window Window) GetWMInfo() (SysWMinfo, error) {
	info := new(C.SDL_SysWMinfo)
	VERSION((*Version)(unsafe.Pointer(&info.version)))
	r := C.SDL_GetWindowWMInfo(window.ptr, info)
	if r == C.SDL_FALSE {
		return SysWMinfo{}, sdlError(int(r))
	}
	return SysWMinfo{info}, nil
}

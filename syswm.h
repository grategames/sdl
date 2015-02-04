// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "SDL.h"
#include "SDL_syswm.h"

extern void *win_hwnd(SDL_SysWMmsg *sysWMmsg);
extern void *win_msg(SDL_SysWMmsg *sysWMmsg);
extern void *win_wParam(SDL_SysWMmsg *sysWMmsg);
extern void *win_lParam(SDL_SysWMmsg *sysWMmsg);
extern void *win_window(SDL_SysWMinfo *info);
extern void *x11_event(SDL_SysWMmsg *msg);
extern void *x11_display(SDL_SysWMinfo *info);
extern void *x11_window(SDL_SysWMinfo *info);
extern void *dfb_event(SDL_SysWMmsg *msg);
extern void *dfb_dfb(SDL_SysWMinfo *info);
extern void *dfb_window(SDL_SysWMinfo *info);
extern void *dfb_surface(SDL_SysWMinfo *info);
extern void *cocoa_window(SDL_SysWMinfo *info);
extern void *uikit_window(SDL_SysWMinfo *info);

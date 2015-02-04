// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "syswm.h"

#if defined(SDL_VIDEO_DRIVER_WINDOWS)
void *win_hwnd(SDL_SysWMmsg *sysWMmsg) { return &sysWMmsg->msg.win.hwnd; }
void *win_msg(SDL_SysWMmsg *sysWMmsg) { return &sysWMmsg->msg.win.msg; }
void *win_wParam(SDL_SysWMmsg *sysWMmsg) { return &sysWMmsg->msg.win.wParam; }
void *win_lParam(SDL_SysWMmsg *sysWMmsg) { return &sysWMmsg->msg.win.lParam; }
void *win_window(SDL_SysWMinfo *info) { return &info->info.win.window; }
#else
void *win_hwnd(SDL_SysWMmsg *sysWMmsg) { return NULL; }
void *win_msg(SDL_SysWMmsg *sysWMmsg) { return NULL; }
void *win_wParam(SDL_SysWMmsg *sysWMmsg) { return NULL; }
void *win_lParam(SDL_SysWMmsg *sysWMmsg) { return NULL; }
void *win_window(SDL_SysWMinfo *info) { return NULL; }
#endif
#if defined(SDL_VIDEO_DRIVER_X11)
void *x11_event(SDL_SysWMmsg *msg) { return &msg->msg.x11.event; }
void *x11_display(SDL_SysWMinfo *info) { return info->info.x11.display; }
void *x11_window(SDL_SysWMinfo *info) { return &info->info.x11.window; }
#else
void *x11_event(SDL_SysWMmsg *msg) { return NULL; }
void *x11_display(SDL_SysWMinfo *info) { return NULL; }
void *x11_window(SDL_SysWMinfo *info) { return NULL; }
#endif
#if defined(SDL_VIDEO_DRIVER_DIRECTFB)
void *dfb_event(SDL_SysWMmsg *msg) { return &msg->msg.dfb.event; }
void *dfb_dfb(SDL_SysWMinfo *info) { return info->info.dfb.dfb; }
void *dfb_window(SDL_SysWMinfo *info) { return info->info.dfb.window; }
void *dfb_surface(SDL_SysWMinfo *info) { return info->info.dfb.surface; }
#else
void *dfb_event(SDL_SysWMmsg *msg) { return NULL; }
void *dfb_dfb(SDL_SysWMinfo *info) { return NULL; }
void *dfb_window(SDL_SysWMinfo *info) { return NULL; }
void *dfb_surface(SDL_SysWMinfo *info) { return NULL; }
#endif
#if defined(SDL_VIDEO_DRIVER_COCOA)
void *cocoa_window(SDL_SysWMinfo *info) { return info->info.cocoa.window; }
#else
void *cocoa_window(SDL_SysWMinfo *info) { return NULL; }
#endif
#if defined(SDL_VIDEO_DRIVER_UIKIT)
void *uikit_window(SDL_SysWMinfo *info) { return info->info.uikit.window; }
#else
void *uikit_window(SDL_SysWMinfo *info) { return NULL; }
#endif

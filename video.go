// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import "unsafe"

type WindowFlags uint32

const (
	// fullscreen window
	WINDOW_FULLSCREEN WindowFlags = C.SDL_WINDOW_FULLSCREEN
	// window usable with OpenGL context
	WINDOW_OPENGL WindowFlags = C.SDL_WINDOW_OPENGL
	// window is visible
	WINDOW_SHOWN WindowFlags = C.SDL_WINDOW_SHOWN
	// window is not visible
	WINDOW_HIDDEN WindowFlags = C.SDL_WINDOW_HIDDEN
	// no window decoration
	WINDOW_BORDERLESS WindowFlags = C.SDL_WINDOW_BORDERLESS
	// window can be resized
	WINDOW_RESIZABLE WindowFlags = C.SDL_WINDOW_RESIZABLE
	// window is minimized
	WINDOW_MINIMIZED WindowFlags = C.SDL_WINDOW_MINIMIZED
	// window is maximized
	WINDOW_MAXIMIZED WindowFlags = C.SDL_WINDOW_MAXIMIZED
	// window has grabbed input focus
	WINDOW_INPUTGRABBED WindowFlags = C.SDL_WINDOW_INPUT_GRABBED
	// window has input focus
	WINDOW_INPUT_FOCUS WindowFlags = C.SDL_WINDOW_INPUT_FOCUS
	// window has mouse focus
	WINDOW_MOUSE_FOCUS        WindowFlags = C.SDL_WINDOW_MOUSE_FOCUS
	WINDOW_FULLSCREEN_DESKTOP             = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	// Window was not created by SDL
	WINDOW_FOREIGN WindowFlags = C.SDL_WINDOW_FOREIGN
)

const (
	// WINDOWPOS_UNDEFINED is used to indicate that you don't care what the
	// window position is.
	WINDOWPOS_UNDEFINED = C.SDL_WINDOWPOS_UNDEFINED
	// WINDOWPOS_CENTERED is used to indicate that you want the window to be
	// centered.
	WINDOWPOS_CENTERED = C.SDL_WINDOWPOS_CENTERED
)

type WindowEventID uint8

const (
	// never used
	WINDOWEVENT_NONE WindowEventID = C.SDL_WINDOWEVENT_NONE
	// Window has been shown
	WINDOWEVENT_SHOWN WindowEventID = C.SDL_WINDOWEVENT_SHOWN
	// Window has been hidden
	WINDOWEVENT_HIDDEN WindowEventID = C.SDL_WINDOWEVENT_HIDDEN
	// Window has been exposed and should be redrawn
	WINDOWEVENT_EXPOSED WindowEventID = C.SDL_WINDOWEVENT_EXPOSED
	// Window has been moved to Data1, Data2.
	WINDOWEVENT_MOVED WindowEventID = C.SDL_WINDOWEVENT_MOVED
	// Window has been resized to Data1xData2.
	WINDOWEVENT_RESIZED WindowEventID = C.SDL_WINDOWEVENT_RESIZED
	// The window size has changed, either as a result of an API call or
	// through the system or user changing the window size.
	WINDOWEVENT_SIZE_CHANGED WindowEventID = C.SDL_WINDOWEVENT_SIZE_CHANGED
	// Window has been minimized
	WINDOWEVENT_MINIMIZED WindowEventID = C.SDL_WINDOWEVENT_MINIMIZED
	// Window has been maximized
	WINDOWEVENT_MAXIMIZED WindowEventID = C.SDL_WINDOWEVENT_MAXIMIZED
	// Window has been restored to the normal size and position
	WINDOWEVENT_RESTORED WindowEventID = C.SDL_WINDOWEVENT_RESTORED
	// Window has gained mouse focus.
	WINDOWEVENT_ENTER WindowEventID = C.SDL_WINDOWEVENT_ENTER
	// Window has lost mouse focus.
	WINDOWEVENT_LEAVE WindowEventID = C.SDL_WINDOWEVENT_LEAVE
	// Window has gained keyboard focus.
	WINDOWEVENT_FOCUS_GAINED WindowEventID = C.SDL_WINDOWEVENT_FOCUS_GAINED
	// Window has lost keyboard focus.
	WINDOWEVENT_FOCUS_LOST WindowEventID = C.SDL_WINDOWEVENT_FOCUS_LOST
	// Window manager requests that the window be closed.
	WINDOWEVENT_CLOSE WindowEventID = C.SDL_WINDOWEVENT_CLOSE
)

type GLattr uint32

const (
	GL_RED_SIZE                   GLattr = C.SDL_GL_RED_SIZE
	GL_GREEN_SIZE                 GLattr = C.SDL_GL_GREEN_SIZE
	GL_BLUE_SIZE                  GLattr = C.SDL_GL_BLUE_SIZE
	GL_ALPHA_SIZE                 GLattr = C.SDL_GL_ALPHA_SIZE
	GL_BUFFER_SIZE                GLattr = C.SDL_GL_BUFFER_SIZE
	GL_DOUBLEBUFFER               GLattr = C.SDL_GL_DOUBLEBUFFER
	GL_DEPTH_SIZE                 GLattr = C.SDL_GL_DEPTH_SIZE
	GL_STENCIL_SIZE               GLattr = C.SDL_GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE             GLattr = C.SDL_GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE           GLattr = C.SDL_GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE            GLattr = C.SDL_GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE           GLattr = C.SDL_GL_ACCUM_ALPHA_SIZE
	GL_STEREO                     GLattr = C.SDL_GL_STEREO
	GL_MULTISAMPLEBUFFERS         GLattr = C.SDL_GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES         GLattr = C.SDL_GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL         GLattr = C.SDL_GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING           GLattr = C.SDL_GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION      GLattr = C.SDL_GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION      GLattr = C.SDL_GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_EGL                GLattr = C.SDL_GL_CONTEXT_EGL
	GL_CONTEXT_FLAGS              GLattr = C.SDL_GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK       GLattr = C.SDL_GL_CONTEXT_PROFILE_MASK
	GL_SHARE_WITH_CURRENT_CONTEXT GLattr = C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT
)

type GLprofile uint32

const (
	GL_CONTEXT_PROFILE_CORE          GLprofile = C.SDL_GL_CONTEXT_PROFILE_CORE
	GL_CONTEXT_PROFILE_COMPATIBILITY GLprofile = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY
	GL_CONTEXT_PROFILE_ES            GLprofile = C.SDL_GL_CONTEXT_PROFILE_ES
)

type GLcontextFlag uint32

const (
	GL_CONTEXT_DEBUG_FLAG              GLcontextFlag = C.SDL_GL_CONTEXT_DEBUG_FLAG
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG GLcontextFlag = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG
	GL_CONTEXT_ROBUST_ACCESS_FLAG      GLcontextFlag = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG
	GL_CONTEXT_RESET_ISOLATION_FLAG    GLcontextFlag = C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG
)

// Window is used to identify a window.
type Window struct {
	ptr *C.SDL_Window
}

// An opaque handle to an OpenGl context.
type GLContext struct {
	ctx C.SDL_GLContext
}

// GetNumVideoDrivers returns the number of video drivers compiled into SDL.
func GetNumVideoDrivers() int {
	return int(C.SDL_GetNumVideoDrivers())
}

// GetVideoDriver returns the name of a built in video driver.
func GetVideoDriver(index int) string {
	return C.GoString(C.SDL_GetVideoDriver(C.int(index)))
}

// VideoInit initializes the video subsystem, optionally specifying a video
// driver. If driver_name is an empty string ("") the default video driver is
// used. It does not initialize a window or graphics mode.
func VideoInit(driver_name string) error {
	if driver_name == "" {
		if r := int(C.SDL_VideoInit(nil)); r != 0 {
			return sdlError(r)
		}
		return nil
	}
	cdriver_name := C.CString(driver_name)
	defer C.free(unsafe.Pointer(cdriver_name))
	if r := int(C.SDL_VideoInit(cdriver_name)); r != 0 {
		return sdlError(r)
	}
	return nil
}

// VideoQuit shuts down the video subsystem.  It closes all windows, and
// restores the original video mode.
func VideoQuit() {
	C.SDL_VideoQuit()
}

// GetCurrentVideoDriver returns the name of the currently initialized video driver.
func GetCurrentVideoDriver() string {
	return C.GoString(C.SDL_GetCurrentVideoDriver())
}

// GetNumVideoDisplays returns the number of available video displays.
func GetNumVideoDisplays() int {
	return int(C.SDL_GetNumVideoDisplays())
}

// GetDisplayName gets the name of a display in UTF-8 encoding.
func GetDisplayName(displayIndex int) (string, error) {
	r := C.SDL_GetDisplayName(C.int(displayIndex))
	if r == nil {
		return "", sdlError(0)
	}
	return C.GoString(r), nil
}

// GetDisplayBounds returns the desktop area represented by a display, with
// the primary display located at 0,0.
func GetDisplayBounds(displayIndex int) (*Rect, error) {
	area := new(Rect)
	r := int(C.SDL_GetDisplayBounds(C.int(displayIndex),
		(*C.SDL_Rect)(unsafe.Pointer(area))))
	if r != 0 {
		return nil, sdlError(r)
	}
	return area, nil
}

// GetNumDisplayModes returns the number of available display modes.
func GetNumDisplayModes(displayIndex int) int {
	return int(C.SDL_GetNumDisplayModes(C.int(displayIndex)))
}

// GetDisplayMode returns the display mode for the given indexes.
//
// The display modes are sorted in this priority:
//  bits per pixel -> more colors to fewer colors
//  width -> largest to smallest
//  height -> largest to smallest
//  refresh rate -> highest to lowest
func GetDisplayMode(displayIndex, modeIndex int) (*DisplayMode, error) {
	mode := new(DisplayMode)
	r := int(C.SDL_GetDisplayMode(C.int(displayIndex), C.int(modeIndex),
		(*C.SDL_DisplayMode)(unsafe.Pointer(mode))))
	if r != 0 {
		return nil, sdlError(r)
	}
	return mode, nil
}

// GetDesktopDisplayMode returns the desktop display mode.
func GetDesktopDisplayMode(displayIndex int) (*DisplayMode, error) {
	mode := new(DisplayMode)
	r := int(C.SDL_GetDesktopDisplayMode(C.int(displayIndex),
		(*C.SDL_DisplayMode)(unsafe.Pointer(mode))))
	if r != 0 {
		return nil, sdlError(r)
	}
	return mode, nil
}

// GetCurrentDisplayMode returns the current display mode for displayIndex.
func GetCurrentDisplayMode(displayIndex int) (*DisplayMode, error) {
	mode := new(DisplayMode)
	r := int(C.SDL_GetCurrentDisplayMode(C.int(displayIndex),
		(*C.SDL_DisplayMode)(unsafe.Pointer(mode))))
	if r != 0 {
		return nil, sdlError(r)
	}
	return mode, nil
}

// GetClosestDisplayMode returns the closest display mode to the requested
// display mode.
//
// The mode format and refresh_rate default to the desktop mode if they are 0.
// The modes are scanned with size being first priority, format being second
// priority, and finally checking the refresh rate.  If all the available modes
// are too small, then nil is returned.
func GetClosestDisplayMode(displayIndex int, requested *DisplayMode) (*DisplayMode, error) {
	closest := new(DisplayMode)

	r := (C.SDL_GetClosestDisplayMode(C.int(displayIndex),
		(*C.SDL_DisplayMode)(unsafe.Pointer(requested)),
		(*C.SDL_DisplayMode)(unsafe.Pointer(closest))))
	if r == nil {
		return nil, sdlError(0)
	}
	return closest, nil
}

// GetDisplayIndex gets the display index associated with window.
func (window Window) GetDisplayIndex() (int, error) {
	r := int(C.SDL_GetWindowDisplayIndex(window.ptr))
	if r == -1 {
		return 0, sdlError(r)
	}
	return r, nil
}

// SetDisplayMode sets the display mode used when the window is fullscreen and
// visible.  If mode is nil the window's dimensions and the desktop format and
// refresh rate are used.
func (window Window) SetDisplayMode(mode *DisplayMode) error {
	r := int(C.SDL_SetWindowDisplayMode(window.ptr,
		(*C.SDL_DisplayMode)(unsafe.Pointer(mode))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetDisplayMode gets the display mode used if the window is fullscreen and
// visible.
func (window Window) GetDisplayMode() (*DisplayMode, error) {
	mode := new(DisplayMode)
	r := int(C.SDL_GetWindowDisplayMode(window.ptr,
		(*C.SDL_DisplayMode)(unsafe.Pointer(mode))))
	if r != 0 {
		return nil, sdlError(r)
	}
	return mode, nil
}

// GetPixelFormat returns the pixel format of the Window.
func (window Window) GetPixelFormat() uint32 {
	return uint32(C.SDL_GetWindowPixelFormat(window.ptr))
}

// CreateWindow creates a window.  If you want the window to be centered set
// x and/or y to WINDOWPOS_CENTERED, if you don't care about the window position
// you can set x and/or y to WINDOWPOS_UNDEFINED.  On error CreateWindow returns
// nil.
func CreateWindow(title string, x, y, w, h int, flags WindowFlags) (Window, error) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	win := Window{}
	win.ptr = C.SDL_CreateWindow(ctitle, C.int(x), C.int(y), C.int(w),
		C.int(h), C.Uint32(flags))
	if win.ptr == nil {
		return win, sdlError(0)
	}
	return win, nil
}

// CreateWindowFrom creates and SDL window from an existing native window.
func CreateWindowFrom(data uintptr) (Window, error) {
	w := Window{}
	w.ptr = C.SDL_CreateWindowFrom(unsafe.Pointer(data))
	if w.ptr == nil {
		return w, sdlError(0)
	}
	return w, nil
}

// GetID returns the numeric ID of the window, for logging purposes.
func (window Window) GetID() uint32 {
	return uint32(C.SDL_GetWindowID(window.ptr))
}

// GetWindowFromID returns a window from a stored ID, or nil if it doesn't exist.
func GetWindowFromID(id uint32) (Window, error) {
	w := Window{}
	w.ptr = C.SDL_GetWindowFromID(C.Uint32(id))
	if w.ptr == nil {
		return w, sdlError(0)
	}
	return w, nil
}

// GetFlags returns the windows flags.
func (window Window) GetFlags() WindowFlags {
	return WindowFlags(C.SDL_GetWindowFlags(window.ptr))
}

// SetTitle sets the title of the window, in UTF-8 format.
func (window Window) SetTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.SDL_SetWindowTitle(window.ptr, ctitle)
}

// GetTitle gets the title of the window, in UTF-8 format.
func (window Window) GetTitle() string {
	return C.GoString(C.SDL_GetWindowTitle(window.ptr))
}

// SetIcon sets the icon for a window.
func (window Window) SetIcon(icon *Surface) {
	C.SDL_SetWindowIcon(window.ptr,
		(*C.SDL_Surface)(unsafe.Pointer(icon)))
}

/*
Skipping
SDL_SetWindowData()
SDL_GetWindowData()

I believe type embedding can be used for the same purpose in a type safe way.
*/

// SetPosition sets the position of the window.
func (window Window) SetPosition(x, y int) {
	C.SDL_SetWindowPosition(window.ptr,
		C.int(x), C.int(y))
}

// GetPosition returns the position of a window.
func (window Window) GetPosition() (x, y int) {
	C.SDL_GetWindowPosition(window.ptr,
		(*C.int)(unsafe.Pointer(&x)), (*C.int)(unsafe.Pointer(&y)))
	return
}

// SetSize sets the size of the window's client area. You can't change the size
// of a fullscreen window, it automatically matches the size of the display
// mode.
func (window Window) SetSize(w, h int) {
	C.SDL_SetWindowSize(window.ptr,
		C.int(w), C.int(h))
}

// GetSize returns the size of the window's client area.
func (window Window) GetSize() (w, h int) {
	C.SDL_GetWindowSize(window.ptr,
		(*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer(&h)))
	return
}

// SetMinimumSize sets the minimum size of window's client area.
//
// Note: You can't change the minimum size of a fullscreen window, it
// automatically matches the size of the displace mode.
func (window Window) SetMinimumSize(min_w, min_h int) {
	C.SDL_SetWindowMinimumSize(window.ptr, C.int(min_w), C.int(min_h))
}

// GetMinimumSize gets the minimum size of window's client area.
func (window Window) GetMinimumSize() (w, h int) {
	C.SDL_GetWindowMinimumSize(window.ptr,
		(*C.int)(unsafe.Pointer(&w)),
		(*C.int)(unsafe.Pointer(&h)))
	return
}

// SetMaximumSize sets the maximum size of a window's client area.
//
// Note: You can't change the maximum size of a fullscreen window, it
// automatically matches the size of the display mode.
func (window Window) SetMaximumSize(max_w, max_h int) {
	C.SDL_SetWindowMaximumSize(window.ptr, C.int(max_w), C.int(max_h))
}

// GetMaximumSize gets the maximum size of a window's client area.
func (window Window) GetMaximumSize() (w, h int) {
	cw, ch := new(C.int), new(C.int)
	C.SDL_GetWindowMaximumSize(window.ptr, cw, ch)
	return int(*cw), int(*ch)
}

// SetBordered will add or remove the window's WINDOW_BORDERLESS flag and add
// or remove the border from the actual window.  This is a no-op if the
// window's border already match the requested state.  If bordered is false
// the border will be removed, if it is true the border will be added.
//
// Note: You can't change the border state of a fullscreen window.
func (window Window) SetBordered(bordered bool) {
	b := C.SDL_FALSE
	if bordered {
		b = C.SDL_TRUE
	}
	C.SDL_SetWindowBordered(window.ptr, C.SDL_bool(b))
}

// Show shows the window.
func (window Window) Show() {
	C.SDL_ShowWindow(window.ptr)
}

// Hide hides the window.
func (window Window) Hide() {
	C.SDL_ShowWindow(window.ptr)
}

// Raise will raise the window above the other windows and set the input focus.
func (window Window) Raise() {
	C.SDL_RaiseWindow(window.ptr)
}

// Maximize makes the window as large as possible.
func (window Window) Maximize() {
	C.SDL_MaximizeWindow(window.ptr)
}

// Minimize minimizes the window to an iconic representation.
func (window Window) Minimize() {
	C.SDL_MinimizeWindow(window.ptr)
}

// Restore restores the size and position of a minimized or maximized window.
func (window Window) Restore() {
	C.SDL_RestoreWindow(window.ptr)
}

// SetFullscreen sets the windows fullscreen state.
func (window Window) SetFullscreen(flags uint32) error {
	r := int(C.SDL_SetWindowFullscreen(window.ptr,
		C.Uint32(flags)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetSurface gets the SDL surface associated with the window.  A new surface
// will be created with the optimal format for the window, if necessary.  This
// surface will be freed when the window is destroyed.
//
// Note: You may not combine this with 3D or the rendering API on this window.
func (window Window) GetSurface() (*Surface, error) {
	surface := C.SDL_GetWindowSurface(window.ptr)
	if surface == nil {
		return nil, sdlError(0)
	}
	return (*Surface)(unsafe.Pointer(surface)), nil
}

// UpdateSurface copies the window surface to the screen.
func (window Window) UpdateSurface() error {
	if r := int(C.SDL_UpdateWindowSurface(window.ptr)); r != 0 {
		return sdlError(r)
	}
	return nil
}

// UpdateSurfaceRects copies rectangles on the window surface to the screen.
func (window Window) UpdateSurfaceRects(rects []Rect) error {
	var ptr *C.SDL_Rect
	if len(rects) > 0 {
		ptr = (*C.SDL_Rect)(unsafe.Pointer(&rects[0]))
	}

	r := int(C.SDL_UpdateWindowSurfaceRects(window.ptr, ptr,
		C.int(len(rects))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// SetGrab sets the window's input grab mode. If grab is true input is grabbed,
// if it is false input is released.
func (window Window) SetGrab(grab bool) {
	var cgrab C.SDL_bool
	if grab {
		cgrab = C.SDL_TRUE
	} else {
		cgrab = C.SDL_FALSE
	}
	C.SDL_SetWindowGrab(window.ptr, cgrab)
}

// GetGrab gets the window's input grab mode.
func (window Window) GetGrab() bool {
	b := C.SDL_GetWindowGrab(window.ptr)
	if b == C.SDL_TRUE {
		return true
	}
	return false
}

// SetBrightness sets the window's brightness (gamma correction).
func (window Window) SetBrightness(brightness float32) error {
	r := int(C.SDL_SetWindowBrightness(window.ptr,
		C.float(brightness)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetWindowBrightness gets the window's brightness (gamma correction).
func (window Window) GetBrightness() float32 {
	return float32(C.SDL_GetWindowBrightness(window.ptr))
}

// SetGammaRamp sets the gamma translation table for the red, green, and blue
// channels of the video hardware.  Each table is an array of 256 16-bit
// quantities, representing a mapping between the input and output for that
// channel.  The input is the index into the array, and the output is the 16-bit
// gamma value at the index, scaled to the output color precision.  If you do
// not want to set a channel you can use nil instead.
func (window Window) SetGamaRamp(red, green, blue *[256]uint16) error {
	r := int(C.SDL_SetWindowGammaRamp(window.ptr,
		(*C.Uint16)(unsafe.Pointer(&red[0])),
		(*C.Uint16)(unsafe.Pointer(&green[0])),
		(*C.Uint16)(unsafe.Pointer(&blue[0]))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetGammaRamp gets the gamma ramp for a window.  If you do not want to get
// a channel you can use nil instead.
func (window Window) GetGammaRamp(red, green, blue *[256]uint16) error {
	r := int(C.SDL_GetWindowGammaRamp(window.ptr,
		(*C.Uint16)(unsafe.Pointer(&red[0])),
		(*C.Uint16)(unsafe.Pointer(&green[0])),
		(*C.Uint16)(unsafe.Pointer(&blue[0]))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// Destroy destroys the window.
func (window Window) Destroy() {
	C.SDL_DestroyWindow(window.ptr)
	window.ptr = nil
}

// IsScreenSaverEnabled returns whether the screensaver is currently enabled
// (default on).
func IsScreenSaverEnabled() bool {
	if b := C.SDL_IsScreenSaverEnabled(); b == C.SDL_TRUE {
		return true
	}
	return false
}

// EnableScreenSaver allows the screen to be blanked by a screensaver.
func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

// DisableScreenSaver prevents the screen from being blanked by a screensaver.
func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

/*
Skipping
SDL_GL_LoadLibrary()
SDL_GL_GetProcAddress()
SDL_GL_UnloadLibrary()
SDL_GL_ExtensionSupported

I think you only need these if you are going to use c function pointers.
*/

// GL_SetAttribute sets and OpenGL window attribute before creation.
func GL_SetAttribute(attr GLattr, value int) error {
	if r := int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value))); r != 0 {
		return sdlError(r)
	}
	return nil
}

// GL_GetAttribute gets the actual value for an attribute from the current context.
func GL_GetAttribute(attr GLattr) (int, error) {
	var value int
	r := int(C.SDL_GL_GetAttribute(C.SDL_GLattr(attr), (*C.int)(unsafe.Pointer(&value))))
	if r != 0 {
		return value, sdlError(r)
	}
	return value, nil
}

// GL_CreateContext creates an OpenGL context for use with an OpenGL window,
// and makes it current.
func (window Window) GL_CreateContext() (GLContext, error) {
	g := GLContext{}
	g.ctx = C.SDL_GL_CreateContext(window.ptr)
	if g.ctx == nil {
		return g, sdlError(0)
	}
	return g, nil
}

// GL_MakeCurrent sets up an OpenGL context for rendering into an OpenGL window.
//
// Note: The context must have been created with a compatible window.
func (window Window) GL_MakeCurrent(context GLContext) error {
	r := int(C.SDL_GL_MakeCurrent(window.ptr,
		context.ctx))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GL_SetSwapInterval sets the swap interval for the current OpenGL context.
//
//  interval = 0  -> immediate updates
//  interval = 1  -> updates synchronized with the vertical retrace
//  interval = -1 -> allow late swaps to happen immediately instead of waiting for the next retrace.
func GL_SetSwapInterval(interval int) error {
	r := int(C.SDL_GL_SetSwapInterval(C.int(interval)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GL_GetSwapInterval gets the swap interval for the current OpenGL context.
// It returns 0 if there is no vertical retrace synchronization, 1 if the
// buffer swap is synchronized with the vertical retrace, and -1 if late swaps
// happen immediately instead of waiting for the next retrace.  If the system
// can't determine the swap interval, or there isn't a valid current context,
// 0 will be returned as a safe default.
func GL_GetSwapInterval() int {
	return int(C.SDL_GL_GetSwapInterval())
}

// GL_Swap swaps the OpenGL buffers for a window, if double-buffering is supported.
func (window Window) GL_Swap() {
	C.SDL_GL_SwapWindow(window.ptr)
}

// Delete deletes the OpenGL context.
func (context GLContext) Delete() {
	C.SDL_GL_DeleteContext(context.ctx)
	context.ctx = nil
}

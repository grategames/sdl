// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"
import "unsafe"

type MessageBoxFlags uint32

const (
	// error dialog
	MESSAGEBOX_ERROR MessageBoxFlags = C.SDL_MESSAGEBOX_ERROR
	// warning dialog
	MESSAGEBOX_WARNING MessageBoxFlags = C.SDL_MESSAGEBOX_WARNING
	// informational dialog
	MESSAGEBOX_INFORMATION MessageBoxFlags = C.SDL_MESSAGEBOX_INFORMATION
)

type MessageBoxButtonFlags uint32

const (
	// Marks the default button when return is hit
	MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT MessageBoxButtonFlags = C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT
	// Marks the default button when escape is hit
	MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT MessageBoxButtonFlags = C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
)

type MessageBoxColorType uint32

const (
	MESSAGEBOX_COLOR_BACKGROUND        MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BACKGROUND
	MESSAGEBOX_COLOR_TEXT              MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_TEXT
	MESSAGEBOX_COLOR_BUTTON_BORDER     MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_BORDER
	MESSAGEBOX_COLOR_BUTTON_BACKGROUND MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND
	MESSAGEBOX_COLOR_BUTTON_SELECTED   MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED
)

// MessageBoxButtonData contains the data for a button in a message box.
type MessageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	Buttonid int32  // User defined button id
	Text     string // Button text (UTF-8)
}

// MessageBoxColor contains the RGB value used in message box color scheme.
type MessageBoxColor struct {
	R uint8
	G uint8
	B uint8
}

// MessageBoxColorScheme is a set of colors to use for message box dialogs.
type MessageBoxColorScheme [5]MessageBoxColor

// MessageBoxData contains all the information needed to create a message box
// dialog.
type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      Window // The parent window, or a zero value Window for no parent.
	Title       string // Title text (UTF-8)
	Message     string // Mesage text (UTF-8)
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme // Can be nil
}

// Show creates a modal messagebox, buttonid is the id of the button pressed
// by the user.
//
// Note: Show should be called on the thread that created the parent winow, or
// on the main thread if the messagebox has no parent.  It will block
// execution of that thread until the user clicks a button or closes the
// messagebox.
func (box *MessageBoxData) Show() (buttonid int32, err error) {
	cBox := C.SDL_MessageBoxData{}

	cBox.flags = C.Uint32(box.Flags)
	cBox.window = box.Window.ptr

	cBox.title = C.CString(box.Title)
	defer C.free(unsafe.Pointer(cBox.title))

	cBox.message = C.CString(box.Message)
	defer C.free(unsafe.Pointer(cBox.message))

	cButtons := make([]C.SDL_MessageBoxButtonData, 0, len(box.Buttons))
	cButton := C.SDL_MessageBoxButtonData{}
	for _, button := range box.Buttons {
		cButton.flags = C.Uint32(button.Flags)
		cButton.buttonid = C.int(button.Buttonid)
		cButton.text = C.CString(button.Text)
		defer C.free(unsafe.Pointer(cButton.text))

		cButtons = append(cButtons, cButton)
	}
	cBox.buttons = &cButtons[0]
	cBox.numbuttons = C.int(len(cButtons))

	cColorScheme := C.SDL_MessageBoxColorScheme{}
	if box.ColorScheme == nil {
		cBox.colorScheme = nil
	} else {
		for k, v := range box.ColorScheme {
			cColorScheme.colors[k].r = C.Uint8(v.R)
			cColorScheme.colors[k].g = C.Uint8(v.G)
			cColorScheme.colors[k].b = C.Uint8(v.B)
		}
		cBox.colorScheme = &cColorScheme
	}

	r := C.SDL_ShowMessageBox(&cBox, (*C.int)(&buttonid))
	if r != 0 {
		err = sdlError(int(r))
	}
	return
}

// ShowSimpleMessageBox creates a simple modal message box.
func ShowSimpleMessageBox(flags MessageBoxFlags, title, message string, window *C.SDL_Window) error {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))

	r := C.SDL_ShowSimpleMessageBox(C.Uint32(flags), cTitle, cMessage, window)
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

// #include "SDL.h"
import "C"

// Window state change event data
type WindowEvent struct {
	Type      EventType // WINDOWEVENT
	Timestamp uint32
	WindowID  uint32 // The asscociated window
	Event     WindowEventID
	_         uint8
	_         uint8
	_         uint8
	Data1     int32 // Event dependent data
	Data2     int32 // Event dependent data
}

// Keyboard button event structure
type KeyboardEvent struct {
	Type      EventType // KEYDOWN OR KEYUP
	Timestamp uint32
	WindowID  uint32 // The window with keyboard focus, if any
	State     uint8  // PRESSED or RELEASED
	Repeat    uint8  // Non-zero if this is a key repeat
	_         uint8
	_         uint8
	Keysym    Keysym // The key that was pressed or released
}

// Keyboard text editing event structure
type TextEditingEvent struct {
	Type      EventType // TEXTEDITING
	Timestamp uint32
	WindowID  uint32   // The window with keyboard focus, if any
	Text      [32]int8 // The editing text
	Start     int32    // The start cursor of selected editing text
	Length    int32    // The length of selected editing text
}

// Keyboard text input event structure
type TextInputEvent struct {
	Type      EventType // TEXTINPUT
	Timestamp uint32
	WindowID  uint32   // The window with keyboard focus, if any
	Text      [32]int8 // The input text
}

// Mouse motion event structure
type MouseMotionEvent struct {
	Type      EventType // MOUSEMOTION
	Timestamp uint32
	WindowID  uint32 // The window with mouse focus, if any
	Which     uint32
	State     uint8 // The current button state
	_         uint8
	_         uint8
	_         uint8
	X         int32 // X coordinate, relative to window
	Y         int32 // Y coordinate, relative to window
	Xrel      int32 // The relative motion in the X direction
	Yrel      int32 // The relative motion in the Y direction
}

// Mouse button event structure
type MouseButtonEvent struct {
	Type      EventType // MOUSEBUTTONDOWN or MOUSEBUTTONUP
	Timestamp uint32
	WindowID  uint32 // The window with mouse focus, if any
	Which     uint32
	Button    uint8 // The mouse button index
	State     uint8 // PRESSED or RELEASED
	_         uint8
	_         uint8
	X         int32 // X coordinate, relative to window
	Y         int32 // Y coordinate, relative to window
}

// Mouse wheel event structure
type MouseWheelEvent struct {
	Type      EventType // MOUSEWHEEL
	Timestamp uint32
	WindowID  uint32 // The window with mouse focus, if any
	Which     uint32
	X         int32 // The amount scrolled horizontally
	Y         int32 // The amount scrolled vertically
}

// Joystick axis motion event structure
type JoyAxisEvent struct {
	Type      EventType // JOYAXISMOTION
	Timestamp uint32
	Which     int32 // The joystick device index
	Axis      uint8 // The joystick axis index
	_         uint8
	_         uint8
	_         uint8
	Value     int16 // The axis value (range: -32768 to 32767)
	_         uint16
}

// Joystick ball event structure
type JoyBallEvent struct {
	Type      EventType // JOYBALLMOTION
	Timestamp uint32
	Which     int32 // The joystick device index
	Ball      uint8 // The joystick trackball index
	_         uint8
	_         uint8
	_         uint8
	Xrel      int16 // The relative motion in the X direction
	Yrel      int16 // The relative motion in the Y direction
}

// Joystick hat postion change event structure
type JoyHatEvent struct {
	Type      EventType // JOYHATMOTION
	Timestamp uint32
	Which     int32 // The joystick device index
	Hat       uint8 // The joystick hat index
	// The hat position value.
	// HAT_LEFTUP, HAT_UP, HAT_RIGHTUP
	// HAT_LEFT, HAT_CENTER, HAT_RIGHT
	// HAT_LEFTDOWN, HAT_DOWN, HAT_RIGHTDOWN
	//
	// Note: zero means the POV is centered.
	Value uint8
	_     uint8
	_     uint8
}

// Joystick button event structure
type JoyButtonEvent struct {
	Type      EventType // JOYBUTTONDOWN or JOYBUTTONUP
	Timestamp uint32
	Which     int32 // The joystick device index
	Button    uint8 // The joystick button index
	State     uint8 // PRESSED or RELEASED
	_         uint8
	_         uint8
}

// Joystick device event structure
type JoyDeviceEvent struct {
	Type      EventType // JOYDEVICEADDED or JOYDEVICEREMOVED
	Timestamp uint32
	Which     int32 // The joystick device index for ADD, instance_id for REMOVE
}

// Game controller axis motion event structure
type ControllerAxisEvent struct {
	Type      EventType // CONTROLLERAXISMOTION
	Timestamp uint32
	Which     int32 // The joystick instance id
	Axis      uint8 // The joystick axis index
	_         uint8
	_         uint8
	_         uint8
	Value     int16 // The axis value (range: -32768 to 32767)
	_         uint16
}

// Game controller button event structure
type ControllerButtonEvent struct {
	Type      EventType // CONTROLLERBUTTONDOWN or CONTROLLERBUTTONUP
	Timestamp uint32
	Which     int32 // The joystick instance id
	Button    uint8 // The joystick button index
	State     uint8 // PRESSED or RELEASED
	_         uint8
	_         uint8
}

// Controller device event structure
type ControllerDeviceEvent struct {
	Type      EventType // CONTROLLERDEVICEADDED or CONTROLLERDEVICEREMOVED
	Timestamp uint32
	Which     int32 // The joystick device index for ADD, instance_id for REMOVE
}

// Touch finger motion/finger event structure
type TouchFingerEvent struct {
	Type      EventType // FINGERMOTION or FINGERDOWN or FINGERUP
	Timestamp uint32
	TouchId   int64 // The touch device id
	FingerId  int64
	X         float32
	Y         float32
	Dx        float32
	Dy        float32
	Pressure  float32
	_         [4]byte
}

// Multiple Finger Gesture Event
type MultiGestureEvent struct {
	Type       EventType // MULTIGESTURE
	Timestamp  uint32
	TouchId    int64 // The touch device index
	DTheta     float32
	DDist      float32
	X          float32 // currently 0...1.
	Y          float32
	NumFingers uint16
	_          uint16
	_          [4]byte
}

type DollarGestureEvent struct {
	Type       EventType // DOLLARGESTURE
	Timestamp  uint32
	TouchId    int64 // The touch device index
	GestureId  int64
	NumFingers uint32
	Error      float32
	X          float32
	Y          float32
}

// An event used to request a file open by the system.  This event is disabled
// by default, you can enable it with EventState.
//
// If you enable this event, you must call FreeFile for all DropEvents.
type DropEvent struct {
	Type      EventType // DROPFILE
	Timestamp uint32
	file      *C.char
}

// The "quit requested" event
type QuitEvent struct {
	Type      EventType // QUIT
	Timestamp uint32
}

// A user-defined event type
type UserEvent struct {
	Type      EventType // USEREVENT through LASTEVENT
	Timestamp uint32
	WindowID  uint32  // The associated window if any
	Code      int32   // User defined event code
	Data1     uintptr // User defined data pointer
	Data2     uintptr // User defined data pointer
}

// A video driver dependent system event.  This event is disabled by default,
// you can enable it with EventState.
type SysWMEvent struct {
	Type      EventType // SYSWMEVENT
	Timestamp uint32
	Msg       *SysWMmsg //driver dependent data
}

// EventUnion is used to hold a SDL_Event and is used for transferring events
// between go and C.  To gain access to the other fields of an event you
// will have to call the Convert method.
type EventUnion struct {
	Type EventType
	_    [52]byte
}

// JoystickGUID encodes the stable unique id for a joystick device.
type JoystickGUID struct {
	Data [16]uint8
}

// Keysym is used in key events.
type Keysym struct {
	Scancode Scancode // SDL physical key code - see Scancode for details
	Sym      Keycode  // SDL virtual key code - see Keycode for details
	Mod      Keymod   // current key modifiers
	_        [2]byte
	unicode  uint32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// Palette contains palette information.  You should never need to create a
// palette manually.  It is automatically created when SDL allocates a
// PixelFormat for a surface.
type Palette struct {
	ncolors  int32
	colors   uintptr
	version  uint32
	refcount int32
}

// Everything in PixelFormat is read-only.
type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	_             [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	refcount      int32
	next          *PixelFormat
}

// Point defines a point
type Point struct {
	X int32
	Y int32
}

// Rect is a rectangle, with the origin at the upper left.
type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

// Surface is a collection of pixels used in software blitting.
//
// Note: This structure should be treated as read-only, except for Pixels,
// which, if not nil, contains the raw pixel data for the surface.
type Surface struct {
	flags  uint32
	Format *PixelFormat // Read-only
	W      int32        // Read-only
	H      int32        // Read-only

	// The length of a row pixels in bytes
	Pitch int32 // Read-only

	pixels uintptr

	// Application data associated with the surface
	Userdata uintptr // Read-write

	locked    int32
	lock_data uintptr

	// Clipping information
	Clip_rect Rect // Read-only

	_map     uintptr
	Refcount int32 // Read-mostly
}

// Version defines a SDL version.
type Version struct {
	Major uint8 // major version
	Minor uint8 // minor version
	Patch uint8 // update version
}

// DisplayMode defines a display mode.
type DisplayMode struct {
	Format      uint32 // pixel format
	W           int32  // width
	H           int32  // height
	RefreshRate int32  // refresh rate (or zero for unspecified)
	Driverdata  *byte  // driver-specific data, initialize to 0
}

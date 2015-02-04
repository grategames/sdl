// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

// The SDL keyboard scancode representation.
//
// Values of this type are used to represent keyboard keys.  The values of
// this type are based on the USB usage page standard:
// http://www.usb.org/developers/devclass_docs/Hut1_12.pdf
type Scancode uint32

const (
	SCANCODE_UNKNOWN Scancode = C.SDL_SCANCODE_UNKNOWN

	SCANCODE_A Scancode = C.SDL_SCANCODE_A
	SCANCODE_B Scancode = C.SDL_SCANCODE_B
	SCANCODE_C Scancode = C.SDL_SCANCODE_C
	SCANCODE_D Scancode = C.SDL_SCANCODE_D
	SCANCODE_E Scancode = C.SDL_SCANCODE_E
	SCANCODE_F Scancode = C.SDL_SCANCODE_F
	SCANCODE_G Scancode = C.SDL_SCANCODE_G
	SCANCODE_H Scancode = C.SDL_SCANCODE_H
	SCANCODE_I Scancode = C.SDL_SCANCODE_I
	SCANCODE_J Scancode = C.SDL_SCANCODE_J
	SCANCODE_K Scancode = C.SDL_SCANCODE_K
	SCANCODE_L Scancode = C.SDL_SCANCODE_L
	SCANCODE_M Scancode = C.SDL_SCANCODE_M
	SCANCODE_N Scancode = C.SDL_SCANCODE_N
	SCANCODE_O Scancode = C.SDL_SCANCODE_O
	SCANCODE_P Scancode = C.SDL_SCANCODE_P
	SCANCODE_Q Scancode = C.SDL_SCANCODE_Q
	SCANCODE_R Scancode = C.SDL_SCANCODE_R
	SCANCODE_S Scancode = C.SDL_SCANCODE_S
	SCANCODE_T Scancode = C.SDL_SCANCODE_T
	SCANCODE_U Scancode = C.SDL_SCANCODE_U
	SCANCODE_V Scancode = C.SDL_SCANCODE_V
	SCANCODE_W Scancode = C.SDL_SCANCODE_W
	SCANCODE_X Scancode = C.SDL_SCANCODE_X
	SCANCODE_Y Scancode = C.SDL_SCANCODE_Y
	SCANCODE_Z Scancode = C.SDL_SCANCODE_Z

	SCANCODE_1 Scancode = C.SDL_SCANCODE_1
	SCANCODE_2 Scancode = C.SDL_SCANCODE_2
	SCANCODE_3 Scancode = C.SDL_SCANCODE_3
	SCANCODE_4 Scancode = C.SDL_SCANCODE_4
	SCANCODE_5 Scancode = C.SDL_SCANCODE_5
	SCANCODE_6 Scancode = C.SDL_SCANCODE_6
	SCANCODE_7 Scancode = C.SDL_SCANCODE_7
	SCANCODE_8 Scancode = C.SDL_SCANCODE_8
	SCANCODE_9 Scancode = C.SDL_SCANCODE_9
	SCANCODE_0 Scancode = C.SDL_SCANCODE_0

	SCANCODE_RETURN    Scancode = C.SDL_SCANCODE_RETURN
	SCANCODE_ESCAPE    Scancode = C.SDL_SCANCODE_ESCAPE
	SCANCODE_BACKSPACE Scancode = C.SDL_SCANCODE_BACKSPACE
	SCANCODE_TAB       Scancode = C.SDL_SCANCODE_TAB
	SCANCODE_SPACE     Scancode = C.SDL_SCANCODE_SPACE

	SCANCODE_MINUS        Scancode = C.SDL_SCANCODE_MINUS
	SCANCODE_EQUALS       Scancode = C.SDL_SCANCODE_EQUALS
	SCANCODE_LEFTBRACKET  Scancode = C.SDL_SCANCODE_LEFTBRACKET
	SCANCODE_RIGHTBRACKET Scancode = C.SDL_SCANCODE_RIGHTBRACKET
	SCANCODE_BACKSLASH    Scancode = C.SDL_SCANCODE_BACKSLASH
	SCANCODE_NONUSHASH    Scancode = C.SDL_SCANCODE_NONUSHASH
	SCANCODE_SEMICOLON    Scancode = C.SDL_SCANCODE_SEMICOLON
	SCANCODE_APOSTROPHE   Scancode = C.SDL_SCANCODE_APOSTROPHE
	SCANCODE_GRAVE        Scancode = C.SDL_SCANCODE_GRAVE
	SCANCODE_COMMA        Scancode = C.SDL_SCANCODE_COMMA
	SCANCODE_PERIOD       Scancode = C.SDL_SCANCODE_PERIOD
	SCANCODE_SLASH        Scancode = C.SDL_SCANCODE_SLASH

	SCANCODE_CAPSLOCK Scancode = C.SDL_SCANCODE_CAPSLOCK

	SCANCODE_F1  Scancode = C.SDL_SCANCODE_F1
	SCANCODE_F2  Scancode = C.SDL_SCANCODE_F2
	SCANCODE_F3  Scancode = C.SDL_SCANCODE_F3
	SCANCODE_F4  Scancode = C.SDL_SCANCODE_F4
	SCANCODE_F5  Scancode = C.SDL_SCANCODE_F5
	SCANCODE_F6  Scancode = C.SDL_SCANCODE_F6
	SCANCODE_F7  Scancode = C.SDL_SCANCODE_F7
	SCANCODE_F8  Scancode = C.SDL_SCANCODE_F8
	SCANCODE_F9  Scancode = C.SDL_SCANCODE_F9
	SCANCODE_F10 Scancode = C.SDL_SCANCODE_F10
	SCANCODE_F11 Scancode = C.SDL_SCANCODE_F11
	SCANCODE_F12 Scancode = C.SDL_SCANCODE_F12

	SCANCODE_PRINTSCREEN Scancode = C.SDL_SCANCODE_PRINTSCREEN
	SCANCODE_SCROLLLOCK  Scancode = C.SDL_SCANCODE_SCROLLLOCK
	SCANCODE_PAUSE       Scancode = C.SDL_SCANCODE_PAUSE
	SCANCODE_INSERT      Scancode = C.SDL_SCANCODE_INSERT
	SCANCODE_HOME        Scancode = C.SDL_SCANCODE_HOME
	SCANCODE_PAGEUP      Scancode = C.SDL_SCANCODE_PAGEUP
	SCANCODE_DELETE      Scancode = C.SDL_SCANCODE_DELETE
	SCANCODE_END         Scancode = C.SDL_SCANCODE_END
	SCANCODE_PAGEDOWN    Scancode = C.SDL_SCANCODE_PAGEDOWN
	SCANCODE_RIGHT       Scancode = C.SDL_SCANCODE_RIGHT
	SCANCODE_LEFT        Scancode = C.SDL_SCANCODE_LEFT
	SCANCODE_DOWN        Scancode = C.SDL_SCANCODE_DOWN
	SCANCODE_UP          Scancode = C.SDL_SCANCODE_UP

	SCANCODE_NUMLOCKCLEAR Scancode = C.SDL_SCANCODE_NUMLOCKCLEAR
	SCANCODE_KP_DIVIDE    Scancode = C.SDL_SCANCODE_KP_DIVIDE
	SCANCODE_KP_MULTIPLY  Scancode = C.SDL_SCANCODE_KP_MULTIPLY
	SCANCODE_KP_MINUS     Scancode = C.SDL_SCANCODE_KP_MINUS
	SCANCODE_KP_PLUS      Scancode = C.SDL_SCANCODE_KP_PLUS
	SCANCODE_KP_ENTER     Scancode = C.SDL_SCANCODE_KP_ENTER
	SCANCODE_KP_1         Scancode = C.SDL_SCANCODE_KP_1
	SCANCODE_KP_2         Scancode = C.SDL_SCANCODE_KP_2
	SCANCODE_KP_3         Scancode = C.SDL_SCANCODE_KP_3
	SCANCODE_KP_4         Scancode = C.SDL_SCANCODE_KP_4
	SCANCODE_KP_5         Scancode = C.SDL_SCANCODE_KP_5
	SCANCODE_KP_6         Scancode = C.SDL_SCANCODE_KP_6
	SCANCODE_KP_7         Scancode = C.SDL_SCANCODE_KP_7
	SCANCODE_KP_8         Scancode = C.SDL_SCANCODE_KP_8
	SCANCODE_KP_9         Scancode = C.SDL_SCANCODE_KP_9
	SCANCODE_KP_0         Scancode = C.SDL_SCANCODE_KP_0
	SCANCODE_KP_PERIOD    Scancode = C.SDL_SCANCODE_KP_PERIOD

	SCANCODE_NONUSBACKSLASH Scancode = C.SDL_SCANCODE_NONUSBACKSLASH
	SCANCODE_APPLICATION    Scancode = C.SDL_SCANCODE_APPLICATION
	SCANCODE_POWER          Scancode = C.SDL_SCANCODE_POWER
	SCANCODE_KP_EQUALS      Scancode = C.SDL_SCANCODE_KP_EQUALS
	SCANCODE_F13            Scancode = C.SDL_SCANCODE_F13
	SCANCODE_F14            Scancode = C.SDL_SCANCODE_F14
	SCANCODE_F15            Scancode = C.SDL_SCANCODE_F15
	SCANCODE_F16            Scancode = C.SDL_SCANCODE_F16
	SCANCODE_F17            Scancode = C.SDL_SCANCODE_F17
	SCANCODE_F18            Scancode = C.SDL_SCANCODE_F18
	SCANCODE_F19            Scancode = C.SDL_SCANCODE_F19
	SCANCODE_F20            Scancode = C.SDL_SCANCODE_F20
	SCANCODE_F21            Scancode = C.SDL_SCANCODE_F21
	SCANCODE_F22            Scancode = C.SDL_SCANCODE_F22
	SCANCODE_F23            Scancode = C.SDL_SCANCODE_F23
	SCANCODE_F24            Scancode = C.SDL_SCANCODE_F24
	SCANCODE_EXECUTE        Scancode = C.SDL_SCANCODE_EXECUTE
	SCANCODE_HELP           Scancode = C.SDL_SCANCODE_HELP
	SCANCODE_MENU           Scancode = C.SDL_SCANCODE_MENU
	SCANCODE_SELECT         Scancode = C.SDL_SCANCODE_SELECT
	SCANCODE_STOP           Scancode = C.SDL_SCANCODE_STOP
	SCANCODE_AGAIN          Scancode = C.SDL_SCANCODE_AGAIN
	SCANCODE_UNDO           Scancode = C.SDL_SCANCODE_UNDO
	SCANCODE_CUT            Scancode = C.SDL_SCANCODE_CUT
	SCANCODE_COPY           Scancode = C.SDL_SCANCODE_COPY
	SCANCODE_PASTE          Scancode = C.SDL_SCANCODE_PASTE
	SCANCODE_FIND           Scancode = C.SDL_SCANCODE_FIND
	SCANCODE_MUTE           Scancode = C.SDL_SCANCODE_MUTE
	SCANCODE_VOLUMEUP       Scancode = C.SDL_SCANCODE_VOLUMEUP
	SCANCODE_VOLUMEDOWN     Scancode = C.SDL_SCANCODE_VOLUMEDOWN
	SCANCODE_KP_COMMA       Scancode = C.SDL_SCANCODE_KP_COMMA
	SCANCODE_KP_EQUALSAS400 Scancode = C.SDL_SCANCODE_KP_EQUALSAS400

	SCANCODE_INTERNATIONAL1 Scancode = C.SDL_SCANCODE_INTERNATIONAL1
	SCANCODE_INTERNATIONAL2 Scancode = C.SDL_SCANCODE_INTERNATIONAL2
	SCANCODE_INTERNATIONAL3 Scancode = C.SDL_SCANCODE_INTERNATIONAL3
	SCANCODE_INTERNATIONAL4 Scancode = C.SDL_SCANCODE_INTERNATIONAL4
	SCANCODE_INTERNATIONAL5 Scancode = C.SDL_SCANCODE_INTERNATIONAL5
	SCANCODE_INTERNATIONAL6 Scancode = C.SDL_SCANCODE_INTERNATIONAL6
	SCANCODE_INTERNATIONAL7 Scancode = C.SDL_SCANCODE_INTERNATIONAL7
	SCANCODE_INTERNATIONAL8 Scancode = C.SDL_SCANCODE_INTERNATIONAL8
	SCANCODE_INTERNATIONAL9 Scancode = C.SDL_SCANCODE_INTERNATIONAL9
	SCANCODE_LANG1          Scancode = C.SDL_SCANCODE_LANG1
	SCANCODE_LANG2          Scancode = C.SDL_SCANCODE_LANG2
	SCANCODE_LANG3          Scancode = C.SDL_SCANCODE_LANG3
	SCANCODE_LANG4          Scancode = C.SDL_SCANCODE_LANG4
	SCANCODE_LANG5          Scancode = C.SDL_SCANCODE_LANG5
	SCANCODE_LANG6          Scancode = C.SDL_SCANCODE_LANG6
	SCANCODE_LANG7          Scancode = C.SDL_SCANCODE_LANG7
	SCANCODE_LANG8          Scancode = C.SDL_SCANCODE_LANG8
	SCANCODE_LANG9          Scancode = C.SDL_SCANCODE_LANG9

	SCANCODE_ALTERASE   Scancode = C.SDL_SCANCODE_ALTERASE
	SCANCODE_SYSREQ     Scancode = C.SDL_SCANCODE_SYSREQ
	SCANCODE_CANCEL     Scancode = C.SDL_SCANCODE_CANCEL
	SCANCODE_CLEAR      Scancode = C.SDL_SCANCODE_CLEAR
	SCANCODE_PRIOR      Scancode = C.SDL_SCANCODE_PRIOR
	SCANCODE_RETURN2    Scancode = C.SDL_SCANCODE_RETURN2
	SCANCODE_SEPARATOR  Scancode = C.SDL_SCANCODE_SEPARATOR
	SCANCODE_OUT        Scancode = C.SDL_SCANCODE_OUT
	SCANCODE_OPER       Scancode = C.SDL_SCANCODE_OPER
	SCANCODE_CLEARAGAIN Scancode = C.SDL_SCANCODE_CLEARAGAIN
	SCANCODE_CRSEL      Scancode = C.SDL_SCANCODE_CRSEL
	SCANCODE_EXSEL      Scancode = C.SDL_SCANCODE_EXSEL

	SCANCODE_KP_00              Scancode = C.SDL_SCANCODE_KP_00
	SCANCODE_KP_000             Scancode = C.SDL_SCANCODE_KP_000
	SCANCODE_THOUSANDSSEPARATOR Scancode = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	SCANCODE_DECIMALSEPARATOR   Scancode = C.SDL_SCANCODE_DECIMALSEPARATOR
	SCANCODE_CURRENCYUNIT       Scancode = C.SDL_SCANCODE_CURRENCYUNIT
	SCANCODE_CURRENCYSUBUNIT    Scancode = C.SDL_SCANCODE_CURRENCYSUBUNIT
	SCANCODE_KP_LEFTPAREN       Scancode = C.SDL_SCANCODE_KP_LEFTPAREN
	SCANCODE_KP_RIGHTPAREN      Scancode = C.SDL_SCANCODE_KP_RIGHTPAREN
	SCANCODE_KP_LEFTBRACE       Scancode = C.SDL_SCANCODE_KP_LEFTBRACE
	SCANCODE_KP_RIGHTBRACE      Scancode = C.SDL_SCANCODE_KP_RIGHTBRACE
	SCANCODE_KP_TAB             Scancode = C.SDL_SCANCODE_KP_TAB
	SCANCODE_KP_BACKSPACE       Scancode = C.SDL_SCANCODE_KP_BACKSPACE
	SCANCODE_KP_A               Scancode = C.SDL_SCANCODE_KP_A
	SCANCODE_KP_B               Scancode = C.SDL_SCANCODE_KP_B
	SCANCODE_KP_C               Scancode = C.SDL_SCANCODE_KP_C
	SCANCODE_KP_D               Scancode = C.SDL_SCANCODE_KP_D
	SCANCODE_KP_E               Scancode = C.SDL_SCANCODE_KP_E
	SCANCODE_KP_F               Scancode = C.SDL_SCANCODE_KP_F
	SCANCODE_KP_XOR             Scancode = C.SDL_SCANCODE_KP_XOR
	SCANCODE_KP_POWER           Scancode = C.SDL_SCANCODE_KP_POWER
	SCANCODE_KP_PERCENT         Scancode = C.SDL_SCANCODE_KP_PERCENT
	SCANCODE_KP_LESS            Scancode = C.SDL_SCANCODE_KP_LESS
	SCANCODE_KP_GREATER         Scancode = C.SDL_SCANCODE_KP_GREATER
	SCANCODE_KP_AMPERSAND       Scancode = C.SDL_SCANCODE_KP_AMPERSAND
	SCANCODE_KP_DBLAMPERSAND    Scancode = C.SDL_SCANCODE_KP_DBLAMPERSAND
	SCANCODE_KP_VERTICALBAR     Scancode = C.SDL_SCANCODE_KP_VERTICALBAR
	SCANCODE_KP_DBLVERTICALBAR  Scancode = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	SCANCODE_KP_COLON           Scancode = C.SDL_SCANCODE_KP_COLON
	SCANCODE_KP_HASH            Scancode = C.SDL_SCANCODE_KP_HASH
	SCANCODE_KP_SPACE           Scancode = C.SDL_SCANCODE_KP_SPACE
	SCANCODE_KP_AT              Scancode = C.SDL_SCANCODE_KP_AT
	SCANCODE_KP_EXCLAM          Scancode = C.SDL_SCANCODE_KP_EXCLAM
	SCANCODE_KP_MEMSTORE        Scancode = C.SDL_SCANCODE_KP_MEMSTORE
	SCANCODE_KP_MEMRECALL       Scancode = C.SDL_SCANCODE_KP_MEMRECALL
	SCANCODE_KP_MEMCLEAR        Scancode = C.SDL_SCANCODE_KP_MEMCLEAR
	SCANCODE_KP_MEMADD          Scancode = C.SDL_SCANCODE_KP_MEMADD
	SCANCODE_KP_MEMSUBTRACT     Scancode = C.SDL_SCANCODE_KP_MEMSUBTRACT
	SCANCODE_KP_MEMMULTIPLY     Scancode = C.SDL_SCANCODE_KP_MEMMULTIPLY
	SCANCODE_KP_MEMDIVIDE       Scancode = C.SDL_SCANCODE_KP_MEMDIVIDE
	SCANCODE_KP_PLUSMINUS       Scancode = C.SDL_SCANCODE_KP_PLUSMINUS
	SCANCODE_KP_CLEAR           Scancode = C.SDL_SCANCODE_KP_CLEAR
	SCANCODE_KP_CLEARENTRY      Scancode = C.SDL_SCANCODE_KP_CLEARENTRY
	SCANCODE_KP_BINARY          Scancode = C.SDL_SCANCODE_KP_BINARY
	SCANCODE_KP_OCTAL           Scancode = C.SDL_SCANCODE_KP_OCTAL
	SCANCODE_KP_DECIMAL         Scancode = C.SDL_SCANCODE_KP_DECIMAL
	SCANCODE_KP_HEXADECIMAL     Scancode = C.SDL_SCANCODE_KP_HEXADECIMAL

	SCANCODE_LCTRL  Scancode = C.SDL_SCANCODE_LCTRL
	SCANCODE_LSHIFT Scancode = C.SDL_SCANCODE_LSHIFT
	SCANCODE_LALT   Scancode = C.SDL_SCANCODE_LALT
	SCANCODE_LGUI   Scancode = C.SDL_SCANCODE_LGUI
	SCANCODE_RCTRL  Scancode = C.SDL_SCANCODE_RCTRL
	SCANCODE_RSHIFT Scancode = C.SDL_SCANCODE_RSHIFT
	SCANCODE_RALT   Scancode = C.SDL_SCANCODE_RALT
	SCANCODE_RGUI   Scancode = C.SDL_SCANCODE_RGUI

	SCANCODE_MODE Scancode = C.SDL_SCANCODE_MODE

	SCANCODE_AUDIONEXT    Scancode = C.SDL_SCANCODE_AUDIONEXT
	SCANCODE_AUDIOPREV    Scancode = C.SDL_SCANCODE_AUDIOPREV
	SCANCODE_AUDIOSTOP    Scancode = C.SDL_SCANCODE_AUDIOSTOP
	SCANCODE_AUDIOPLAY    Scancode = C.SDL_SCANCODE_AUDIOPLAY
	SCANCODE_AUDIOMUTE    Scancode = C.SDL_SCANCODE_AUDIOMUTE
	SCANCODE_MEDIASELECT  Scancode = C.SDL_SCANCODE_MEDIASELECT
	SCANCODE_WWW          Scancode = C.SDL_SCANCODE_WWW
	SCANCODE_MAIL         Scancode = C.SDL_SCANCODE_MAIL
	SCANCODE_CALCULATOR   Scancode = C.SDL_SCANCODE_CALCULATOR
	SCANCODE_COMPUTER     Scancode = C.SDL_SCANCODE_COMPUTER
	SCANCODE_AC_SEARCH    Scancode = C.SDL_SCANCODE_AC_SEARCH
	SCANCODE_AC_HOME      Scancode = C.SDL_SCANCODE_AC_HOME
	SCANCODE_AC_BACK      Scancode = C.SDL_SCANCODE_AC_BACK
	SCANCODE_AC_FORWARD   Scancode = C.SDL_SCANCODE_AC_FORWARD
	SCANCODE_AC_STOP      Scancode = C.SDL_SCANCODE_AC_STOP
	SCANCODE_AC_REFRESH   Scancode = C.SDL_SCANCODE_AC_REFRESH
	SCANCODE_AC_BOOKMARKS Scancode = C.SDL_SCANCODE_AC_BOOKMARKS

	SCANCODE_BRIGHTNESSDOWN Scancode = C.SDL_SCANCODE_BRIGHTNESSDOWN
	SCANCODE_BRIGHTNESSUP   Scancode = C.SDL_SCANCODE_BRIGHTNESSUP
	SCANCODE_DISPLAYSWITCH  Scancode = C.SDL_SCANCODE_DISPLAYSWITCH
	SCANCODE_KBDILLUMTOGGLE Scancode = C.SDL_SCANCODE_KBDILLUMTOGGLE
	SCANCODE_KBDILLUMDOWN   Scancode = C.SDL_SCANCODE_KBDILLUMDOWN
	SCANCODE_KBDILLUMUP     Scancode = C.SDL_SCANCODE_KBDILLUMUP
	SCANCODE_EJECT          Scancode = C.SDL_SCANCODE_EJECT
	SCANCODE_SLEEP          Scancode = C.SDL_SCANCODE_SLEEP

	SCANCODE_APP1 Scancode = C.SDL_SCANCODE_APP1
	SCANCODE_APP2 Scancode = C.SDL_SCANCODE_APP2

	NUM_SCANCODES Scancode = C.SDL_NUM_SCANCODES
)
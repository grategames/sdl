// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unsafe"
)

func isPtr(k reflect.Kind) bool {
	if k == reflect.Ptr || k == reflect.UnsafePointer || k == reflect.Uintptr {
		return true
	}
	return false
}

func abiTest(i1, i2 interface{}) error {
	t1, t2 := reflect.TypeOf(i1), reflect.TypeOf(i2)

	if t1.Kind() != reflect.Struct {
		return fmt.Errorf("Type %s is not a struct.", t1.Name())
	}

	if t2.Kind() != reflect.Struct {
		return fmt.Errorf("Type %s is not a struct.", t2.Name())
	}

	if t1.Size() != t2.Size() {
		return fmt.Errorf("Types %s and %s are not the same size: %d %d.", t1.Name(), t2.Name(), t1.Size(), t2.Size())
	}

	t1Fields := make(map[string]string)
	for i := 0; i < t1.NumField(); i++ {
		field := t1.Field(i).Name
		if key, keep := filterField(field); keep {
			t1Fields[key] = field
		}
	}

	t2Fields := make(map[string]string)
	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i).Name
		if key, keep := filterField(field); keep {
			t2Fields[key] = field
		}
	}

	if len(t1Fields) != len(t2Fields) {
		t1fields := sort.StringSlice{}
		for k := range t1Fields {
			t1fields = append(t1fields, k)
		}
		t1fields.Sort()

		t2fields := sort.StringSlice{}
		for k := range t2Fields {
			t2fields = append(t2fields, k)
		}
		t2fields.Sort()

		return fmt.Errorf("Field count mismatch:\n%s fields: %s\n%s fields: %s", t1.Name(), t1fields, t2.Name(), t2fields)
	}

	for k := range t1Fields {
		if _, ok := t2Fields[k]; !ok {
			return fmt.Errorf("%s does not have field %s.", t2.Name(), k)
		}
		f1, ok := t1.FieldByName(t1Fields[k])
		if !ok {
			return fmt.Errorf("Failed to find field %s in $s.", t1Fields[k], t1.Name())
		}
		f2, ok := t2.FieldByName(t2Fields[k])
		if !ok {
			return fmt.Errorf("Failed to find field %s in %s.", t2Fields[k], t1.Name())
		}
		if f1.Offset != f2.Offset {
			return fmt.Errorf("Field %s (in %s and %s) has different offsets: %d, %d.", k, t1.Name(), t2.Name(), f1.Offset, f2.Offset)
		}

		kind1, kind2 := f1.Type.Kind(), f2.Type.Kind()
		if kind1 != kind2 {
			if !isPtr(kind1) || !isPtr(kind2) {
				return fmt.Errorf("Field %s (in %s and %s) has different kind: %s, %s.", k, t1.Name(), t2.Name(), kind1, kind2)
			}
		}

		if f1.Type.Size() != f2.Type.Size() {
			return fmt.Errorf("Field %s (in %s and %s) is not the same size: %d, %d.", k, t1.Name(), t2.Name(), f1.Type.Size(), f2.Type.Size())
		}
	}
	return nil
}

func filterField(field string) (string, bool) {
	if field == "_" || strings.HasPrefix(field, "pad") {
		return "", false
	}

	key := strings.Replace(field, "_", "", -1)
	key = strings.ToLower(key)

	return key, true
}

type iPair struct {
	i1 interface{}
	i2 interface{}
}

var structTests = []iPair{
	{WindowEvent{}, testWindowEvent{}},
	{KeyboardEvent{}, testKeyboardEvent{}},
	{TextEditingEvent{}, testTextEditingEvent{}},
	{TextInputEvent{}, testTextInputEvent{}},
	{MouseMotionEvent{}, testMouseMotionEvent{}},
	{MouseButtonEvent{}, testMouseButtonEvent{}},
	{MouseWheelEvent{}, testMouseWheelEvent{}},
	{JoyAxisEvent{}, testJoyAxisEvent{}},
	{JoyBallEvent{}, testJoyBallEvent{}},
	{JoyHatEvent{}, testJoyHatEvent{}},
	{JoyButtonEvent{}, testJoyButtonEvent{}},
	{JoyDeviceEvent{}, testJoyDeviceEvent{}},
	{ControllerAxisEvent{}, testControllerAxisEvent{}},
	{ControllerButtonEvent{}, testControllerButtonEvent{}},
	{ControllerDeviceEvent{}, testControllerDeviceEvent{}},
	{TouchFingerEvent{}, testTouchFingerEvent{}},
	{MultiGestureEvent{}, testMultiGestureEvent{}},
	{DollarGestureEvent{}, testDollarGestureEvent{}},
	{DropEvent{}, testDropEvent{}},
	{QuitEvent{}, testQuitEvent{}},
	{UserEvent{}, testUserEvent{}},
	{SysWMEvent{}, testSysWMEvent{}},
	{JoystickGUID{}, testJoystickGUID{}},
	{Keysym{}, testKeysym{}},
	{Color{}, testColor{}},
	{Palette{}, testPalette{}},
	{PixelFormat{}, testPixelFormat{}},
	{Point{}, testPoint{}},
	{Rect{}, testRect{}},
	{Surface{}, testSurface{}},
	{Version{}, testVersion{}},
	{DisplayMode{}, testDisplayMode{}},
}

func TestStructs(t *testing.T) {
	for _, pair := range structTests {
		err := abiTest(pair.i1, pair.i2)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}
}

func TestUnions(t *testing.T) {
	if int(unsafe.Sizeof(EventUnion{})) != len(testEventUnion{}) {
		t.Fatal("EventUnion not the same size as SDL_Event")
	}
}

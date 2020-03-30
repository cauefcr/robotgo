// Copyright 2016 The cauefcr Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/cauefcr/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package robotgo

import (
	Hook "github.com/cauefcr/ghook"
)

// const (
// 	HookEnabled  = 1 //iota
// 	HookDisabled = 2
// 	KeyDown      = 3
// 	KeyHold      = 4
// 	KeyUp        = 5
// 	MouseUp      = 6
// 	MouseHold    = 7
// 	MouseDown    = 8
// 	MouseMove    = 9
// 	MouseDrag    = 10
// 	MouseWheel   = 11
// 	FakeEvent    = 12
// 	//Keychar could be   v
// 	CharUndefined = 0xFFFF
// 	WheelUp       = -1
// 	WheelDown     = 1
// )

/*
 ___________    ____  _______ .__   __. .___________.
|   ____\   \  /   / |   ____||  \ |  | |           |
|  |__   \   \/   /  |  |__   |   \|  | `---|  |----`
|   __|   \      /   |   __|  |  . `  |     |  |
|  |____   \    /    |  |____ |  |\   |     |  |
|_______|   \__/     |_______||__| \__|     |__|
*/

// AddEvent add event listener,
//
// parameters for the string type,
// the keyboard corresponding key parameters,
//
// mouse arguments: mleft, center, mright, wheelDown, wheelUp,
// wheelLeft, wheelRight.
// func AddEvent(key string) bool {
// 	var (
// 		// cs   *C.char
// 		mArr = []string{"mleft", "center", "mright", "wheelDown",
// 			"wheelUp", "wheelLeft", "wheelRight"}
// 		mouseBool bool
// 	)

// 	for i := 0; i < len(mArr); i++ {
// 		if key == mArr[i] {
// 			mouseBool = true
// 		}
// 	}

// 	if len(key) > 1 && !mouseBool {
// 		key = strconv.Itoa(int(Keycode[key]))
// 	}

// 	geve := hook.AddEvent(key)
// 	// defer C.free(unsafe.Pointer(cs))

// 	if geve == 0 {
// 		return true
// 	}

// 	return false
// }

// StopEvent stop event listener
// func StopEvent() {
// 	hook.StopEvent()
// }

// Start start global event hook
// return event channel
func EventStart() chan Hook.Event {
	return Hook.Start()
}

// End removes global event hook
func EventEnd() {
	Hook.End()
}

func EventProcess(Events chan Hook.Event) chan bool {
	return Hook.Process(Events)
}

func EventHook(When uint8, keysPressed []string, Callback func(Hook.Event)) {
	Hook.Register(When, keysPressed, Callback)
	return
}

// AddEvents add global event hook
//
// robotgo.AddEvents("q")
// robotgo.AddEvents("q", "ctrl")
// robotgo.AddEvents("q", "ctrl", "shift")
// func AddEvents(key string, arr ...string) bool {
// 	s := hook.Start()
// 	// defer hook.End()

// 	ct := false
// 	k := 0
// 	for {
// 		e := <-s

// 		l := len(arr)
// 		if l > 0 {
// 			for i := 0; i < l; i++ {
// 				ukey := Keycode[arr[i]]

// 				if e.Kind == hook.KeyHold && e.Keycode == ukey {
// 					k++
// 				}

// 				if k == l {
// 					ct = true
// 				}

// 				if e.Kind == hook.KeyUp && e.Keycode == ukey {
// 					if k > 0 {
// 						k--
// 					}
// 					// time.Sleep(10 * time.Microsecond)
// 					ct = false
// 				}
// 			}
// 		} else {
// 			ct = true
// 		}

// 		if ct && e.Kind == hook.KeyUp && e.Keycode == Keycode[key] {
// 			hook.End()
// 			// k = 0
// 			break
// 		}
// 	}

// 	return true
// }

// AddMouse add mouse event hook
//
// mouse arguments: left, center, right, wheelDown, wheelUp,
// wheelLeft, wheelRight.
//
// robotgo.AddMouse("left")
// robotgo.AddMouse("left", 100, 100)
// func AddMouse(btn string, x ...int16) bool {
// 	s := hook.Start()
// 	ukey := MouseMap[btn]

// 	ct := false
// 	for {
// 		e := <-s

// 		if len(x) > 1 {
// 			if e.Kind == hook.MouseMove && e.X == x[0] && e.Y == x[1] {
// 				ct = true
// 			}
// 		} else {
// 			ct = true
// 		}

// 		if ct && e.Kind == hook.MouseDown && e.Button == ukey {
// 			hook.End()
// 			break
// 		}
// 	}

// 	return true
// }

// AddMousePos add listen mouse event pos hook
// func AddMousePos(x, y int16) bool {
// 	s := hook.Start()

// 	for {
// 		e := <-s
// 		if e.Kind == hook.MouseMove && e.X == x && e.Y == y {
// 			hook.End()
// 			break
// 		}
// 	}

// 	return true
// }

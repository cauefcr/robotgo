package robotgo

/*
#include "cdeps/hook/iohook.h"
*/
import "C"

//to-do: make it more like idiomatic go code
// several things are not working, like the mouse position
// import "strconv"

const(
	EVENT_TROWAWAY Event_type = iota
	EVENT_HOOK_ENABLED
	EVENT_HOOK_DISABLED
	EVENT_KEY_TYPED
	EVENT_KEY_PRESSED
	EVENT_KEY_RELEASED
	EVENT_MOUSE_CLICKED
	EVENT_MOUSE_PRESSED
	EVENT_MOUSE_RELEASED
	EVENT_MOUSE_MOVED
	EVENT_MOUSE_DRAGGED
	EVENT_MOUSE_WHEEL
)

type Keyboard_data struct{
	Keycode uint16
	Rawcode uint16
	Keychar rune
}

type Mouse_data struct{
	Button uint16
	Clicks uint16
	X int16
	Y int16
}

type Mouse_wheel_data struct{
	Clicks uint16
	Amount uint16
	Rotation int16
	X int16
	Y int16
	Kind uint8//can't be type
	Direction uint8
}

type Event_keyboard struct{
	Kind Event_type
	time uint64
	Mask uint16
	Reserved uint16
	data Keyboard_data
}

func (e Event_keyboard) Type() Event_type {
	return e.Kind
}

func (e Event_keyboard) Time() uint64 {
	return e.time
}

func (e Event_keyboard) Data() Keyboard_data {
	return e.data
}

type Event_mouse struct{
	Kind Event_type
	time uint64
	Mask uint16
	Reserved uint16
	data Mouse_data
}

func (e Event_mouse) Type() Event_type {
	return e.Kind
}

func (e Event_mouse) Time() uint64 {
	return e.time
}

func (e Event_mouse) Data() Mouse_data {
	return e.data
}

type Event_mouse_wheel struct{
	Kind Event_type
	time uint64
	Mask uint16
	Reserved uint16
	data Mouse_wheel_data
}

func (e Event_mouse_wheel) Type() Event_type {
	return e.Kind
}

func (e Event_mouse_wheel) Time() uint64 {
	return e.time
}

func (e Event_mouse_wheel) Data() Mouse_wheel_data {
	return e.data
}

//export send_key_event
func send_key_event(kev C.struct__event_keyboard){
	Event__ <- Event_keyboard{
		Kind: Event_type(C.int(kev._type)),
		time: uint64(C.uint64_t(kev.time)),
		Mask: uint16(C.uint16_t(kev.mask)),
		Reserved: uint16(C.uint16_t(kev.reserved)),
		data: Keyboard_data{
			Keycode: uint16(C.uint16_t(kev.data.keycode)),
			Rawcode: uint16(C.uint16_t(kev.data.rawcode)),
			Keychar: rune(C.uint16_t(kev.data.keychar)),
		},
	};
}
//export send_mouse_event
func send_mouse_event(mev C.struct__event_mouse){
	Event__ <- Event_mouse{
		Kind: Event_type(C.int(mev._type)),
		time: uint64(C.uint64_t(mev.time)),
		Mask: uint16(C.uint16_t(mev.mask)),
		Reserved: uint16(C.uint16_t(mev.reserved)),
		data: Mouse_data{
			Button: uint16(C.uint16_t(mev.data.button)),
			Clicks: uint16(C.uint16_t(mev.data.clicks)),
		},
	};
}

//export send_mouse_wheel_event
func send_mouse_wheel_event(mwe C.struct__event_mouse_wheel){
	Event__ <- Event_mouse_wheel{
		Kind: Event_type(C.int(mwe._type)),
		time: uint64(C.uint64_t(mwe.time)),
		Mask: uint16(C.uint16_t(mwe.mask)),
		Reserved: uint16(C.uint16_t(mwe.reserved)),
		data: Mouse_wheel_data{
			Clicks: uint16(C.uint16_t(mwe.data.clicks)),
			Amount: uint16(C.uint16_t(mwe.data.amount)),
			Rotation: int16(C.int16_t(mwe.data.rotation)),
			X: int16(C.int16_t(mwe.data.x)),
			Y: int16(C.int16_t(mwe.data.y)),
			Kind: uint8(C.uint8_t(mwe.data._type)),
			Direction: uint8(C.uint8_t(mwe.data.direction)),
		},
	};
}

package xel

import (
	"fmt"
	"github.com/amortaza/go-glfw"
	"github.com/amortaza/go-hal"
)

var gUserOnKey func(
	key hal.KeyboardKey,
	action hal.ButtonAction,
	alt,
	ctrl,
	shift bool)

func glfw_onKey(
		window *glfw.Window,
		key glfw.Key,
		scancode int,
		action glfw.Action,
		mods glfw.ModifierKey) {

	var _key hal.KeyboardKey
	var _action hal.ButtonAction

	if key == glfw.KeyA {
		_key = hal.Key_A

	} else if key == glfw.KeyB {
		_key = hal.Key_B

	} else if key == glfw.KeyC {
		_key = hal.Key_C

	} else if key == glfw.KeyD {
		_key = hal.Key_D

	} else if key == glfw.KeyE {
		_key = hal.Key_E

	} else if key == glfw.KeyF {
		_key = hal.Key_F

	} else if key == glfw.KeyG {
		_key = hal.Key_G

	} else if key == glfw.KeyH {
		_key = hal.Key_H

	} else if key == glfw.KeyI {
		_key = hal.Key_I

	} else if key == glfw.KeyJ {
		_key = hal.Key_J

	} else if key == glfw.KeyK {
		_key = hal.Key_K

	} else if key == glfw.KeyL {
		_key = hal.Key_L

	} else if key == glfw.KeyM {
		_key = hal.Key_M

	} else if key == glfw.KeyN {
		_key = hal.Key_N

	} else if key == glfw.KeyO {
		_key = hal.Key_O

	} else if key == glfw.KeyP {
		_key = hal.Key_P

	} else if key == glfw.KeyQ {
		_key = hal.Key_Q

	} else if key == glfw.KeyR {
		_key = hal.Key_R

	} else if key == glfw.KeyS {
		_key = hal.Key_S

	} else if key == glfw.KeyT {
		_key = hal.Key_T

	} else if key == glfw.KeyU {
		_key = hal.Key_U

	} else if key == glfw.KeyV {
		_key = hal.Key_V

	} else if key == glfw.KeyW {
		_key = hal.Key_W

	} else if key == glfw.KeyX {
		_key = hal.Key_X

	} else if key == glfw.KeyY {
		_key = hal.Key_Y

	} else if key == glfw.KeyZ {
		_key = hal.Key_Z

	} else if key == glfw.Key1 {
		_key = hal.Key_1

	} else if key == glfw.Key2 {
		_key = hal.Key_2

	} else if key == glfw.Key3 {
		_key = hal.Key_3

	} else if key == glfw.Key4 {
		_key = hal.Key_4

	} else if key == glfw.Key0 {
		_key = hal.Key_0

	} else if key == glfw.Key5 {
		_key = hal.Key_5

	} else if key == glfw.Key6 {
		_key = hal.Key_6

	} else if key == glfw.Key7 {
		_key = hal.Key_7

	} else if key == glfw.Key8 {
		_key = hal.Key_8

	} else if key == glfw.Key9 {
		_key = hal.Key_9

	} else if key == glfw.KeyKP1 {
		_key = hal.Key_PAD_1

	} else if key == glfw.KeyKP2 {
		_key = hal.Key_PAD_2

	} else if key == glfw.KeyKP3 {
		_key = hal.Key_PAD_3

	} else if key == glfw.KeyKP4 {
		_key = hal.Key_PAD_4

	} else if key == glfw.KeyKP0 {
		_key = hal.Key_PAD_0

	} else if key == glfw.KeyKP5 {
		_key = hal.Key_PAD_5

	} else if key == glfw.KeyKP6 {
		_key = hal.Key_PAD_6

	} else if key == glfw.KeyKP7 {
		_key = hal.Key_PAD_7

	} else if key == glfw.KeyKP8 {
		_key = hal.Key_PAD_8

	} else if key == glfw.KeyKP9 {
		_key = hal.Key_PAD_9

	} else if key == glfw.KeyF1 {
		_key = hal.Key_F1

	} else if key == glfw.KeyF2 {
		_key = hal.Key_F2

	} else if key == glfw.KeyF3 {
		_key = hal.Key_F3

	} else if key == glfw.KeyF4 {
		_key = hal.Key_F4

	} else if key == glfw.KeyF5 {
		_key = hal.Key_F5

	} else if key == glfw.KeyF6 {
		_key = hal.Key_F6

	} else if key == glfw.KeyF7 {
		_key = hal.Key_F7

	} else if key == glfw.KeyF8 {
		_key = hal.Key_F8

	} else if key == glfw.KeyF9 {
		_key = hal.Key_F9

	} else if key == glfw.KeyF10 {
		_key = hal.Key_F10

	} else if key == glfw.KeyF11 {
		_key = hal.Key_F11

	} else if key == glfw.KeyF12 {
		_key = hal.Key_F12

	} else if key == glfw.KeyKPDecimal {
		_key = hal.Key_PAD_DECIMAL

	} else if key == glfw.KeyKPDivide {
		_key = hal.Key_PAD_DIVIDE

	} else if key == glfw.KeyKPMultiply {
		_key = hal.Key_PAD_MULTIPLY

	} else if key == glfw.KeyKPSubtract {
		_key = hal.Key_PAD_SUBTRACT

	} else if key == glfw.KeyKPAdd {
		_key = hal.Key_PAD_ADD

	} else if key == glfw.KeyKPEnter {
		_key = hal.Key_PAD_ENTER

	} else if key == glfw.KeyKPEqual {
		_key = hal.Key_PAD_EQUAL

	} else if key == glfw.KeyRight {
		_key = hal.Key_RIGHT

	} else if key == glfw.KeyLeft {
		_key = hal.Key_LEFT

	} else if key == glfw.KeyUp {
		_key = hal.Key_UP

	} else if key == glfw.KeyDown {
		_key = hal.Key_DOWN

	} else if key == glfw.KeyPageUp {
		_key = hal.Key_PAGE_UP

	} else if key == glfw.KeyPageDown {
		_key = hal.Key_PAGE_DOWN

	} else if key == glfw.KeyHome {
		_key = hal.Key_HOME

	} else if key == glfw.KeyEnd {
		_key = hal.Key_END

	} else if key == glfw.KeyDelete {
		_key = hal.Key_DELETE

	} else if key == glfw.KeyInsert {
		_key = hal.Key_INSERT

	} else if key == glfw.KeyBackspace {
		_key = hal.Key_BACKSPACE

	} else if key == glfw.KeyTab {
		_key = hal.Key_TAB

	} else if key == glfw.KeyEnter {
		_key = hal.Key_ENTER

	} else if key == glfw.KeyEscape {
		_key = hal.Key_ESCAPE

	} else if key == glfw.KeyRightBracket {
		_key = hal.Key_RIGHT_BRACKET

	} else if key == glfw.KeyLeftBracket {
		_key = hal.Key_LEFT_BRACKET

	} else if key == glfw.KeyBackslash {
		_key = hal.Key_BACK_SLASH

	} else if key == glfw.KeyEqual {
		_key = hal.Key_EQUAL

	} else if key == glfw.KeySlash {
		_key = hal.Key_FORWARD_SLASH

	} else if key == glfw.KeySemicolon {
		_key = hal.Key_SEMICOLON

	} else if key == glfw.KeyPeriod {
		_key = hal.Key_PERIOD

	} else if key == glfw.KeyMinus {
		_key = hal.Key_MINUS

	} else if key == glfw.KeyComma {
		_key = hal.Key_COMMA

	} else if key == glfw.KeySpace {
		_key = hal.Key_SPACE

	} else if key == glfw.KeyApostrophe {
		_key = hal.Key_APOSTROPHE

	} else {
		fmt.Println("Unrecognized key pressed %i", key)
		return
	}

	if action == glfw.Press {
		_action = hal.Button_Action_Down

	} else if action == glfw.Release {
		_action = hal.Button_Action_Up

	} else if action == glfw.Repeat {
		_action = hal.Button_Action_Down

	} else {
		fmt.Println("Unrecognized key action %i in xel-keyboard.go", action)
		return
	}

	if gUserOnKey != nil {
		gUserOnKey(_key, _action, mods & glfw.ModAlt != 0, mods & glfw.ModControl != 0, mods & glfw.ModShift != 0)
	}
}


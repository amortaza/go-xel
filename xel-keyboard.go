package xel

import (
	"github.com/goxjs/glfw"
	"github.com/amortaza/go-bellina/constants"
	"fmt"
)

var _onKey func(key bl.KeyboardKey, action bl.ButtonAction, alt, ctrl, shift bool)

const (
	Key_APOSTROPHE bl.KeyboardKey = 1 + iota
	Key_SPACE
	Key_COMMA
	Key_MINUS
	Key_PERIOD
	Key_FORWARD_SLASH
	Key_SEMICOLON
	Key_EQUAL
	Key_LEFT_BRACKET
	Key_BACK_SLASH
	Key_RIGHT_BRACKET
	Key_ESCAPE
	Key_ENTER
	Key_TAB
	Key_BACKSPACE
	Key_INSERT
	Key_DELETE

	Key_RIGHT
	Key_LEFT
	Key_UP
	Key_DOWN
	Key_PAGE_UP
	Key_PAGE_DOWN
	Key_HOME
	Key_END

	Key_F1
	Key_F2
	Key_F3
	Key_F4
	Key_F5
	Key_F6
	Key_F7
	Key_F8
	Key_F9
	Key_F10
	Key_F11
	Key_F12

	Key_PAD_0
	Key_PAD_1
	Key_PAD_2
	Key_PAD_3
	Key_PAD_4
	Key_PAD_5
	Key_PAD_6
	Key_PAD_7
	Key_PAD_8
	Key_PAD_9

	Key_PAD_DECIMAL
	Key_PAD_DIVIDE
	Key_PAD_MULTIPLY
	Key_PAD_SUBTRACT
	Key_PAD_ADD
	Key_PAD_ENTER
	Key_PAD_EQUAL

	Key_1
	Key_2
	Key_3
	Key_4
	Key_5
	Key_6
	Key_7
	Key_8
	Key_9
	Key_0

	Key_A
	Key_B
	Key_C
	Key_D
	Key_E
	Key_F
	Key_G
	Key_H
	Key_I
	Key_J
	Key_K
	Key_L
	Key_M
	Key_N
	Key_O
	Key_P
	Key_Q
	Key_R
	Key_S
	Key_T
	Key_U
	Key_V
	Key_W
	Key_X
	Key_Y
	Key_Z
)

func __onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	var _key bl.KeyboardKey
	var _action bl.ButtonAction

	if key == glfw.KeyA {
		_key = Key_A
	} else if key == glfw.KeyB {
		_key = Key_B
	} else if key == glfw.KeyC {
		_key = Key_C
	} else if key == glfw.KeyD {
		_key = Key_D
	} else if key == glfw.KeyE {
		_key = Key_E
	} else if key == glfw.KeyF {
		_key = Key_F
	} else if key == glfw.KeyG {
		_key = Key_G
	} else if key == glfw.KeyH {
		_key = Key_H
	} else if key == glfw.KeyI {
		_key = Key_I
	} else if key == glfw.KeyJ {
		_key = Key_J
	} else if key == glfw.KeyK {
		_key = Key_K
	} else if key == glfw.KeyL {
		_key = Key_L
	} else if key == glfw.KeyM {
		_key = Key_M
	} else if key == glfw.KeyN {
		_key = Key_N
	} else if key == glfw.KeyO {
		_key = Key_O
	} else if key == glfw.KeyP {
		_key = Key_P
	} else if key == glfw.KeyQ {
		_key = Key_Q
	} else if key == glfw.KeyR {
		_key = Key_R
	} else if key == glfw.KeyS {
		_key = Key_S
	} else if key == glfw.KeyT {
		_key = Key_T
	} else if key == glfw.KeyU {
		_key = Key_U
	} else if key == glfw.KeyV {
		_key = Key_V
	} else if key == glfw.KeyW {
		_key = Key_W
	} else if key == glfw.KeyX {
		_key = Key_X
	} else if key == glfw.KeyY {
		_key = Key_Y
	} else if key == glfw.KeyZ {
		_key = Key_Z

	} else if key == glfw.Key1 {
		_key = Key_1
	} else if key == glfw.Key2 {
		_key = Key_2
	} else if key == glfw.Key3 {
		_key = Key_3
	} else if key == glfw.Key4 {
		_key = Key_4
	} else if key == glfw.Key0 {
		_key = Key_0
	} else if key == glfw.Key5 {
		_key = Key_5
	} else if key == glfw.Key6 {
		_key = Key_6
	} else if key == glfw.Key7 {
		_key = Key_7
	} else if key == glfw.Key8 {
		_key = Key_8
	} else if key == glfw.Key9 {
		_key = Key_9

	} else if key == glfw.KeyKP1 {
		_key = Key_PAD_1
	} else if key == glfw.KeyKP2 {
		_key = Key_PAD_2
	} else if key == glfw.KeyKP3 {
		_key = Key_PAD_3
	} else if key == glfw.KeyKP4 {
		_key = Key_PAD_4
	} else if key == glfw.KeyKP0 {
		_key = Key_PAD_0
	} else if key == glfw.KeyKP5 {
		_key = Key_PAD_5
	} else if key == glfw.KeyKP6 {
		_key = Key_PAD_6
	} else if key == glfw.KeyKP7 {
		_key = Key_PAD_7
	} else if key == glfw.KeyKP8 {
		_key = Key_PAD_8
	} else if key == glfw.KeyKP9 {
		_key = Key_PAD_9

	} else if key == glfw.KeyF1 {
		_key = Key_F1
	} else if key == glfw.KeyF2 {
		_key = Key_F2
	} else if key == glfw.KeyF3 {
		_key = Key_F3
	} else if key == glfw.KeyF4 {
		_key = Key_F4
	} else if key == glfw.KeyF5 {
		_key = Key_F5
	} else if key == glfw.KeyF6 {
		_key = Key_F6
	} else if key == glfw.KeyF7 {
		_key = Key_F7
	} else if key == glfw.KeyF8 {
		_key = Key_F8
	} else if key == glfw.KeyF9 {
		_key = Key_F9
	} else if key == glfw.KeyF10 {
		_key = Key_F10
	} else if key == glfw.KeyF11 {
		_key = Key_F11
	} else if key == glfw.KeyF12 {
		_key = Key_F12

	} else if key == glfw.KeyKPDecimal {
		_key = Key_PAD_DECIMAL
	} else if key == glfw.KeyKPDivide {
		_key = Key_PAD_DIVIDE
	} else if key == glfw.KeyKPMultiply {
		_key = Key_PAD_MULTIPLY
	} else if key == glfw.KeyKPSubtract {
		_key = Key_PAD_SUBTRACT
	} else if key == glfw.KeyKPAdd {
		_key = Key_PAD_ADD
	} else if key == glfw.KeyKPEnter {
		_key = Key_PAD_ENTER
	} else if key == glfw.KeyKPEqual {
		_key = Key_PAD_EQUAL

	} else if key == glfw.KeyRight {
		_key = Key_RIGHT
	} else if key == glfw.KeyLeft {
		_key = Key_LEFT
	} else if key == glfw.KeyUp {
		_key = Key_UP
	} else if key == glfw.KeyDown {
		_key = Key_DOWN
	} else if key == glfw.KeyPageUp {
		_key = Key_PAGE_UP
	} else if key == glfw.KeyPageDown {
		_key = Key_PAGE_DOWN
	} else if key == glfw.KeyHome {
		_key = Key_HOME
	} else if key == glfw.KeyEnd {
		_key = Key_END

	} else if key == glfw.KeyDelete {
		_key = Key_DELETE
	} else if key == glfw.KeyInsert {
		_key = Key_INSERT
	} else if key == glfw.KeyBackspace {
		_key = Key_BACKSPACE
	} else if key == glfw.KeyTab {
		_key = Key_TAB
	} else if key == glfw.KeyEnter {
		_key = Key_ENTER
	} else if key == glfw.KeyEscape {
		_key = Key_ESCAPE
	} else if key == glfw.KeyRightBracket {
		_key = Key_RIGHT_BRACKET
	} else if key == glfw.KeyLeftBracket {
		_key = Key_LEFT_BRACKET
	} else if key == glfw.KeyBackslash {
		_key = Key_BACK_SLASH
	} else if key == glfw.KeyEqual {
		_key = Key_EQUAL
	} else if key == glfw.KeySlash {
		_key = Key_FORWARD_SLASH
	} else if key == glfw.KeySemicolon {
		_key = Key_SEMICOLON
	} else if key == glfw.KeyPeriod {
		_key = Key_PERIOD
	} else if key == glfw.KeyMinus {
		_key = Key_MINUS
	} else if key == glfw.KeyComma {
		_key = Key_COMMA
	} else if key == glfw.KeySpace {
		_key = Key_SPACE
	} else if key == glfw.KeyApostrophe {
		_key = Key_APOSTROPHE

	} else {
		fmt.Println("Unrecognized key pressed %i", key)
		return
	}

	if action == glfw.Press {
		_action = Button_Action_Down

	} else if action == glfw.Release {
		_action = Button_Action_Up


	} else if action == glfw.Repeat {
		_action = Button_Action_Down

	} else {
		//fmt.Println("Unrecognized key acti %i in xel-keyboard.go", action)
		return
	}


	if _onKey != nil {
		_onKey(_key, _action, mods & glfw.ModAlt != 0, mods & glfw.ModControl != 0, mods & glfw.ModShift != 0)
	}
}

func KeyToChar(key bl.KeyboardKey, shift, numlock bool) string {
	if key == Key_SPACE {
		return " "
	}

	if key == Key_A {
		if shift {
			return "A"
		} else {
			return "a"
		}
	}

	if key == Key_B {
		if shift {
			return "B"
		} else {
			return "b"
		}
	}

	if key == Key_C {
		if shift {
			return "C"
		} else {
			return "c"
		}
	}

	if key == Key_D {
		if shift {
			return "D"
		} else {
			return "d"
		}
	}

	if key == Key_E {
		if shift {
			return "E"
		} else {
			return "e"
		}
	}

	if key == Key_F {
		if shift {
			return "F"
		} else {
			return "f"
		}
	}

	if key == Key_G {
		if shift {
			return "G"
		} else {
			return "g"
		}
	}

	if key == Key_H {
		if shift {
			return "H"
		} else {
			return "h"
		}
	}

	if key == Key_I {
		if shift {
			return "I"
		} else {
			return "i"
		}
	}

	if key == Key_J {
		if shift {
			return "J"
		} else {
			return "j"
		}
	}

	if key == Key_K {
		if shift {
			return "K"
		} else {
			return "k"
		}
	}

	if key == Key_L {
		if shift {
			return "L"
		} else {
			return "l"
		}
	}

	if key == Key_M {
		if shift {
			return "M"
		} else {
			return "m"
		}
	}

	if key == Key_N {
		if shift {
			return "N"
		} else {
			return "n"
		}
	}

	if key == Key_O {
		if shift {
			return "O"
		} else {
			return "o"
		}
	}

	if key == Key_P {
		if shift {
			return "P"
		} else {
			return "p"
		}
	}

	if key == Key_Q {
		if shift {
			return "Q"
		} else {
			return "q"
		}
	}

	if key == Key_R {
		if shift {
			return "R"
		} else {
			return "r"
		}
	}

	if key == Key_S {
		if shift {
			return "S"
		} else {
			return "s"
		}
	}

	if key == Key_T {
		if shift {
			return "T"
		} else {
			return "t"
		}
	}

	if key == Key_U {
		if shift {
			return "U"
		} else {
			return "u"
		}
	}

	if key == Key_V {
		if shift {
			return "V"
		} else {
			return "v"
		}
	}

	if key == Key_W {
		if shift {
			return "W"
		} else {
			return "w"
		}
	}

	if key == Key_X {
		if shift {
			return "X"
		} else {
			return "x"
		}
	}

	if key == Key_Y {
		if shift {
			return "Y"
		} else {
			return "y"
		}
	}

	if key == Key_Z {
		if shift {
			return "Z"
		} else {
			return "z"
		}
	}

	if key == Key_RIGHT_BRACKET {
		if shift {
			return "}"
		} else {
			return "]"
		}
	}

	if key == Key_LEFT_BRACKET {
		if shift {
			return "{"
		} else {
			return "["
		}
	}

	if key == Key_BACK_SLASH {
		if shift {
			return "|"
		} else {
			return "\\"
		}
	}

	if key == Key_EQUAL {
		if shift {
			return "+"
		} else {
			return "="
		}
	}

	if key == Key_FORWARD_SLASH {
		if shift {
			return "?"
		} else {
			return "/"
		}
	}

	if key == Key_SEMICOLON {
		if shift {
			return ":"
		} else {
			return ";"
		}
	}

	if key == Key_PERIOD {
		if shift {
			return ">"
		} else {
			return "."
		}
	}

	if key == Key_MINUS {
		if shift {
			return "_"
		} else {
			return "-"
		}
	}

	if key == Key_COMMA {
		if shift {
			return "<"
		} else {
			return ","
		}
	}

	if key == Key_APOSTROPHE {
		if shift {
			return "~"
		} else {
			return "`"
		}
	}

	if key == Key_0 {
		if shift {
			return ")"
		} else {
			return "0"
		}
	}

	if key == Key_1 {
		if shift {
			return "!"
		} else {
			return "1"
		}
	}

	if key == Key_2 {
		if shift {
			return "@"
		} else {
			return "2"
		}
	}

	if key == Key_3 {
		if shift {
			return "#"
		} else {
			return "3"
		}
	}

	if key == Key_4 {
		if shift {
			return "$"
		} else {
			return "4"
		}
	}

	if key == Key_5 {
		if shift {
			return "%"
		} else {
			return "5"
		}
	}

	if key == Key_6 {
		if shift {
			return "^"
		} else {
			return "6"
		}
	}

	if key == Key_7 {
		if shift {
			return "&"
		} else {
			return "7"
		}
	}

	if key == Key_8 {
		if shift {
			return "*"
		} else {
			return "8"
		}
	}

	if key == Key_9 {
		if shift {
			return "("
		} else {
			return "9"
		}
	}

	if key == Key_PAD_1 {

		if numlock {
			return "1"
		}
	}

	if key == Key_PAD_2 {

		if numlock {
			return "2"
		}
	}

	if key == Key_PAD_3 {

		if numlock {
			return "3"
		}
	}

	if key == Key_PAD_4 {

		if numlock {
			return "4"
		}
	}

	if key == Key_PAD_5 {

		if numlock {
			return "5"
		}
	}

	if key == Key_PAD_6 {

		if numlock {
			return "6"
		}
	}

	if key == Key_PAD_7 {

		if numlock {
			return "7"
		}
	}

	if key == Key_PAD_8 {

		if numlock {
			return "8"
		}
	}

	if key == Key_PAD_9 {

		if numlock {
			return "9"
		}
	}

	if key == Key_PAD_0 {

		if numlock {
			return "0"
		}
	}

	if key == Key_PAD_DECIMAL {

		if numlock {
			return "."
		}
	}

	if key == Key_PAD_DIVIDE {
		return "/"
	}

	if key == Key_PAD_MULTIPLY {
		return "*"
	}

	if key == Key_PAD_SUBTRACT {
		return "-"
	}

	if key == Key_PAD_ADD {
		return "+"
	}

	if key == Key_PAD_EQUAL {
		return "="
	}

	return ""
}

type KeyBehaviorType int

const (
	Key_Behavior_HOME KeyBehaviorType = 1 + iota
	Key_Behavior_INSERT
	Key_Behavior_PAGE_UP
	Key_Behavior_PAGE_DOWN
	Key_Behavior_DOWN
	Key_Behavior_UP
	Key_Behavior_RIGHT
	Key_Behavior_LEFT
	Key_Behavior_NIL
	Key_Behavior_END
	Key_Behavior_FNUM
	Key_Behavior_DELETE
	Key_Behavior_ENTER
	Key_Behavior_BACKSPACE
	Key_Behavior_TAB
	Key_Behavior_ESCAPE
	Key_Behavior_UNKNOWN
	Key_Behavior_CHAR
)

func KeyToBehavior(key bl.KeyboardKey, shift, numlock bool) KeyBehaviorType {

	if KeyToChar(key, shift, numlock) != "" {
		return Key_Behavior_CHAR
	}

	if key == Key_PAD_1 {

		if !numlock {
			return Key_Behavior_END
		}
	}

	if key == Key_PAD_2 {

		if !numlock {
			return Key_Behavior_DOWN
		}
	}

	if key == Key_PAD_3 {

		if !numlock {
			return Key_Behavior_PAGE_DOWN
		}
	}

	if key == Key_PAD_4 {

		if !numlock {
			return Key_Behavior_LEFT
		}
	}

	if key == Key_PAD_5 {

		if !numlock {
			return Key_Behavior_NIL
		}
	}

	if key == Key_PAD_6 {

		if !numlock {
			return Key_Behavior_RIGHT
		}
	}

	if key == Key_PAD_7 {

		if !numlock {
			return Key_Behavior_HOME
		}
	}

	if key == Key_PAD_8 {

		if !numlock {
			return Key_Behavior_UP
		}
	}

	if key == Key_PAD_9 {

		if !numlock {
			return Key_Behavior_PAGE_UP
		}
	}

	if key == Key_PAD_0 {

		if !numlock {
			return Key_Behavior_INSERT
		}
	}

	if key == Key_F1 || key == Key_F2 || key == Key_F3 || key == Key_F4 || key == Key_F5 || key == Key_F6 || key == Key_F7 || key == Key_F8 || key == Key_F9 || key == Key_F10 || key == Key_F11 || key == Key_F12 {

		return Key_Behavior_FNUM
	}

	if key == Key_PAD_DECIMAL {

		if !numlock {
			return Key_Behavior_DELETE
		}
	}

	if key == Key_PAD_ENTER {
		return Key_Behavior_ENTER
	}

	if key == Key_RIGHT {
		return Key_Behavior_RIGHT
	}

	if key == Key_LEFT {
		return Key_Behavior_LEFT
	}

	if key == Key_UP {
		return Key_Behavior_UP
	}

	if key == Key_DOWN {
		return Key_Behavior_DOWN
	}

	if key == Key_PAGE_UP {
		return Key_Behavior_PAGE_UP
	}

	if key == Key_PAGE_DOWN {
		return Key_Behavior_PAGE_DOWN
	}

	if key == Key_HOME {
		return Key_Behavior_HOME
	}

	if key == Key_END {
		return Key_Behavior_END
	}

	if key == Key_DELETE {
		return Key_Behavior_DELETE
	}

	if key == Key_INSERT {
		return Key_Behavior_INSERT
	}

	if key == Key_BACKSPACE {
		return Key_Behavior_BACKSPACE
	}

	if key == Key_TAB {
		return Key_Behavior_TAB
	}

	if key == Key_ENTER {
		return Key_Behavior_ENTER
	}

	if key == Key_ESCAPE {
		return Key_Behavior_ESCAPE
	}

	return Key_Behavior_UNKNOWN
}

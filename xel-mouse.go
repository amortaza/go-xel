package xel

import (
	"github.com/goxjs/glfw"
	"fmt"
	"github.com/amortaza/go-bellina"
)

var gUserOnMouseMove func(x, y int)
var gUserOnMouseButton func(button bl.MouseButton, action bl.ButtonAction)

func SetMouseCursor(cursor bl.MouseCursor) {
	fmt.Println("SetMouseCursor has not been implemented in xel2")
}

func __onMouseMove(window *glfw.Window, x, y float64) {
	MouseX, MouseY = int(x), int(y)

	if gUserOnMouseMove != nil {
		gUserOnMouseMove(int(x), int(y))
	}
}

func __onMouseButton(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	var _button bl.MouseButton
	var _action bl.ButtonAction

	if button == glfw.MouseButtonLeft {
		_button = Mouse_Button_Left

	} else if button == glfw.MouseButtonRight {
		_button = Mouse_Button_Right

	} else {
		fmt.Println("Unrecognized mouse button %i", button)
		return
	}

	if action == glfw.Press {
		_action = Button_Action_Down
	} else if action == glfw.Release {
		_action = Button_Action_Up
	} else {
		fmt.Println("Unrecognized action %i", action)
		return
	}

	if gUserOnMouseButton != nil {
		gUserOnMouseButton(_button, _action)
	}
}

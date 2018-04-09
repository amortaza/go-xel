package xel

import (
	"github.com/amortaza/go-glfw"
	//"fmt"
	"fmt"
	"github.com/amortaza/go-hal"
)

var gUserOnMouseMove func(x, y int)
var gUserOnMouseButton func(button hal.MouseButton, action hal.ButtonAction)

func SetMouseCursor(cursor hal.MouseCursor) {

	if cursor == hal.MouseCursor_Arrow {
		g_window.SetCursor(glfw.ArrowCursor)

	} else if cursor == hal.MouseCursor_Horiz_Resize {
		g_window.SetCursor(glfw.HResizeCursor)

	} else if cursor == hal.MouseCursor_Vert_Resize {
		g_window.SetCursor(glfw.VResizeCursor)

	} else if cursor == hal.MouseCursor_IBeam {
		g_window.SetCursor(glfw.IBeamCursor)

	} else if cursor == hal.MouseCursor_Hand {
		g_window.SetCursor(glfw.HandCursor)

	} else if cursor == hal.MouseCursor_CrossHair {
		g_window.SetCursor(glfw.CrossHairCursor)

	} else {
		fmt.Println("Did not recognize the Mouse cursor in xel2.SetMouseCursor")
	}
}

func glfw_onMouseMove(
		window *glfw.Window,
		x, y float64) {

	MouseX, MouseY = int(x), int(y)

	if gUserOnMouseMove != nil {
		gUserOnMouseMove(int(x), int(y))
	}
}

func glfw_onMouseButton(
		window *glfw.Window,
		button glfw.MouseButton,
		action glfw.Action,
		mods glfw.ModifierKey) {

	var _button hal.MouseButton
	var _action hal.ButtonAction

	if button == glfw.MouseButtonLeft {
		_button = hal.Mouse_Button_Left

	} else if button == glfw.MouseButtonRight {
		_button = hal.Mouse_Button_Right

	} else {

		fmt.Println("Unrecognized mouse button %i", button)
		return
	}

	if action == glfw.Press {
		_action = hal.Button_Action_Down

	} else if action == glfw.Release {
		_action = hal.Button_Action_Up

	} else {

		fmt.Println("Unrecognized action %i", action)
		return
	}

	if gUserOnMouseButton != nil {
		gUserOnMouseButton(_button, _action)
	}
}


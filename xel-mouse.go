package xel

import (
	"github.com/amortaza/go-glfw"
	//"fmt"
	"fmt"
)

var MouseX, MouseY int

type MouseButton int
type ButtonAction int

var gUserOnMouseMove func(x, y int)
var gUserOnMouseButton func(button MouseButton, action ButtonAction)

const (
	Mouse_Button_Left MouseButton = 1 + iota
	Mouse_Button_Right
)

/*
func SetMouseCursor(cursor bl.MouseCursor) {
	if cursor == bl.MouseCursor_Arrow {
		gWindow.SetCursor(glfw.ArrowCursor)

	} else if cursor == bl.MouseCursor_Horiz_Resize {
		gWindow.SetCursor(glfw.HResizeCursor)

	} else if cursor == bl.MouseCursor_Vert_Resize {
		gWindow.SetCursor(glfw.VResizeCursor)

	} else if cursor == bl.MouseCursor_IBeam {
		gWindow.SetCursor(glfw.IBeamCursor)

	} else if cursor == bl.MouseCursor_Hand {
		gWindow.SetCursor(glfw.HandCursor)

	} else if cursor == bl.MouseCursor_CrossHair {
		gWindow.SetCursor(glfw.CrossHairCursor)
	} else {
		fmt.Println("Did not recognize the Mouse cursor in xel2.SetMouseCursor")
	}
}
*/
func xel_onMouseMove(window *glfw.Window, x, y float64) {
	MouseX, MouseY = int(x), int(y)

	if gUserOnMouseMove != nil {
		gUserOnMouseMove(int(x), int(y))
	}
}

func xel_onMouseButton(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	var _button MouseButton
	var _action ButtonAction

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

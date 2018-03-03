package xel

import (
	"github.com/amortaza/go-glfw"
)

var gWindow *glfw.Window

// cannot change the width/height types - must be "int"
func xel_onResize(window *glfw.Window, width int, height int) {

	WinWidth, WinHeight = width, height

	if gUserOnResize != nil {
		gUserOnResize(width, height)
	}
}

func createWindow(title string) {

	gWindow, _ = glfw.CreateWindow(int(WinWidth), int(WinHeight), title, nil, nil)

	gWindow.MakeContextCurrent()

	glfw.SwapInterval(0)

	gWindow.SetCursorPosCallback(xel_onMouseMove)
	gWindow.SetMouseButtonCallback(xel_onMouseButton)
	gWindow.SetKeyCallback(xel_onKey)
	gWindow.SetSizeCallback(xel_onResize)

	width, height := gWindow.GetSize()
	xel_onResize(gWindow, width, height)
}

var gUserOnAfterGL func()
var gUserOnBeforeWindowDelete func()
var gUserOnResize func(width, height int)


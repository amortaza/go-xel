package xel

import (
	"github.com/goxjs/glfw"
)

var gWindow *glfw.Window
var gTitle string

var gUserOnAfterGL func()
var gUserOnTick func()
var gUserOnBeforeDelete func()
var gUserOnResize func(width, height int)

// cannot change the width/height types - must be "int"
func __onResize(window *glfw.Window, width int, height int) {
	WinWidth, WinHeight = width, height

	if gUserOnResize != nil {
		gUserOnResize(width, height)
	}
}

func createWindow() {

	gWindow, _ = glfw.CreateWindow(int(WinWidth), int(WinHeight), gTitle, nil, nil)

	gWindow.MakeContextCurrent()
}

package xel

import (
	"github.com/amortaza/go-glfw"
	"fmt"
)

var g_window *glfw.Window

// cannot change the width/height types - must be "int"
func xel_onResize(window *glfw.Window, width int, height int) {

	WinWidth, WinHeight = width, height

	if g_user_OnResize != nil {
		g_user_OnResize(width, height)
	}
}

func createWindow(title string) {

	g_window, _ = glfw.CreateWindow(int(WinWidth), int(WinHeight), title, nil, nil)

	g_window.SetPos(g_win_left, g_win_top)

	g_window.MakeContextCurrent()

	fmt.Println("(+) GLFW Initialized")

	glfw.SwapInterval(0)

	g_window.SetCursorPosCallback(glfw_onMouseMove)
	g_window.SetMouseButtonCallback(glfw_onMouseButton)
	g_window.SetKeyCallback(glfw_onKey)
	g_window.SetSizeCallback(xel_onResize)

	width, height := g_window.GetSize()
	xel_onResize(g_window, width, height)
}

var g_user_OnAfterGL func()
var g_user_OnBefore_WindowDelete func()
var g_user_OnResize func(width, height int)


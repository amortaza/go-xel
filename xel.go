package xel

import (
	"github.com/amortaza/go-glfw"
	"github.com/goxjs/gl"
	"time"
	"fmt"
	"github.com/amortaza/go-hal"
)

var g_win_left, g_win_top int

var WinWidth, WinHeight int
var MouseX, MouseY int

func Init(left, top, width, height int) {

	if err := glfw.Init(gl.ContextWatcher); err != nil {
		panic("failed to initialize glfw")
	}

	g_win_left, g_win_top = left, top

	glfw.WindowHint(glfw.Resizable, int(glfw.Resizable))
	glfw.WindowHint(glfw.Samples, 4);

	WinWidth, WinHeight = width, height
}

func SetCallbacks(
			onAfterGL,
			onTick func(),
			onBeforeWindowDelete func(),
			onResize func(width, height int),
			onMouseMove func(x, y int),
			onMouseButton func(button hal.MouseButton, action hal.ButtonAction),
			onKey func(key hal.KeyboardKey, action hal.ButtonAction, alt, ctrl, shift bool)) {

	g_user_OnAfterGL = onAfterGL
	gUserOnTick = onTick
	g_user_OnBefore_WindowDelete = onBeforeWindowDelete
	g_user_OnResize = onResize
	gUserOnMouseMove = onMouseMove
	gUserOnMouseButton = onMouseButton
	gUserOnKey = onKey
}

func Loop(title string) {

	createWindow(title)

	if g_user_OnAfterGL != nil {
		g_user_OnAfterGL();
	}

	glfw.SwapInterval(1)

	for !g_window.ShouldClose() {

		then := time.Now().UnixNano()

		if gUserOnTick != nil {
			gUserOnTick()
		}

		g_window.SwapBuffers()

		glfw.PollEvents()

		for time.Now().UnixNano() - then < 25000000 { // 25 ms
			time.Sleep(2 * time.Millisecond) // 2ms
		}
	}

	if g_user_OnBefore_WindowDelete != nil {
		g_user_OnBefore_WindowDelete();
	}

	glfw.Terminate()

	fmt.Println("(-) GLFW Uninitialized")
}

var gUserOnTick func()
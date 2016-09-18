package xel

import (
	"github.com/goxjs/glfw"
	"github.com/goxjs/gl"
	"github.com/amortaza/go-bellina"
	"runtime"
	"time"
	"fmt"
)

var WinWidth, WinHeight int
var MouseX, MouseY int

func Init(title string, width, height int) {
	gTitle = title

	if err := glfw.Init(gl.ContextWatcher); err != nil {
		panic("failed to initialize glfw")
	}

	glfw.WindowHint(glfw.Resizable, int(glfw.Resizable))
	glfw.WindowHint(glfw.Samples, 4);

	WinWidth, WinHeight = width, height
}

func Uninit() {
	glfw.Terminate()
	fmt.Println("(-) GLFW terminated")
}

func SetCallbacks(	onAfterGL,
			onTick func(),
			onBeforeDelete func(),
			onResize func(width, height int),
			onMouseMove func(x, y int),
			onMouseButton func(button bl.MouseButton, action bl.ButtonAction),
			onKey func(key bl.KeyboardKey, action bl.ButtonAction, alt, ctrl, shift bool)) {

	gUserOnAfterGL = onAfterGL
	gUserOnTick = onTick
	gUserOnBeforeDelete = onBeforeDelete
	gUserOnResize = onResize
	gUserOnMouseMove = onMouseMove
	gUserOnMouseButton = onMouseButton
	_onKey = onKey
}

func Loop() {

	createWindow()

	glfw.SwapInterval(0)

	gWindow.SetCursorPosCallback(__onMouseMove)
	gWindow.SetMouseButtonCallback(__onMouseButton)
	gWindow.SetKeyCallback(__onKey)
	gWindow.SetSizeCallback(__onResize)

	width, height := gWindow.GetSize()
	__onResize(gWindow, width, height)

	if gUserOnAfterGL != nil {
		gUserOnAfterGL();
	}

	for !gWindow.ShouldClose() {
		then := time.Now().UnixNano()

		if gUserOnTick != nil {
			gUserOnTick()
		}

		gWindow.SwapBuffers()

		glfw.PollEvents()

		if true {
			for time.Now().UnixNano() - then < 25000000 {
				// 25ms
				time.Sleep(2 * time.Millisecond) // 2ms
			}

			if false {
				runtime.GC()
			}
		}

	}

	if gUserOnBeforeDelete != nil {
		gUserOnBeforeDelete();
	}
}
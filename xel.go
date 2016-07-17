package xel

import (
	"github.com/goxjs/glfw"
	"github.com/goxjs/gl"
	"runtime"
	"time"
	"fmt"
)

type MouseButton int
type KeyboardKey int

type ButtonAction int

const (
	Mouse_Button_Left MouseButton = 1 + iota
	Mouse_Button_Right
)

const (
	Button_Action_Down ButtonAction = 1 + iota
	Button_Action_Up
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
			onMouseButton func(button MouseButton, action ButtonAction),
			onKey func(key KeyboardKey, action ButtonAction, alt, ctrl, shift bool)) {

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

		for time.Now().UnixNano() - then < 30000000 { // 30ms
			time.Sleep(10 * time.Millisecond) // 10ms
		}

		// #ifdef nonprod
		runtime.GC()
	}

	if gUserOnBeforeDelete != nil {
		gUserOnBeforeDelete();
	}
}


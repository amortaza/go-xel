package xel

import (
	"github.com/amortaza/go-glfw"
	"github.com/goxjs/gl"
	"time"
	"fmt"
)

func Init(width, height int) {

	if err := glfw.Init(gl.ContextWatcher); err != nil {
		panic("failed to initialize glfw")
	}

	glfw.WindowHint(glfw.Resizable, int(glfw.Resizable))
	glfw.WindowHint(glfw.Samples, 4);

	WinWidth, WinHeight = width, height
}

func SetCallbacks(
			onAfterLoadGL,
			onTick func(),
			onBeforeUnloadGL func(),
			onResize func(width, height int),
			onMouseMove func(x, y int),
			onMouseButton func(button MouseButton, action ButtonAction),
			onKey func(key KeyboardKey, action ButtonAction, alt, ctrl, shift bool)) {

	gUserOnAfterLoadGL = onAfterLoadGL
	gUserOnTick = onTick
	gUserOnBeforeUnloadGL = onBeforeUnloadGL
	gUserOnResize = onResize
	gUserOnMouseMove = onMouseMove
	gUserOnMouseButton = onMouseButton
	gUserOnKey = onKey
}

func Loop(title string) {

	createWindow(title)

	if gUserOnAfterLoadGL != nil {
		gUserOnAfterLoadGL();
	}

	for !gWindow.ShouldClose() {

		then := time.Now().UnixNano()

		if gUserOnTick != nil {
			gUserOnTick()
		}

		gWindow.SwapBuffers()

		glfw.PollEvents()

		for time.Now().UnixNano() - then < 25000000 { // 25 ms
			time.Sleep(2 * time.Millisecond) // 2ms
		}
	}

	if gUserOnBeforeUnloadGL != nil {
		gUserOnBeforeUnloadGL();
	}

	glfw.Terminate()

	fmt.Println("(-) GLFW terminated")
}

var gUserOnTick func()
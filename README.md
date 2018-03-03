# what does go-xel do?

Gives Go (golang) applications a Window with an OpenGL (glfw) context with callbacks for Mouse and Keyboard.

# why does go-xel exist?

go-xel is used by the `Bellina UI Library`

# usage

Setup in Windows
>set GOROOT=*wherever your go is installed*

>set GOPATH=*wherever your go project is*

Put the go-xel repository under 
>%GOPATH%\src\github.com\amortaza\go-xel

For example:

`cd %GOPATH%\src\github.com\amortaza\`

`git clone git@github.com:amortaza/go-xel.git`

# note about the examples

You may need pre-requisites to run `example-xel-nanovg.go`.  For full instructions on how to setup your environment, please refer to `github.com/amortaza/go-bellina-tutorials/tutorial-00`.

Pay careful attention to the init() function in the example.
runtime.LockOSThread() must be called in init().

```go
package main

import (
	"github.com/amortaza/go-xel"
	"runtime"
)

func onLoop() {
	// this gets called every loop...use it wisely
}

func onAfterGL() {
	// This callback gets called after the OpenGL context has been
	// initialized.
	// All OpenGL calls are valid now
}

func onBeforeWindowDelete() {
	// Right before the window is deleted, this callback can be used
	// to free up resources.
}

func onResize(width, height int) {
	// This callback gets called for every resize of the window.
}

func onMouseMove(x, y int) {
	// This callback gets called for every movement of the mouse.
}

func onMouseButton(button xel.MouseButton, action xel.ButtonAction) {
	// This callback gets called for every press of the mouse button.
}

func onKey(key xel.KeyboardKey, action xel.ButtonAction, alt, ctrl, shift bool) {
	// This callback gets called for every press of the keyboard.
}

func init() {
	// This bit is required for GLFW.
	runtime.LockOSThread()
}

func main() {

	// Create an 800 by 600 pixel window (is this 1997 again!)
	xel.Init(800, 600)

	// Setup the callbacks
	xel.SetCallbacks(onAfterGL, onLoop, onBeforeWindowDelete, onResize, onMouseMove, onMouseButton, onKey)

	// Start the loop!
	xel.Loop("Hello, World!")
}
```

# Questions and comments

Please do not hesitate to create issues, or email me.

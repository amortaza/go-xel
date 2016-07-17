package main

import (
	"github.com/amortaza/go-xel2"
	gl3 "github.com/chsc/gogl/gl33"
)


func afterGL() {
	e := gl3.Init()

	if e != nil {
		panic("Unable to initialize gl3")
	}

	gl3.ClearColor(0.3, 0.3, 0.32, 1.0)
	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
}

func onDelete() {
}

func onLoop() {
	gl3.Viewport(0, 0, 640, 480)

	gl3.ClearColor(0.3, 0.3, 0.32, 1.0)

	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
}

func onResize(width, height int) {
}

func main() {
	xel.Init("Welcome, home!", 640, 480)

	xel.SetCallbacks(afterGL, onLoop, onDelete, onResize, nil, nil, nil )
	xel.Loop()
	xel.Uninit()
}


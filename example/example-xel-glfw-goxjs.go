package main

import "github.com/amortaza/go-xel-glfw-goxjs"

func main() {
	xel.Init("Welcome, home!", 640, 480)
	xel.Loop()
	xel.Uninit()
}

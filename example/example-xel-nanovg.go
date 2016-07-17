package main

import (
	"github.com/amortaza/go-xel2"
	"github.com/shibukawa/nanovgo"
	"fmt"
	"github.com/shibukawa/nanovgo/sample/demo"
	"log"
	gl3 "github.com/chsc/gogl/gl33"
)

var ctx *nanovgo.Context
var demoData *demo.DemoData

func afterGL() {
	var err error

	ctx, err = nanovgo.NewContext(0 /*nanovgo.AntiAlias | nanovgo.StencilStrokes | nanovgo.Debug*/)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("(+) Created nanovgo context")

	demoData = LoadDemo(ctx)

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	gl3.ClearColor(0.3, 0.3, 0.32, 1.0)
	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
}

func onDelete() {
	demoData.FreeData(ctx)

	fmt.Println("(-) Deleting nanovg context")
	ctx.Delete()
}

var t float32
var w, h int = 0,0

func onLoop() {
	t++
	gl3.Viewport(0, 0, gl3.Sizei(w), gl3.Sizei(h))
	gl3.ClearColor(0.3, 0.3, 0.32, 1.0)
	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
	gl3.Enable(gl3.BLEND)
	gl3.BlendFunc(gl3.SRC_ALPHA, gl3.ONE_MINUS_SRC_ALPHA)
	gl3.Enable(gl3.CULL_FACE)
	gl3.Disable(gl3.DEPTH_TEST)

	ctx.BeginFrame(w, h, 1)

	demo.RenderDemo(ctx, float32(1), float32(1), float32(w), float32(h), t, false, demoData)
	//fps.RenderGraph(ctx, 5, 5)

	ctx.EndFrame()

	gl3.Enable(gl3.DEPTH_TEST)

}

func onResize(a,b int) {
	w,h=a,b
	fmt.Println(w, " ", h)
}

func main() {
	xel.Init("Welcome, home!", 640, 480)

	xel.SetCallbacks(afterGL, onLoop, onDelete, onResize, nil, nil, nil )
	xel.Loop()
	xel.Uninit()
}

func LoadDemo(ctx *nanovgo.Context) *demo.DemoData {
	d := &demo.DemoData{}
	for i := 0; i < 12; i++ {
		path := fmt.Sprintf("github.com/shibukawa/nanovgo/sample/images/image%d.jpg", i+1)
		d.Images = append(d.Images, ctx.CreateImage(path, 0))
		if d.Images[i] == 0 {
			log.Fatalf("Could not load %s", path)
		}
	}

	d.FontIcons = ctx.CreateFont("icons", "github.com/shibukawa/nanovgo/sample/entypo.ttf")
	if d.FontIcons == -1 {
		log.Fatalln("Could not add font icons.")
	}
	d.FontNormal = ctx.CreateFont("sans", "github.com/shibukawa/nanovgo/sample/Roboto-Regular.ttf")
	if d.FontNormal == -1 {
		log.Fatalln("Could not add font italic.")
	}
	d.FontBold = ctx.CreateFont("sans-bold", "github.com/shibukawa/nanovgo/sample/Roboto-Bold.ttf")
	if d.FontBold == -1 {
		log.Fatalln("Could not add font bold.")
	}
	return d
}

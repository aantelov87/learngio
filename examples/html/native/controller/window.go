package controller

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/p9c/learngio/examples/html/native/model"
)

var (
	list = &layout.List{
		Axis: layout.Vertical,
	}
)

func createBrowserWindow(gtx *layout.Context, th *material.Theme, document *model.NodeDOM) model.AppWindow {
	defaultCursor := glfw.CreateStandardCursor(glfw.ArrowCursor)
	pageTile := getPageTitle(document) + " - thdwb"

	browserWindow := model.AppWindow{
		Width:  900,
		Height: 600,

		ViewportWidth:  0,
		ViewportHeight: 0,

		AddressbarWidth:  0,
		AddressbarHeight: 0,

		CursorX: 0,
		CursorY: 0,

		DefaultCursor: defaultCursor,

		Title:  pageTile,
		Redraw: true,
		Resize: true,
		Reflow: true,

		ViewportOffset: 0,

		DOM: document,
	}

	//glfw.WindowHint(glfw.StencilBits, 8)
	//glfw.WindowHint(glfw.DepthBits, 0)

	//window, err := glfw.CreateWindow(browserWindow.Width, browserWindow.Height, browserWindow.Title, nil, nil)
	//if err != nil {
	//	log.Fatalf("Error creating window: %v", err)
	//}
	//window.MakeContextCurrent()

	//err = gl.Init()
	//if err != nil {
	//	log.Fatalf("Error initializing GL: %v", err)
	//}
	//
	//glfw.SwapInterval(1)
	//gl.Enable(gl.MULTISAMPLE)

	//addressbarBackend, err := goglbackend.New(0, 0, 0, 0, nil)
	//if err != nil {
	//	log.Fatalf("Error loading canvas GL assets: %v", err)
	//}
	//
	//viewportBackend, err := goglbackend.New(0, 0, 0, 0, nil)
	//if err != nil {
	//	log.Fatalf("Error loading canvas GL assets: %v", err)
	//}

	//browserWindow.GlfwWindow = window
	//browserWindow.AddressbarBackend = addressbarBackend
	//browserWindow.ViewportBackend = viewportBackend

	return browserWindow
}

func browserWindowMainLoop(gtx *layout.Context, th *material.Theme, document *model.NodeDOM) {
	//for !browserWindow.GlfwWindow.ShouldClose() {
	//	if browserWindow.Redraw {
	//		browserWindow.GlfwWindow.MakeContextCurrent()
	//		//windowWidth, windowHeight := browserWindow.GlfwWindow.GetSize()
	//
	//		if windowWidth != browserWindow.Width || windowHeight != browserWindow.Height {
	//			browserWindow.Width = windowWidth
	//			browserWindow.Height = windowHeight
	//			browserWindow.Resize = true
	//			browserWindow.Reflow = true
	//		}
	//
	//		browserWindow.AddressbarWidth = browserWindow.Width
	//		browserWindow.AddressbarHeight = 40
	//
	//		browserWindow.ViewportWidth = browserWindow.Width
	//		browserWindow.ViewportHeight = browserWindow.Height - browserWindow.AddressbarHeight
	//
	//		//browserWindow.AddressbarBackend.SetBounds(0, browserWindow.ViewportHeight, browserWindow.AddressbarWidth, browserWindow.AddressbarHeight)
	//		//browserWindow.ViewportBackend.SetBounds(0, 0, browserWindow.ViewportWidth, browserWindow.ViewportHeight)
	//
	//		updateAddressBar(browserWindow)
	//
	//if browserWindow.Reflow {
	//	browserWindow.DOM.Style.Width = float64(browserWindow.ViewportWidth)
	//	browserWindow.DOM.Style.Height = float64(browserWindow.ViewportHeight)
	//
	//	browserWindow.DOM.Style.Top = 0
	//	browserWindow.DOM.Style.Left = 0
	widgets := []func(){}

	nodeChildren := getNodeChildren(document)
	for i := 0; i < len(nodeChildren); i++ {
		ReflowNode(nodeChildren[i], nodeChildren[i], 0)
		///
		widgets = append(widgets, func() {
			ch := th.Caption(fmt.Sprint(nodeChildren[i].Element))
			ch.Layout(gtx)
		})

		//}

		////
	}
	//
	//	browserWindow.Reflow = false
	//}

	//updateViewport(browserWindow)

	//browserWindow.Resize = false
	//browserWindow.Redraw = false
	//browserWindow.GlfwWindow.SwapBuffers()
	//}

	//glfw.WaitEvents()
	//}
	list.Layout(gtx, len(widgets), func(i int) {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})

}

func updateAddressBar(browserWindow *model.AppWindow) {
	//oldAddressBarInput := getUIElementByID(browserWindow.UIElements, "addressbarInput")

	if browserWindow.Resize {
		browserWindow.Initialized = false
		browserWindow.UIElements = nil
	}

	if browserWindow.Initialized {
		for i := 0; i < len(browserWindow.UIElements); i++ {
			browserWindow.UIElements[i].Redraw()
		}
	} else {
		//w := float64(browserWindow.AddressbarWidth)
		//h := float64(browserWindow.AddressbarHeight)

		//addressbarBackground := Box("addressbarBackground", 0, 0, float64(browserWindow.Width), float64(browserWindow.AddressbarHeight), browserWindow.Addressbar)

		//addressbarText := browserWindow.Location
		//if oldAddressBarInput != nil {
		//	addressbarText = oldAddressBarInput.Text
		//}

		//addressbarInput := Input("addressbarInput", w, h, browserWindow.Addressbar, addressbarText)

		//browserWindow.UIElements = append(browserWindow.UIElements, &addressbarBackground, &addressbarInput)
		browserWindow.Initialized = true
	}
}

func updateViewport(browserWindow *model.AppWindow) {
	//viewport := browserWindow.Viewport
	vO := float64(browserWindow.ViewportOffset)

	//w := float64(browserWindow.ViewportWidth)
	//h := float64(browserWindow.ViewportHeight)

	//viewport.SetFillStyle("#FFF")
	//viewport.FillRect(0, 0, w, h)
	renderNode(browserWindow.DOM, browserWindow, vO)
}

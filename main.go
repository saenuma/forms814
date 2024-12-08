package main

import (
	"runtime"
	"time"
	"os"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	rootPath, err := GetRootPath()
	if err != nil {
		panic(err)
	}


	// make default project
	os.MkdirAll(filepath.Join(rootPath, "first_proj"), 0777)

	runtime.LockOSThread()

	window := g143.NewWindow(1200, 800, ProgTitle, false)
	ProjectName = "first_proj"
	DrawBeginView(window, "first_proj")

	// respond to the mouse
	window.SetMouseButtonCallback(projViewMouseCallback)
	// respond to the keyboard
	window.SetKeyCallback(ProjKeyCallback)
	// save the project file
	window.SetCloseCallback(SaveProjectCloseCallback)
	// quick hover effect
	window.SetCursorPosCallback(getHoverCB(ProjObjCoords))

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		time.Sleep(time.Second/time.Duration(FPS) - time.Since(t))
	}
}

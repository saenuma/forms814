package main

import (
	"os"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func ProjKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if key == glfw.KeyBackspace && len(NameInputEnteredTxt) != 0 {
		NameInputEnteredTxt = NameInputEnteredTxt[:len(NameInputEnteredTxt)-1]
	} else if key == glfw.KeySpace {
		NameInputEnteredTxt += " "
	} else if key == glfw.KeyEnter && len(NameInputEnteredTxt) != 0 {
		// create file
		rootPath, _ := GetRootPath()

		ProjectName = NameInputEnteredTxt + ".f8p"
		outPath := filepath.Join(rootPath, ProjectName)
		os.WriteFile(outPath, []byte(""), 0777)

		// // move to work view
		DrawWorkView(window, 1)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		return
	} else {
		NameInputEnteredTxt += glfw.GetKeyName(key, scancode)
	}

	nIRS := ProjObjCoords[PROJ_NameInput]
	theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
	theCtx.drawInput(PROJ_NameInput, nIRS.OriginX, nIRS.OriginY, nIRS.Width, NameInputEnteredTxt, true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

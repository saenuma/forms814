package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func projViewMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range ProjObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := GetRootPath()

	switch widgetCode {
	case PROJ_NewProject:
		if NameInputEnteredTxt == "" {
			return
		}

		// create file
		ProjectName = NameInputEnteredTxt + ".f8p"
		outPath := filepath.Join(rootPath, ProjectName)
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		DrawWorkView(window, 1)
		// window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		// quick hover effect
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))
	}

	if widgetCode > 1000 && widgetCode < 2000 {
		num := widgetCode - 1000 - 1
		projectFile := GetProjectFiles()[num]

		ProjectName = projectFile.Name

		// load instructions
		obj := make([]map[string]string, 0)
		rootPath, _ := GetRootPath()
		inPath := filepath.Join(rootPath, ProjectName)
		rawBytes, _ := os.ReadFile(inPath)
		json.Unmarshal(rawBytes, &obj)

		Instructions = append(Instructions, obj...)

		// move to work view
		DrawWorkView(window, 1)
		// window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))
	}
}

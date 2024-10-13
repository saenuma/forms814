package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"

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
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
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
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))
	}
}

func workViewMouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range WKObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			// break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case WK_AddFormBtn:
		// tmpFrame = CurrentWindowFrame
		DrawFormDialog(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(fdMouseBtnCallback)
		// window.SetKeyCallback(VaikeyCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(FDObjCoords))

	case WK_OpenWDBtn:
		rootPath, _ := GetRootPath()
		ExternalLaunch(rootPath)

	}
}

func fdMouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range FDObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			// break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case FD_CloseBtn:
		IsUpdateDialog = false
		IsInsertBeforeDialog = false

		DrawWorkView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

	case FD_NameInput, FD_LabelInput, FD_SelectOptionsInput:
		FD_SelectedInput = widgetCode

		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		if widgetCode == FD_SelectOptionsInput {
			theCtx.drawTextInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width,
				widgetRS.Height, EnteredTxts[widgetCode], true)
		} else {
			theCtx.drawInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, EnteredTxts[widgetCode], true)
		}

		// disable other inputs
		allInputs := []int{FD_NameInput, FD_LabelInput, FD_SelectOptionsInput}
		index := slices.Index(allInputs, widgetCode)
		leftInputs := slices.Delete(slices.Clone(allInputs), index, index+1)
		for _, inputId := range leftInputs {
			inputRS := FDObjCoords[inputId]
			if inputId == FD_SelectOptionsInput {
				theCtx.drawTextInput(inputId, inputRS.OriginX, inputRS.OriginY, inputRS.Width,
					inputRS.Height, EnteredTxts[inputId], false)
			} else {
				theCtx.drawInput(inputId, inputRS.OriginX, inputRS.OriginY, inputRS.Width, EnteredTxts[inputId], false)
			}
		}

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	default:
		FD_SelectedInput = 0

	}

	// for generated buttons
	if widgetCode > 200 && widgetCode < 300 {
		attribId := widgetCode - 200 - 1
		attrib := attributes[attribId]

		if value, ok := AttribState[attrib]; ok {
			AttribState[attrib] = !value
		} else {
			AttribState[attrib] = true
		}

		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawCheckbox(widgetCode, widgetRS.OriginX, widgetRS.OriginY, AttribState[attrib])

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()
		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if widgetCode > 300 && widgetCode < 400 {
		fTypeId := widgetCode - 300 - 1
		fType := supportedFields[fTypeId]

		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		for i := range supportedFields {
			cBtnId := 300 + i + 1
			cBtnRect := FDObjCoords[cBtnId]
			theCtx.drawCheckbox(cBtnId, cBtnRect.OriginX, cBtnRect.OriginY, false)
		}

		theCtx.drawCheckbox(widgetCode, widgetRS.OriginX, widgetRS.OriginY, true)

		SelectedFieldType = fType
		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()
		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	}

}

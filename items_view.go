package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tidwall/pretty"
)

func DrawWorkView(window *glfw.Window) {
	window.SetTitle(fmt.Sprintf("Form: %s / %s ---- %s", ProjectName, FormName, ProgTitle))

	WKObjCoords = make(map[int]g143.Rect)

	wWidth, wHeight := window.GetSize()
	theCtx := New2dCtx(wWidth, wHeight, &WKObjCoords)

	// draw top buttons
	bBRect := theCtx.drawButtonB(WK_BackBtn, 10, 10, "Back", "#fff", "#5C909C", "#286775")
	aFBX := nextHorizontalCoords(bBRect, 20)
	aFBRect := theCtx.drawButtonB(WK_AddFormBtn, aFBX, 10, "Add Form Item", "#fff", "#5F7E5D", "#889B87")

	currentX, currentY := 20, aFBRect.OriginY+aFBRect.Height+15
	for i, fObj := range FormObjects {
		theCtx.ggCtx.SetHexColor("#444")
		str2 := fmt.Sprintf("%d. label: %s name: %s fieldtype: %s attributes: %s",
			i+1, fObj["label"],
			fObj["name"], fObj["fieldtype"], fObj["attributes"])
		theCtx.ggCtx.DrawString(str2, float64(currentX), float64(currentY)+FontSize)

		addBeforeBtnId := 2000 + 1 + i
		aBBRect := theCtx.drawButtonA(addBeforeBtnId, currentX, currentY+30, "add before", "#fff", "#4E962D")
		eFOBX := nextHorizontalCoords(aBBRect, 20)
		eFOBtnId := 3000 + i + 1
		eFOBRect := theCtx.drawButtonA(eFOBtnId, eFOBX, currentY+30, "edit", "#fff", "#968D2D")
		dFOBX := nextHorizontalCoords(eFOBRect, 20)
		dFOBtnId := 4000 + i + 1
		theCtx.drawButtonA(dFOBtnId, dFOBX, currentY+30, "delete", "#fff", "#962D2D")
		currentY += 70
	}

	// send the frame to glfw window
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
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
		window.SetKeyCallback(FDKeyCallback)
		window.SetCharCallback(FDCharCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(FDObjCoords))

	case WK_BackBtn:
		// save formobject
		jsonBytes, _ := json.Marshal(FormObjects)
		prettyJsonBytes := pretty.Pretty(jsonBytes)
		rootPath, _ := GetRootPath()
		outPath := filepath.Join(rootPath, ProjectName, FormName)
		os.WriteFile(outPath, prettyJsonBytes, 0777)

		// draw projects selection view
		ProjectName = "first_proj"
		FormName = ""
		DrawBeginView(window, "first_proj")
		window.SetMouseButtonCallback(projViewMouseCallback)
		window.SetKeyCallback(ProjKeyCallback)
		window.SetCharCallback(nil)
		window.SetCursorPosCallback(getHoverCB(ProjObjCoords))
	}

	// for generated buttons
	if widgetCode > 2000 && widgetCode < 3000 {
		// add before selection
		objNum := widgetCode - 2000 - 1
		IsInsertBeforeDialog = true
		ToInsertBefore = objNum

		DrawFormDialog(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(fdMouseBtnCallback)
		window.SetKeyCallback(FDKeyCallback)
		window.SetCharCallback(FDCharCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(FDObjCoords))

	} else if widgetCode > 3000 && widgetCode < 4000 {
		// edit selection
		objNum := widgetCode - 3000 - 1
		ToUpdateInstrNum = objNum
		IsUpdateDialog = true

		DrawFormDialog(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(fdMouseBtnCallback)
		window.SetKeyCallback(FDKeyCallback)
		window.SetCharCallback(FDCharCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(FDObjCoords))

	} else if widgetCode > 4000 && widgetCode < 5000 {
		// delete from instructions slice
		objNum := widgetCode - 4000 - 1
		FormObjects = slices.Delete(FormObjects, objNum, objNum+1)

		WKObjCoords = make(map[int]g143.Rect)
		DrawWorkView(window)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

	}
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawBeginView(window *glfw.Window, project string) {
	ProjObjCoords = make(map[int]g143.Rect)
	wWidth, wHeight := window.GetSize()

	theCtx := New2dCtx(wWidth, wHeight, &ProjObjCoords)

	fontPath := GetDefaultFontPath()
	theCtx.ggCtx.LoadFontFace(fontPath, 30)

	theCtx.ggCtx.SetHexColor(fontColor)
	theCtx.ggCtx.DrawString("Projects", 20, 10+30)
	theCtx.ggCtx.LoadFontFace(fontPath, 20)

	theCtx.drawButtonB(BEGINV_OpenWDBtn, 200, 10, "Open Folder", "#fff", "#5C909C", "#286775")

	projects := GetProjects()
	currentX := 40
	currentY := 30 + 10 + 40

	projectsPaneWidth := (wWidth / 2) - 50
	for i, pf := range projects {

		btnId := 1000 + (i + 1)
		pfRect := theCtx.drawButtonA(btnId, currentX, currentY, pf, "#fff", "#5F699F")

		newX := currentX + pfRect.Width + 10
		if newX > (projectsPaneWidth - pfRect.Width) {
			currentY += 50
			currentX = 40
		} else {
			currentX += pfRect.Width + 10
		}

	}

	// new project form
	pnIRect := theCtx.drawInput(BEGINV_NameInput, 40, wHeight-60, 350, "enter project name", false)
	pnBtnX := nextHorizontalCoords(pnIRect, 30)
	theCtx.drawButtonA(BEGINV_NewProject, pnBtnX, pnIRect.OriginY, "New Project", fontColor, "#B3AE97")

	// second column
	secondColumnX := projectsPaneWidth + 50
	theCtx.ggCtx.LoadFontFace(fontPath, 30)
	theCtx.ggCtx.DrawString(fmt.Sprintf("Form Objects of %s", project), float64(secondColumnX), 10+30)

	// forms of project
	formsOfCurrentProject := GetProjectFiles(project)
	currentX = secondColumnX + 40
	currentY = 30 + 10 + 20
	theCtx.ggCtx.LoadFontFace(fontPath, 20)

	for i, fName := range formsOfCurrentProject {

		btnId := 2000 + (i + 1)
		pfRect := theCtx.drawButtonA(btnId, currentX, currentY, fName, "#fff", "#707695")

		newX := currentX + pfRect.Width + 10
		if newX > (wWidth - pfRect.Width) {
			currentY += 50
			currentX = secondColumnX + 40
		} else {
			currentX += pfRect.Width + 10
		}

	}

	// new formObject form
	fnIRect := theCtx.drawInput(BEGINV_FNameInput, secondColumnX+40, wHeight-60, 350, "enter form name", false)
	fnBtnX := nextHorizontalCoords(fnIRect, 30)
	theCtx.drawButtonA(BEGINV_NewForm, fnBtnX, fnIRect.OriginY, "New Form", fontColor, "#B3AE97")

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

func beginViewMouseCB(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range ProjObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := GetRootPath()

	switch widgetCode {
	case BEGINV_NewProject:
		// create project folder
		rootPath, _ := GetRootPath()

		outPath := filepath.Join(rootPath, NameInputEnteredTxt)
		os.MkdirAll(outPath, 0777)

		// redraw current view
		ProjectName = NameInputEnteredTxt
		DrawBeginView(window, NameInputEnteredTxt)
		window.SetCursorPosCallback(getHoverCB(ProjObjCoords))

	case BEGINV_OpenWDBtn:
		rootPath, _ := GetRootPath()
		ExternalLaunch(rootPath)

	case BEGINV_NewForm:
		if FNameInputEnteredTxt == "" {
			return
		}

		// create file
		FormName = FNameInputEnteredTxt + ".f8p"
		outPath := filepath.Join(rootPath, ProjectName, FormName)
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		drawItemsView(window)
		window.SetMouseButtonCallback(itemsViewMouseBtnCB)
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		// quick hover effect
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

	case BEGINV_NameInput, BEGINV_FNameInput:
		PV_SelectedInput = widgetCode

		theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
		theCtx.drawInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, EnteredTxts[widgetCode], true)

		// disable other inputs
		allInputs := []int{BEGINV_NameInput, BEGINV_FNameInput}
		index := slices.Index(allInputs, widgetCode)
		leftInputs := slices.Delete(slices.Clone(allInputs), index, index+1)
		for _, inputId := range leftInputs {
			inputRS := ProjObjCoords[inputId]
			theCtx.drawInput(inputId, inputRS.OriginX, inputRS.OriginY, inputRS.Width, EnteredTxts[inputId], false)
		}

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	}

	if widgetCode > 1000 && widgetCode < 2000 {
		num := widgetCode - 1000 - 1
		ProjectName = GetProjects()[num]

		rootPath, _ := GetRootPath()
		outPath := filepath.Join(rootPath, ProjectName)
		os.MkdirAll(outPath, 0777)

		// redraw current view
		DrawBeginView(window, ProjectName)
		window.SetCursorPosCallback(getHoverCB(ProjObjCoords))
	} else if widgetCode > 2000 && widgetCode < 3000 {
		num := widgetCode - 2000 - 1
		formsOfCurrentProject := GetProjectFiles(ProjectName)

		// create file
		FormName = formsOfCurrentProject[num]

		// load instructions
		obj := make([]map[string]string, 0)
		rootPath, _ := GetRootPath()
		inPath := filepath.Join(rootPath, ProjectName, FormName)
		rawBytes, _ := os.ReadFile(inPath)
		json.Unmarshal(rawBytes, &obj)

		FormObjects = obj

		// move to work view
		drawItemsView(window)
		window.SetMouseButtonCallback(itemsViewMouseBtnCB)
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		// quick hover effect
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))
	}
}

func ProjKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if PV_SelectedInput == BEGINV_NameInput {

		if key == glfw.KeyBackspace && len(NameInputEnteredTxt) != 0 {
			NameInputEnteredTxt = NameInputEnteredTxt[:len(NameInputEnteredTxt)-1]
		} else if key == glfw.KeyMinus && mods == glfw.ModShift {
			NameInputEnteredTxt = NameInputEnteredTxt + "_"
		} else if key == glfw.KeySpace {
			NameInputEnteredTxt += " "
		} else if key == glfw.KeyEnter && len(NameInputEnteredTxt) != 0 {
			// create project folder
			rootPath, _ := GetRootPath()

			outPath := filepath.Join(rootPath, NameInputEnteredTxt)
			os.MkdirAll(outPath, 0777)

			// redraw current view
			ProjectName = NameInputEnteredTxt
			DrawBeginView(window, NameInputEnteredTxt)
			window.SetCursorPosCallback(getHoverCB(ProjObjCoords))

			// window.SetMouseButtonCallback(itemsViewMouseBtnCB)
			// window.SetKeyCallback(nil)
			// window.SetScrollCallback(FirstUIScrollCallback)
			return
		} else {
			NameInputEnteredTxt += glfw.GetKeyName(key, scancode)
		}

		nIRS := ProjObjCoords[BEGINV_NameInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
		theCtx.drawInput(BEGINV_NameInput, nIRS.OriginX, nIRS.OriginY, nIRS.Width, NameInputEnteredTxt, true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if PV_SelectedInput == BEGINV_FNameInput {

		if key == glfw.KeyBackspace && len(FNameInputEnteredTxt) != 0 {
			FNameInputEnteredTxt = FNameInputEnteredTxt[:len(FNameInputEnteredTxt)-1]
		} else if key == glfw.KeyMinus && mods == glfw.ModShift {
			FNameInputEnteredTxt = FNameInputEnteredTxt + "_"
		} else if key == glfw.KeySpace {
			FNameInputEnteredTxt += " "
		} else if key == glfw.KeyEnter && len(FNameInputEnteredTxt) != 0 {
			// create file
			rootPath, _ := GetRootPath()

			FormName = FNameInputEnteredTxt + ".f8p"
			outPath := filepath.Join(rootPath, ProjectName, FormName)
			os.WriteFile(outPath, []byte(""), 0777)

			// // move to work view
			drawItemsView(window)
			window.SetMouseButtonCallback(itemsViewMouseBtnCB)
			window.SetKeyCallback(nil)
			// window.SetScrollCallback(FirstUIScrollCallback)
			return
		} else {
			FNameInputEnteredTxt += glfw.GetKeyName(key, scancode)
		}

		fnIRS := ProjObjCoords[BEGINV_FNameInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
		theCtx.drawInput(BEGINV_FNameInput, fnIRS.OriginX, fnIRS.OriginY, fnIRS.Width, FNameInputEnteredTxt, true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	}
}

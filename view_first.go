package main

import (
	"fmt"

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

	theCtx.drawButtonB(PROJ_OpenWDBtn, 200, 10, "Open Folder", "#fff", "#5C909C", "#286775")

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
	pnIRect := theCtx.drawInput(PROJ_NameInput, 40, wHeight-60, 350, "enter project name", false)
	pnBtnX := nextHorizontalCoords(pnIRect, 30)
	theCtx.drawButtonA(PROJ_NewProject, pnBtnX, pnIRect.OriginY, "New Project", fontColor, "#B3AE97")

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
	fnIRect := theCtx.drawInput(PROJ_FNameInput, secondColumnX+40, wHeight-60, 350, "enter form name", false)
	fnBtnX := nextHorizontalCoords(fnIRect, 30)
	theCtx.drawButtonA(PROJ_NewForm, fnBtnX, fnIRect.OriginY, "New Form", fontColor, "#B3AE97")

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

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

package main

import (
	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawBeginView(window *glfw.Window) {
	ProjObjCoords = make(map[int]g143.Rect)
	wWidth, wHeight := window.GetSize()

	theCtx := New2dCtx(wWidth, wHeight, &ProjObjCoords)

	fontPath := GetDefaultFontPath()
	theCtx.ggCtx.LoadFontFace(fontPath, 30)

	theCtx.ggCtx.SetHexColor(fontColor)
	theCtx.ggCtx.DrawString("New Project", 20, 10+30)

	theCtx.ggCtx.LoadFontFace(fontPath, 20)
	pnIRect := theCtx.drawInput(PROJ_NameInput, 40, 60, 420, "enter project name", false)
	pnBtnX, pnBtnY := nextHorizontalCoords(pnIRect, 30)
	theCtx.drawButtonA(PROJ_NewProject, pnBtnX, pnBtnY, "New Project", fontColor, "#B3AE97")

	// second row border
	_, borderY := nextVerticalCoords(pnIRect, 10)
	theCtx.ggCtx.SetHexColor("#999")
	theCtx.ggCtx.DrawRoundedRectangle(10, float64(borderY), float64(wWidth)-20, 2, 2)
	theCtx.ggCtx.Fill()

	theCtx.ggCtx.LoadFontFace(fontPath, 30)
	theCtx.ggCtx.SetHexColor(fontColor)
	theCtx.ggCtx.DrawString("Continue Projects", 20, float64(borderY)+12+30)
	theCtx.ggCtx.LoadFontFace(fontPath, 20)

	projectFiles := GetProjectFiles()
	currentX := 40
	currentY := borderY + 22 + 30 + 10
	for i, pf := range projectFiles {

		btnId := 1000 + (i + 1)
		pfRect := theCtx.drawButtonA(btnId, currentX, currentY, pf.Name, "#fff", "#5F699F")

		newX := currentX + pfRect.Width + 10
		if newX > (wWidth - pfRect.Width) {
			currentY += 50
			currentX = 40
		} else {
			currentX += pfRect.Width + 10
		}

	}

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

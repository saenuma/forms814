package main

import (
	"image"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kovidgoyal/imaging"
)

func DrawFormDialog(window *glfw.Window, currentFrame image.Image) {
	FDObjCoords = make(map[int]g143.Rect)
	// InputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()
	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	theCtx := Continue2dCtx(img, &FDObjCoords)

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 200

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	theCtx.ggCtx.SetHexColor("#fff")
	theCtx.ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	theCtx.ggCtx.Fill()

	// Add Form
	aFLX, aFLY := dialogOriginX+20, dialogOriginY+20
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("Add Form Configuration", float64(aFLX), float64(aFLY)+FontSize)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnRect := theCtx.drawButtonA(FD_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(FD_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

	// name input
	theCtx.ggCtx.SetHexColor("#444")
	fNLY := aFLY + 30
	theCtx.ggCtx.DrawString("field name:", float64(aFLX), float64(fNLY)+FontSize)

	// send the frame to glfw window
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

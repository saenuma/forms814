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
	dialogWidth := 950
	dialogHeight := 600

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	theCtx.ggCtx.SetHexColor("#fff")
	theCtx.ggCtx.DrawRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight))
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
	fNLY := aFLY + 50
	theCtx.ggCtx.DrawString("field name: ", float64(aFLX), float64(fNLY)+FontSize)
	fNLW, _ := theCtx.ggCtx.MeasureString("field name:")

	fNIX := aFLX + int(fNLW) + 20
	fNIRect := theCtx.drawInput(FD_NameInput, fNIX, fNLY, 250, "", false)

	currentX, currentY := fNIRect.OriginX+fNIRect.Width+30, fNLY
	// attributes checkboxes
	for i, field := range attributes {
		cBtnId := 200 + i + 1
		cIRect := theCtx.drawCheckbox(cBtnId, currentX, currentY, false)
		cILX, _ := nextHorizontalCoords(cIRect, 10)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(field, float64(cILX), float64(currentY)+FontSize)
		fieldW, _ := theCtx.ggCtx.MeasureString(field)
		currentX = cILX + int(fieldW) + 20
	}

	// label input
	theCtx.ggCtx.SetHexColor("#444")
	_, fLLY := nextVerticalCoords(fNIRect, 15)
	theCtx.ggCtx.DrawString("field label:", float64(aFLX), float64(fLLY)+FontSize)
	fLLW, _ := theCtx.ggCtx.MeasureString("field label:")
	fLIX := aFLX + int(fLLW) + 20
	fLIW := dialogWidth - int(fLLW) - 80
	fLIRect := theCtx.drawInput(FD_LabelInput, fLIX, fLLY, fLIW, "", false)

	// field type
	theCtx.ggCtx.SetHexColor("#444")
	_, fTLY := nextVerticalCoords(fLIRect, 15)
	theCtx.ggCtx.DrawString("field type:", float64(aFLX), float64(fTLY)+FontSize)
	fTLW, _ := theCtx.ggCtx.MeasureString("field type:")
	currentX, currentY = aFLX+int(fTLW)+20, fTLY
	for i, field := range supportedFields {
		cBtnId := 300 + i + 1
		cIRect := theCtx.drawCheckbox(cBtnId, currentX, currentY, false)
		cILX, _ := nextHorizontalCoords(cIRect, 10)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(field, float64(cILX), float64(currentY)+FontSize)
		fieldW, _ := theCtx.ggCtx.MeasureString(field)
		currentX = cILX + int(fieldW) + 20
	}

	// for select fields
	theCtx.ggCtx.SetHexColor("#444")
	sFOLY := currentY + 40
	theCtx.ggCtx.DrawString("select field options:", float64(aFLX), float64(sFOLY)+FontSize)
	theCtx.drawTextInput(FD_SelectOptionsInput, aFLX, sFOLY+30, 400, 300, "", false)

	// send the frame to glfw window
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

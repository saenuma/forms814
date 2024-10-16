package main

import (
	"image"
	"slices"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kovidgoyal/imaging"
)

func DrawFormDialog(window *glfw.Window, currentFrame image.Image) {
	FDObjCoords = make(map[int]g143.Rect)

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
	str1 := "Add Form Configuration"
	if IsUpdateDialog {
		str1 = "Edit Form Configuration"
	}
	theCtx.ggCtx.DrawString(str1, float64(aFLX), float64(aFLY)+FontSize)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnRect := theCtx.drawButtonA(FD_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(FD_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

	// name input
	theCtx.ggCtx.SetHexColor("#444")
	fNLY := aFLY + 50
	theCtx.ggCtx.DrawString("field name: ", float64(aFLX), float64(fNLY)+FontSize)
	fNLW, _ := theCtx.ggCtx.MeasureString("field name:")

	var val string
	if IsUpdateDialog {
		val = FormObjects[ToUpdateInstrNum]["name"]
		EnteredTxts[FD_NameInput] = val
	}
	fNIX := aFLX + int(fNLW) + 20
	fNIRect := theCtx.drawInput(FD_NameInput, fNIX, fNLY, 250, val, false)

	currentX, currentY := fNIRect.OriginX+fNIRect.Width+30, fNLY
	// attributes checkboxes
	var cIRect g143.Rect
	for i, attribName := range attributes {
		cBtnId := 200 + i + 1
		if IsUpdateDialog {
			selectedAttribs := strings.Split(FormObjects[ToUpdateInstrNum]["attributes"], ";")
			selected := false
			if slices.Index(selectedAttribs, attribName) != -1 {
				selected = true
			}
			cIRect = theCtx.drawCheckbox(cBtnId, currentX, currentY, selected)
			AttribState[attribName] = selected
		} else {
			cIRect = theCtx.drawCheckbox(cBtnId, currentX, currentY, false)
		}
		cILX, _ := nextHorizontalCoords(cIRect, 10)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(attribName, float64(cILX), float64(currentY)+FontSize)
		fieldW, _ := theCtx.ggCtx.MeasureString(attribName)
		currentX = cILX + int(fieldW) + 20
	}

	// label input
	theCtx.ggCtx.SetHexColor("#444")
	_, fLLY := nextVerticalCoords(fNIRect, 15)
	theCtx.ggCtx.DrawString("field label:", float64(aFLX), float64(fLLY)+FontSize)
	fLLW, _ := theCtx.ggCtx.MeasureString("field label:")
	fLIX := aFLX + int(fLLW) + 20
	fLIW := dialogWidth - int(fLLW) - 80
	var val3 string
	if IsUpdateDialog {
		val3 = FormObjects[ToUpdateInstrNum]["label"]
		EnteredTxts[FD_LabelInput] = val3
	}
	fLIRect := theCtx.drawInput(FD_LabelInput, fLIX, fLLY, fLIW, val3, false)

	// field type
	theCtx.ggCtx.SetHexColor("#444")
	_, fTLY := nextVerticalCoords(fLIRect, 15)
	theCtx.ggCtx.DrawString("field type:", float64(aFLX), float64(fTLY)+FontSize)
	fTLW, _ := theCtx.ggCtx.MeasureString("field type:")
	currentX, currentY = aFLX+int(fTLW)+20, fTLY
	for i, field := range supportedFields {
		cBtnId := 300 + i + 1
		var cIRect g143.Rect
		if IsUpdateDialog {
			SelectedFieldType = FormObjects[ToUpdateInstrNum]["fieldtype"]
			selected := (field == SelectedFieldType)
			cIRect = theCtx.drawCheckbox(cBtnId, currentX, currentY, selected)
		} else {
			cIRect = theCtx.drawCheckbox(cBtnId, currentX, currentY, false)
		}
		cILX, _ := nextHorizontalCoords(cIRect, 10)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(field, float64(cILX), float64(currentY)+FontSize)
		fieldW, _ := theCtx.ggCtx.MeasureString(field)
		newX := cILX + int(fieldW) + 20
		if newX > (dialogOriginX+dialogWidth - 10 - int(fieldW)) {
			currentY += 40
			currentX = dialogOriginX+20+int(fTLW)+20
		} else {
			currentX = newX
		}
	}

	// for select fields
	theCtx.ggCtx.SetHexColor("#444")
	sFOLY := currentY + 40
	theCtx.ggCtx.DrawString("select field options:", float64(aFLX), float64(sFOLY)+FontSize)
	var val2 string
	if IsUpdateDialog {
		val2 = FormObjects[ToUpdateInstrNum]["select_options"]
		EnteredTxts[FD_SelectOptionsInput] = val2
	}
	sOIRect := theCtx.drawTextInput(FD_SelectOptionsInput, aFLX, sFOLY+30, 400, 250, val2, false)

	// for int fields
	theCtx.ggCtx.SetHexColor("#444")
	nFOLX, _ := nextHorizontalCoords(sOIRect, 40)
	theCtx.ggCtx.DrawString("int field options:", float64(nFOLX), float64(sFOLY+FontSize))

	lTLY := sFOLY + 40
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("linked table:", float64(nFOLX), float64(lTLY+FontSize))
	lTLW, _ := theCtx.ggCtx.MeasureString("linked table:")

	var val4 string
	if IsUpdateDialog {
		val4 = FormObjects[ToUpdateInstrNum]["linked_table"]
		EnteredTxts[FD_LinkedTableInput] = val4
	}
	lTIX := nFOLX + int(lTLW) + 20
	lTIRect := theCtx.drawInput(FD_LinkedTableInput, lTIX, lTLY, 250, val4, false)

	// min value input
	_, mVLY := nextVerticalCoords(lTIRect, 10)
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("min value:", float64(nFOLX), float64(mVLY)+FontSize)
	mVLW, _ := theCtx.ggCtx.MeasureString("min value:")

	var val5 string
	if IsUpdateDialog {
		val5 = FormObjects[ToUpdateInstrNum]["min_value"]
		EnteredTxts[FD_MinValueInput] = val5
	}
	mVIX := nFOLX + int(mVLW) + 20
	mVIRect := theCtx.drawInput(FD_MinValueInput, mVIX, mVLY, 250, val5, false)

	// max value input
	_, mV2LY := nextVerticalCoords(mVIRect, 15)
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("max value:", float64(nFOLX), float64(mV2LY)+FontSize)

	var val6 string
	if IsUpdateDialog {
		val6 = FormObjects[ToUpdateInstrNum]["max_value"]
		EnteredTxts[FD_MaxValueInput] = val6
	}
	mV2IX := nFOLX + int(mVLW) + 20
	theCtx.drawInput(FD_MaxValueInput, mV2IX, mV2LY, 250, val6, false)

	// send the frame to glfw window
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

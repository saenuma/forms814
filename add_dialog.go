package main

import (
	"image"
	"slices"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kovidgoyal/imaging"
)

func drawFormDialog(window *glfw.Window, currentFrame image.Image) {
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
	str1 := "Add Form Item"
	if IsUpdateDialog {
		str1 = "Edit Form Item"
	}
	theCtx.ggCtx.DrawString(str1, float64(aFLX), float64(aFLY)+FontSize)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnStr := "Add"
	if IsUpdateDialog {
		addBtnStr = "Edit"
	}
	addBtnRect := theCtx.drawButtonA(FD_AddBtn, addBtnOriginX, dialogOriginY+20, addBtnStr, "#fff", "#333")
	closeBtnX := nextX(addBtnRect, 10)
	theCtx.drawButtonA(FD_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", fontColor, "#aaa")

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
		cILX := nextX(cIRect, 10)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(attribName, float64(cILX), float64(currentY)+FontSize)
		fieldW, _ := theCtx.ggCtx.MeasureString(attribName)
		currentX = cILX + int(fieldW) + 20
	}

	// label input
	theCtx.ggCtx.SetHexColor("#444")
	fLLY := nextY(fNIRect, 15)
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
	fTLY := nextY(fLIRect, 15)
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
		cILX := nextX(cIRect, 10)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(field, float64(cILX), float64(currentY)+FontSize)
		fieldW, _ := theCtx.ggCtx.MeasureString(field)
		newX := cILX + int(fieldW) + 20
		if i == len(supportedFields)-1 {
			break
		}
		if newX > (dialogOriginX+dialogWidth)-200 {
			currentY += 40
			currentX = dialogOriginX + 20 + int(fTLW) + 20
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
		// EnteredTxts[FD_SelectOptionsInput] = val2
	}
	pBY := sFOLY + 30
	sPBRS := theCtx.drawButtonA(FD_SelectPasteBtn, aFLX, pBY, "paste", "#444", "#aaa")
	sEBX := nextX(sPBRS, 20)
	sEBRS := theCtx.drawButtonA(FD_SelectEmptyBtn, sEBX, pBY, "empty", "#444", "#aaa")
	sOIY := nextY(sEBRS, 10)
	sOIRect := theCtx.drawTextInput(FD_SelectOptionsInput, aFLX, sOIY, 400, 200, val2, false)

	// for int fields
	theCtx.ggCtx.SetHexColor("#444")
	nFOLX := nextX(sOIRect, 40)
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
	mVLY := nextY(lTIRect, 10)
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
	mV2LY := nextY(mVIRect, 15)
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

func formDialogMouseBtnCB(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
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
		FD_SelectedInput = 0

		EnteredTxts = map[int]string{
			FD_LabelInput: "", FD_NameInput: "", FD_SelectOptionsInput: "",
			FD_LinkedTableInput: "", FD_MinValueInput: "", FD_MaxValueInput: "",
		}

		drawItemsView(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(itemsViewMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

	case FD_NameInput, FD_LabelInput, FD_LinkedTableInput, FD_MinValueInput, FD_MaxValueInput:
		FD_SelectedInput = widgetCode

		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, EnteredTxts[widgetCode], true)

		// disable other inputs
		allInputs := []int{FD_NameInput, FD_LabelInput, FD_LinkedTableInput, FD_MinValueInput, FD_MaxValueInput}
		index := slices.Index(allInputs, widgetCode)
		leftInputs := slices.Delete(slices.Clone(allInputs), index, index+1)
		for _, inputId := range leftInputs {
			inputRS := FDObjCoords[inputId]
			theCtx.drawInput(inputId, inputRS.OriginX, inputRS.OriginY, inputRS.Width, EnteredTxts[inputId], false)
		}

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case FD_SelectEmptyBtn:
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)

		EnteredTxts[FD_SelectOptionsInput] = ""
		sIRect := FDObjCoords[FD_SelectOptionsInput]
		theCtx.drawTextInput(FD_SelectOptionsInput, sIRect.OriginX, sIRect.OriginY, sIRect.Width,
			sIRect.Height, EnteredTxts[FD_SelectOptionsInput], false)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case FD_SelectPasteBtn:
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)

		tmp := strings.ReplaceAll(glfw.GetClipboardString(), "\r", "")
		var newTmp string
		for _, part := range strings.Split(tmp, "\n") {
			newTmp += strings.TrimSpace(part) + "\n"
		}
		EnteredTxts[FD_SelectOptionsInput] = newTmp
		sIRect := FDObjCoords[FD_SelectOptionsInput]
		theCtx.drawTextInput(FD_SelectOptionsInput, sIRect.OriginX, sIRect.OriginY, sIRect.Width,
			sIRect.Height, EnteredTxts[FD_SelectOptionsInput], false)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case FD_AddBtn:
		for _, obj := range FormObjects {
			if obj["name"] == EnteredTxts[FD_NameInput] && !IsUpdateDialog {
				return
			}
		}

		item := map[string]string{
			"name":      EnteredTxts[FD_NameInput],
			"fieldtype": SelectedFieldType,
		}

		// enforce the keys: name, label, fieldtype
		for _, v := range item {
			if v == "" {
				return
			}
		}

		attribs := make([]string, 0)
		for k, v := range AttribState {
			if v {
				attribs = append(attribs, k)
			}
		}

		if !slices.Contains(attribs, "hidden") && len(EnteredTxts[FD_LabelInput]) == 0 {
			return
		}

		item["label"] = EnteredTxts[FD_LabelInput]

		item["attributes"] = strings.Join(attribs, ";")
		selectFields := []string{"select", "multi_display_select", "single_display_select"}
		if slices.Index(selectFields, SelectedFieldType) != -1 && len(EnteredTxts[FD_SelectOptionsInput]) != 0 {
			item["select_options"] = EnteredTxts[FD_SelectOptionsInput]
		}

		if SelectedFieldType == "int" {
			item["linked_table"] = EnteredTxts[FD_LinkedTableInput]
			item["min_value"] = EnteredTxts[FD_MinValueInput]
			item["max_value"] = EnteredTxts[FD_MaxValueInput]
		}

		if IsUpdateDialog {
			FormObjects[ToUpdateInstrNum] = item
			IsUpdateDialog = false
		} else {
			if IsInsertBeforeDialog {
				FormObjects = slices.Insert(FormObjects, ToInsertBefore, item)
				IsInsertBeforeDialog = false
			} else {
				FormObjects = append(FormObjects, item)
			}
		}

		FD_SelectedInput = 0
		EnteredTxts = map[int]string{
			FD_LabelInput: "", FD_NameInput: "", FD_SelectOptionsInput: "",
			FD_LinkedTableInput: "", FD_MinValueInput: "", FD_MaxValueInput: "",
		}
		AttribState = make(map[string]bool)

		drawItemsView(window)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(itemsViewMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

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

func formDialogKeyCB(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if FD_SelectedInput == FD_NameInput {
		val := EnteredTxts[FD_NameInput]
		if key == glfw.KeyBackspace && len(EnteredTxts[FD_NameInput]) != 0 {
			EnteredTxts[FD_NameInput] = val[:len(val)-1]
		} else if key == glfw.KeyMinus && mods == glfw.ModShift {
			EnteredTxts[FD_NameInput] = val + "_"
		} else {
			EnteredTxts[FD_NameInput] = val + glfw.GetKeyName(key, scancode)
		}

		nIRS := FDObjCoords[FD_NameInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawInput(FD_NameInput, nIRS.OriginX, nIRS.OriginY, nIRS.Width, EnteredTxts[FD_NameInput], true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if FD_SelectedInput == FD_LabelInput {
		val := EnteredTxts[FD_LabelInput]
		if key == glfw.KeyBackspace && len(EnteredTxts[FD_LabelInput]) != 0 {
			EnteredTxts[FD_LabelInput] = val[:len(val)-1]
		}

		sIRect := FDObjCoords[FD_LabelInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawInput(FD_LabelInput, sIRect.OriginX, sIRect.OriginY, sIRect.Width, EnteredTxts[FD_LabelInput], true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if FD_SelectedInput == FD_LinkedTableInput {
		val := EnteredTxts[FD_LinkedTableInput]
		if key == glfw.KeyBackspace && len(val) != 0 {
			EnteredTxts[FD_LinkedTableInput] = val[:len(val)-1]
		} else if key == glfw.KeyMinus && mods == glfw.ModShift {
			EnteredTxts[FD_LinkedTableInput] = val + "_"
		} else {
			EnteredTxts[FD_LinkedTableInput] = val + glfw.GetKeyName(key, scancode)
		}

		nIRS := FDObjCoords[FD_LinkedTableInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawInput(FD_LinkedTableInput, nIRS.OriginX, nIRS.OriginY, nIRS.Width, EnteredTxts[FD_LinkedTableInput], true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if FD_SelectedInput == FD_MinValueInput || FD_SelectedInput == FD_MaxValueInput {
		inputCode := FD_MinValueInput
		if FD_SelectedInput == FD_MaxValueInput {
			inputCode = FD_MaxValueInput
		}
		val := EnteredTxts[inputCode]
		if IsKeyNumeric(key) {
			EnteredTxts[inputCode] = val + glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeyBackspace && len(val) != 0 {
			EnteredTxts[inputCode] = val[:len(val)-1]
		}

		nIRS := FDObjCoords[inputCode]
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawInput(inputCode, nIRS.OriginX, nIRS.OriginY, nIRS.Width, EnteredTxts[inputCode], true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	}

}

func formDialogCharCB(window *glfw.Window, char rune) {
	wWidth, wHeight := window.GetSize()

	if FD_SelectedInput == FD_LabelInput {
		val := EnteredTxts[FD_LabelInput]
		EnteredTxts[FD_LabelInput] = val + string(char)

		sIRect := FDObjCoords[FD_LabelInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawInput(FD_LabelInput, sIRect.OriginX, sIRect.OriginY, sIRect.Width, EnteredTxts[FD_LabelInput], true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	} else if FD_SelectedInput == FD_SelectOptionsInput {
		val := EnteredTxts[FD_SelectOptionsInput]
		EnteredTxts[FD_SelectOptionsInput] = val + string(char)

		sIRect := FDObjCoords[FD_SelectOptionsInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		theCtx.drawTextInput(FD_SelectOptionsInput, sIRect.OriginX, sIRect.OriginY, sIRect.Width,
			sIRect.Height, EnteredTxts[FD_SelectOptionsInput], true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	}

}

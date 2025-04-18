package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tidwall/pretty"
)

func projViewMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
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
	case PROJ_NewProject:
		// create project folder
		rootPath, _ := GetRootPath()

		outPath := filepath.Join(rootPath, NameInputEnteredTxt)
		os.MkdirAll(outPath, 0777)

		// redraw current view
		ProjectName = NameInputEnteredTxt
		DrawBeginView(window, NameInputEnteredTxt)
		window.SetCursorPosCallback(getHoverCB(ProjObjCoords))

	case PROJ_OpenWDBtn:
		rootPath, _ := GetRootPath()
		ExternalLaunch(rootPath)

	case PROJ_NewForm:
		if FNameInputEnteredTxt == "" {
			return
		}

		// create file
		FormName = FNameInputEnteredTxt + ".f8p"
		outPath := filepath.Join(rootPath, ProjectName, FormName)
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		DrawWorkView(window)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		// quick hover effect
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

	case PROJ_NameInput, PROJ_FNameInput:
		PV_SelectedInput = widgetCode

		theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
		theCtx.drawInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, EnteredTxts[widgetCode], true)

		// disable other inputs
		allInputs := []int{PROJ_NameInput, PROJ_FNameInput}
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
		DrawWorkView(window)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		// quick hover effect
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

		DrawWorkView(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		// window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

	case FD_NameInput, FD_LabelInput, FD_SelectOptionsInput, FD_LinkedTableInput, FD_MinValueInput, FD_MaxValueInput:
		FD_SelectedInput = widgetCode

		theCtx := Continue2dCtx(CurrentWindowFrame, &FDObjCoords)
		if widgetCode == FD_SelectOptionsInput {
			theCtx.drawTextInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width,
				widgetRS.Height, EnteredTxts[widgetCode], true)
		} else {
			theCtx.drawInput(widgetCode, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, EnteredTxts[widgetCode], true)
		}

		// disable other inputs
		allInputs := []int{FD_NameInput, FD_LabelInput, FD_SelectOptionsInput,
			FD_LinkedTableInput, FD_MinValueInput, FD_MaxValueInput}
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

		EnteredTxts = map[int]string{
			FD_LabelInput: "", FD_NameInput: "", FD_SelectOptionsInput: "",
			FD_LinkedTableInput: "", FD_MinValueInput: "", FD_MaxValueInput: "",
		}
		AttribState = make(map[string]bool)

		DrawWorkView(window)
		window.SetCursorPosCallback(getHoverCB(WKObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
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

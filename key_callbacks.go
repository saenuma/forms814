package main

import (
	"os"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func ProjKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if PV_SelectedInput == PROJ_NameInput {

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

			// window.SetMouseButtonCallback(workViewMouseBtnCallback)
			// window.SetKeyCallback(nil)
			// window.SetScrollCallback(FirstUIScrollCallback)
			return
		} else {
			NameInputEnteredTxt += glfw.GetKeyName(key, scancode)
		}

		nIRS := ProjObjCoords[PROJ_NameInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
		theCtx.drawInput(PROJ_NameInput, nIRS.OriginX, nIRS.OriginY, nIRS.Width, NameInputEnteredTxt, true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if PV_SelectedInput == PROJ_FNameInput {

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
			DrawWorkView(window)
			window.SetMouseButtonCallback(workViewMouseBtnCallback)
			window.SetKeyCallback(nil)
			// window.SetScrollCallback(FirstUIScrollCallback)
			return
		} else {
			FNameInputEnteredTxt += glfw.GetKeyName(key, scancode)
		}

		fnIRS := ProjObjCoords[PROJ_FNameInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
		theCtx.drawInput(PROJ_FNameInput, fnIRS.OriginX, fnIRS.OriginY, fnIRS.Width, FNameInputEnteredTxt, true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	}
}

func FDKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
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

func FDCharCallback(window *glfw.Window, char rune) {
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

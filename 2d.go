package main

import (
	"image"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
)

type Ctx struct {
	WindowWidth     int
	WindowHeight    int
	ggCtx           *gg.Context
	ObjCoords       *map[int]g143.Rect
	CurrentFontSize int
}

func New2dCtx(wWidth, wHeight int, objCoords *map[int]g143.Rect) Ctx {
	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	ctx := Ctx{WindowWidth: wWidth, WindowHeight: wHeight, ggCtx: ggCtx,
		ObjCoords: objCoords, CurrentFontSize: 20}
	return ctx
}

func Continue2dCtx(img image.Image, objCoords *map[int]g143.Rect) Ctx {
	ggCtx := gg.NewContextForImage(img)

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	ctx := Ctx{WindowWidth: img.Bounds().Dx(), WindowHeight: img.Bounds().Dy(), ggCtx: ggCtx,
		ObjCoords: objCoords, CurrentFontSize: 20}
	return ctx
}

func (ctx *Ctx) setFontSize(fontSize int) {
	// load font
	fontPath := GetDefaultFontPath()
	err := ctx.ggCtx.LoadFontFace(fontPath, float64(fontSize))
	if err != nil {
		panic(err)
	}

	ctx.CurrentFontSize = fontSize
}

func (ctx *Ctx) drawButtonA(btnId, originX, originY int, text, textColor, bgColor string) g143.Rect {
	// draw bounding rect
	textW, textH := ctx.ggCtx.MeasureString(text)
	width, height := textW+float64(ctx.CurrentFontSize), textH+float64(ctx.CurrentFontSize)
	ctx.ggCtx.SetHexColor(bgColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(width), float64(height))
	ctx.ggCtx.Fill()

	textOffsetY := float64(ctx.CurrentFontSize) / 5
	textOffsetX := float64(ctx.CurrentFontSize) / 2
	// draw text
	ctx.ggCtx.SetHexColor(textColor)
	ctx.ggCtx.DrawString(text, float64(originX+int(textOffsetX)), float64(originY+ctx.CurrentFontSize)+textOffsetY)

	// save dimensions
	btnARect := g143.NewRect(originX, originY, int(width), int(height))
	(*ctx.ObjCoords)[btnId] = btnARect
	return btnARect
}

func (ctx *Ctx) drawInput(inputId, originX, originY, inputWidth int, value string, isSelected bool) g143.Rect {
	height := 30
	if isSelected {
		ctx.ggCtx.SetHexColor("#C3983D")
	} else {
		ctx.ggCtx.SetHexColor(fontColor)
	}
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(inputWidth), float64(height))
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#fff")
	ctx.ggCtx.DrawRectangle(float64(originX)+2, float64(originY)+2, float64(inputWidth)-4, float64(height)-4)
	ctx.ggCtx.Fill()

	entryRect := g143.Rect{Width: inputWidth, Height: height, OriginX: originX, OriginY: originY}
	(*ctx.ObjCoords)[inputId] = entryRect

	if len(value) != 0 {
		ctx.ggCtx.SetHexColor("#444")
		ctx.ggCtx.DrawString(value, float64(originX+15), float64(originY)+FontSize)
	}
	return entryRect
}

func (ctx *Ctx) drawCheckbox(inputId, originX, originY int, isSelected bool) g143.Rect {
	width := 30
	height := 30
	ctx.ggCtx.SetHexColor(fontColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(width), float64(height))
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#fff")
	ctx.ggCtx.DrawRectangle(float64(originX)+2, float64(originY)+2, float64(width)-4, float64(height)-4)
	ctx.ggCtx.Fill()

	entryRect := g143.Rect{Width: width, Height: height, OriginX: originX, OriginY: originY}
	(*ctx.ObjCoords)[inputId] = entryRect

	if isSelected {
		ctx.ggCtx.SetHexColor("#444")
		ctx.ggCtx.DrawRectangle(float64(originX)+4, float64(originY)+4, float64(width)-8, float64(height)-8)
		ctx.ggCtx.Fill()
	}
	return entryRect
}

func (ctx *Ctx) drawTextInput(inputId, originX, originY, inputWidth, height int, values string, isSelected bool) g143.Rect {
	if isSelected {
		ctx.ggCtx.SetHexColor("#C3983D")
	} else {
		ctx.ggCtx.SetHexColor(fontColor)
	}
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(inputWidth), float64(height))
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#fff")
	ctx.ggCtx.DrawRectangle(float64(originX)+2, float64(originY)+2, float64(inputWidth)-4, float64(height)-4)
	ctx.ggCtx.Fill()

	entryRect := g143.Rect{Width: inputWidth, Height: height, OriginX: originX, OriginY: originY}
	(*ctx.ObjCoords)[inputId] = entryRect

	if len(strings.TrimSpace(values)) != 0 {
		strs := strings.Split(values, "\n")
		currentY := originY
		for _, str := range strs {
			ctx.ggCtx.SetHexColor("#444")
			ctx.ggCtx.DrawString(str, float64(originX+15), float64(currentY)+FontSize)
			currentY += FontSize + 5
			if currentY >= (originY + height) {
				break
			}
		}
	}
	return entryRect
}

func (ctx *Ctx) windowRect() g143.Rect {
	return g143.NewRect(0, 0, ctx.WindowWidth, ctx.WindowHeight)
}

func nextX(aRect g143.Rect, margin int) int {
	return aRect.OriginX + aRect.Width + margin
}

func nextY(aRect g143.Rect, margin int) int {
	return aRect.OriginY + aRect.Height + margin
}

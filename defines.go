package main

import (
	"image"
	"time"

	g143 "github.com/bankole7782/graphics143"
)

const (
	FPS       = 24
	FontSize  = 20
	fontColor = "#444"
	PageSize  = 20

	PROJ_NameInput  = 11
	PROJ_NewProject = 12

	WK_AddFormBtn = 21
	WK_OpenWDBtn  = 22

	ProgTitle = "forms814: a tool for website developers on records collection"
)

var (
	CurrentWindowFrame image.Image
	Instructions       []map[string]string
	ProjectName        string
	ToUpdateInstrNum   int // for updating an instruction
	IsUpdateDialog     bool
	CurrentPage        int

	ToInsertBefore       int // for inbetween dialogs
	IsInsertBeforeDialog bool

	// view projects
	ProjObjCoords map[int]g143.Rect
	WKObjCoords   map[int]g143.Rect

	NameInputEnteredTxt string

	cursorEventsCount = 0
)

type ToSortProject struct {
	Name    string
	ModTime time.Time
}

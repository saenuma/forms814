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

	FD_AddBtn             = 31
	FD_CloseBtn           = 32
	FD_NameInput          = 33
	FD_LabelInput         = 34
	FD_SelectOptionsInput = 35

	ProgTitle = "forms814: a tool for website developers on records collection"
)

var (
	CurrentWindowFrame image.Image
	ProjectName        string
	ToUpdateInstrNum   int // for updating an instruction
	IsUpdateDialog     bool
	CurrentPage        int

	ToInsertBefore       int // for inbetween dialogs
	IsInsertBeforeDialog bool

	// view projects
	ProjObjCoords map[int]g143.Rect
	WKObjCoords   map[int]g143.Rect
	FDObjCoords   map[int]g143.Rect

	NameInputEnteredTxt string

	cursorEventsCount = 0

	supportedFields = []string{"number", "string", "text", "email", "select", "date",
		"datetime"}
	attributes = []string{"required", "unique", "nindex", "hidden"}

	FD_SelectedInput = 0

	EnteredTxts map[int]string = map[int]string{
		FD_LabelInput: "", FD_NameInput: "", FD_SelectOptionsInput: "",
	}
	AttribState       map[string]bool = make(map[string]bool)
	SelectedFieldType string

	FormObjects []map[string]string
)

type ToSortProject struct {
	Name    string
	ModTime time.Time
}

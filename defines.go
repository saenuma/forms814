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

	BEGINV_NameInput  = 11
	BEGINV_NewProject = 12
	BEGINV_NewForm    = 13
	BEGINV_FNameInput = 14
	BEGINV_OpenWDBtn  = 15

	ITEMSV_AddFormBtn = 21
	ITEMSV_OpenWDBtn  = 22
	ITEMSV_BackBtn    = 23

	FD_AddBtn             = 31
	FD_CloseBtn           = 32
	FD_NameInput          = 33
	FD_LabelInput         = 34
	FD_SelectOptionsInput = 35
	FD_LinkedTableInput   = 36
	FD_MinValueInput      = 37
	FD_MaxValueInput      = 38
	FD_SelectPasteBtn     = 41
	FD_SelectEmptyBtn     = 42

	ProgTitle = "forms814: a tool for website developers on records collection"
)

var (
	CurrentWindowFrame image.Image
	ProjectName        string
	FormName           string
	ToUpdateInstrNum   int // for updating an instruction
	IsUpdateDialog     bool
	CurrentPage        int

	ToInsertBefore       int // for inbetween dialogs
	IsInsertBeforeDialog bool

	// view projects
	ProjObjCoords map[int]g143.Rect
	WKObjCoords   map[int]g143.Rect
	FDObjCoords   map[int]g143.Rect

	NameInputEnteredTxt  string
	FNameInputEnteredTxt string

	cursorEventsCount = 0

	supportedFields = []string{"int", "string", "text", "email", "date", "float",
		"datetime", "select", "multi_display_select", "single_display_select", "check"}
	attributes = []string{"required", "unique", "nindex", "hidden"}

	FD_SelectedInput = 0
	PV_SelectedInput = 0

	EnteredTxts map[int]string = map[int]string{
		FD_LabelInput: "", FD_NameInput: "", FD_SelectOptionsInput: "",
		FD_LinkedTableInput: "", FD_MinValueInput: "", FD_MaxValueInput: "",
	}
	AttribState       map[string]bool = make(map[string]bool)
	SelectedFieldType string

	FormObjects []map[string]string
)

type ToSortProject struct {
	Name    string
	ModTime time.Time
}

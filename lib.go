package main

import (
	"encoding/json"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
)

func GetDefaultFontPath() string {
	fontPath := filepath.Join(os.TempDir(), "f814_font.ttf")
	os.WriteFile(fontPath, DefaultFont, 0777)
	return fontPath
}

func GetRootPath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "os error")
	}

	dd := os.Getenv("SNAP_USER_COMMON")

	if strings.HasPrefix(dd, filepath.Join(hd, "snap", "go")) || dd == "" {
		dd = filepath.Join(hd, "forms814")
		os.MkdirAll(dd, 0777)
	}

	return dd, nil
}

func DoesPathExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func UntestedRandomString(length int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsKeyNumeric(key glfw.Key) bool {
	numKeys := []glfw.Key{glfw.Key0, glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4,
		glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9}

	for _, numKey := range numKeys {
		if key == numKey {
			return true
		}
	}

	return false
}

func SaveProjectCloseCallback(w *glfw.Window) {
	jsonBytes, _ := json.Marshal(Instructions)
	rootPath, _ := GetRootPath()
	outPath := filepath.Join(rootPath, ProjectName)
	os.WriteFile(outPath, jsonBytes, 0777)
}

func GetProjectFiles() []ToSortProject {
	// display some project names
	rootPath, _ := GetRootPath()
	dirEs, _ := os.ReadDir(rootPath)

	projectFiles := make([]ToSortProject, 0)
	for _, dirE := range dirEs {
		if dirE.IsDir() {
			continue
		}

		if strings.HasSuffix(dirE.Name(), ".f8p") {
			fInfo, _ := dirE.Info()
			projectFiles = append(projectFiles, ToSortProject{dirE.Name(), fInfo.ModTime()})
		}
	}

	slices.SortFunc(projectFiles, func(a, b ToSortProject) int {
		return b.ModTime.Compare(a.ModTime)
	})

	return projectFiles
}

func TotalPages() int {
	return int(math.Ceil(float64(len(Instructions)) / float64(PageSize)))
}

func GetPageInstructions(page int) []map[string]string {
	beginIndex := (page - 1) * PageSize
	endIndex := beginIndex + PageSize

	var retInstructions []map[string]string
	if len(Instructions) <= PageSize {
		retInstructions = Instructions
	} else if page == 1 {
		retInstructions = Instructions[:PageSize]
	} else if endIndex > len(Instructions) {
		retInstructions = Instructions[beginIndex:]
	} else {
		retInstructions = Instructions[beginIndex:endIndex]
	}
	return retInstructions
}
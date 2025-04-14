package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/pkg/errors"
)

func phtml(f8pPath string) (string, error) {
	formObjects, err := getFormObjects(f8pPath)
	if err != nil {
		return "", err
	}

	var html string
	for _, obj := range formObjects {
		if slices.Index(strings.Split(obj["attributes"], ";"), "hidden") != -1 {
			continue
		}

		html += "<div>"
		html += fmt.Sprintf("<div><label for='id_%s'>%s</label></div>", obj["name"], obj["label"])

		if obj["fieldtype"] == "int" {
			html += fmt.Sprintf("<input type='number' name='%s' id='id_%s' min='%s' max='%s' ",
				obj["name"], obj["name"], obj["min_value"], obj["max_value"])
			if slices.Index(strings.Split(obj["attributes"], ";"), "required") != -1 {
				html += " required"
			}
			html += "/>"
		} else if obj["fieldtype"] == "float" {
			html += fmt.Sprintf("<input type='number' name='%s' id='id_%s' min='%s' max='%s' step='0.0001'",
				obj["name"], obj["name"], obj["min_value"], obj["max_value"])
			if slices.Index(strings.Split(obj["attributes"], ";"), "required") != -1 {
				html += " required"
			}
			html += "/>"

		} else if slices.Index([]string{"string", "email", "date", "datetime"}, obj["fieldtype"]) != -1 {
			fieldType := obj["fieldtype"]
			if fieldType == "datetime" {
				fieldType += "-local"
			}
			if fieldType == "string" {
				fieldType = "text"
			}
			html += fmt.Sprintf("<input type='%s' name='%s' id='id_%s' ", fieldType,
				obj["name"], obj["name"])
			if slices.Index(strings.Split(obj["attributes"], ";"), "required") != -1 {
				html += " required"
			}
			html += "/>"
		} else if obj["fieldtype"] == "select" {
			html += fmt.Sprintf("<select id='id_%s' name='%s'", obj["name"], obj["name"])
			if slices.Index(strings.Split(obj["attributes"], ";"), "required") != -1 {
				html += " required"
			}
			html += ">"
			for _, opt := range strings.Split(obj["select_options"], "\n") {
				html += "<option>" + opt + "</option>"
			}
			html += "</select>"
		} else if obj["fieldtype"] == "multi_display_select" {
			html += "<div>"
			for _, opt := range strings.Split(obj["select_options"], "\n") {
				html += fmt.Sprintf("<input type='checkbox' id='id_%s' name='%s' value='%s' /> %s", obj["name"],
					obj["name"], opt, opt)
			}
			html += "</div>"
		} else if obj["fieldtype"] == "single_display_select" {
			html += "<div>"
			for _, opt := range strings.Split(obj["select_options"], "\n") {
				html += fmt.Sprintf("<input type='radio' id='id_%s' name='%s' value='%s' /> %s", obj["name"],
					obj["name"], opt, opt)
			}
			html += "</div>"
		} else if obj["fieldtype"] == "text" {
			html += fmt.Sprintf("<textarea id='id_%s' name='%s'", obj["name"], obj["name"])
			if slices.Index(strings.Split(obj["attributes"], ";"), "required") != -1 {
				html += " required"
			}
			html += "></textarea>"
		} else if obj["fieldtype"] == "check" {
			html += fmt.Sprintf("<input type='checkbox' id='id_%s' name='%s' /> %s", obj["name"],
				obj["name"], obj["label"])
		}
		html += "</div>"
	}

	return html, nil
}

func getFormObjects(formObjectPath string) ([]map[string]string, error) {
	if !strings.HasSuffix(formObjectPath, ".f8p") {
		return nil, errors.New("expecting a .f8p file")
	}

	rawJSON, err := os.ReadFile(formObjectPath)
	if err != nil {
		return nil, errors.Wrap(err, "json error")
	}

	formObjects := make([]map[string]string, 0)
	json.Unmarshal(rawJSON, &formObjects)

	return formObjects, nil
}

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var templates map[string]*template.Template

func compileTemplate(relFilePath string, templateData interface{}, doGoFmt bool) string {
	if t, ok := templates[relFilePath]; ok {
		var bt bytes.Buffer
		err := t.Execute(&bt, templateData)
		checkError(err)

		str := bt.String()
		if doGoFmt {
			formatted, err := format.Source([]byte(str))
			checkError(err)
			return string(formatted)
		}
		return str
	} else {
		panic("Template not found in map, path was: " + relFilePath)
	}
}

func parseTemplates() {
	t := map[string]*template.Template{}

	fullTemplatesDir, err := filepath.Abs(templatesDir)
	checkError(err)

	err = filepath.Walk(fullTemplatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".tpl") {
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			p, err := template.New("").Parse(string(fileContent))
			if err != nil {
				return err
			}

			relPath := strings.Trim(strings.Replace(path[len(fullTemplatesDir):], "\\", "/", -1), "/")
			t[relPath] = p
		}

		return nil
	})
	checkError(err)

	templates = t
}

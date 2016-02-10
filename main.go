package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	templatesDir  string = "./templates"
	inputFileFlag        = flag.String("file", "", "The input YAML file")
	distDirFlag          = flag.String("out", "", "The output dir")
)

func handleStartupPanic() {
	if r := recover(); r != nil {
		errStr := ""
		switch t := r.(type) {
		case error:
			errStr = t.Error()
			break
		default:
			errStr = fmt.Sprintf("%#v", r)
		}
		Logger.Error.Printf("ERROR at STARTUP: %s", errStr)
	}
}

func main() {
	InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	defer handleStartupPanic()

	flag.Parse()

	if len(*inputFileFlag) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if len(*distDirFlag) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	yamlSetup := ReadYamlSetupFile(*inputFileFlag)

	parseTemplates()
	for templateRelPath, _ := range templates {
		compiledText := compileTemplate(templateRelPath, yamlSetup, true)

		distRelPath := strings.Replace(templateRelPath, ".go.tpl", ".go", -1)
		filePath := filepath.Join(*distDirFlag, distRelPath)
		fileDir := filepath.Dir(filePath)
		err := os.MkdirAll(fileDir, 0600)
		checkError(err)

		err = ioutil.WriteFile(filePath, []byte(compiledText), 0600)
		checkError(err)
	}
}

package main

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"strings"
)

func ReadYamlSetupFile(filePath string) *YamlSetup {
	fileBytes, err := ioutil.ReadFile(filePath)
	checkError(err)

	y := &yamlSetup{}
	err = yaml.Unmarshal(fileBytes, y)
	checkError(err)

	Y := &YamlSetup{yamlSetup: y}
	Y.validate()
	Y.init()

	return Y
}

type YamlSetup struct {
	*yamlSetup
	HasFlags  bool
	IsService bool
}

func (y *YamlSetup) ContainsService(serviceName string) bool {
	for _, s := range y.Services {
		if strings.TrimSpace(strings.ToLower(s)) == strings.ToLower(serviceName) {
			return true
		}
	}
	return false
}

func (y *YamlSetup) validate() {
	if y.Features != nil && y.Features.Service != nil && strings.Contains(y.Features.Service.Name, " ") {
		panic("The service name cannot contain a space")
	}
}

func (y *YamlSetup) init() {
	y.HasFlags = len(y.Flags) > 0
	y.IsService = y.Features != nil && y.Features.Service != nil
}

type yamlSetup struct {
	GoPackageName string
	Flags         []string
	Services      []string
	Security      []string
	Features      *struct {
		Service *struct {
			Name       string
			HandleStop bool
		}
	}
}

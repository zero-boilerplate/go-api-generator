package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestReadYAML(t *testing.T) {
	Convey("Testing reading of YAML", t, func() {
		yamlSetup := ReadYamlSetupFile("./.example.go-api.yml")

		So(yamlSetup, ShouldNotBeNil)
		Convey("Should contain Mysql service name (case insensitive)", func() {
			So(yamlSetup.ContainsService("Mysql"), ShouldBeTrue)
			So(yamlSetup.ContainsService("Mysql1"), ShouldBeFalse)
			So(yamlSetup.ContainsService("mysql"), ShouldBeTrue)
		})
	})
}

func TestGenerateService(t *testing.T) {
	Convey("Testing generation of app.go file", t, func() {
		parseTemplates()

		yamlSetup := ReadYamlSetupFile("./.example.go-api.yml")

		fileContent := compileTemplate("app.go.tpl", yamlSetup, true)

		expectedLines := []string{
			`package main`,
			``,
			`import (`,
			`    "github.com/ayufan/golang-kardianos-service"`,
			`)`,
			``,
			`type app struct {`,
			`    logger service.Logger`,
			`}`,
			``,
			`func (a *app) Run(logger service.Logger) {`,
			`    a.logger = logger`,
			`    a.RunApp()`,
			`}`,
			``,
			`func (a *app) OnStop() {`,
			`    // Any work in Stop should be quick, usually a few seconds at most.`,
			`    a.StopApp()`,
			`}`,
		}

		lines := strings.Split(strings.TrimSpace(fileContent), "\n")
		for i, el := range expectedLines {
			So(strings.TrimSpace(el), ShouldEqual, strings.TrimSpace(lines[i]))
		}
	})
}

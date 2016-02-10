package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCompileTemplate(t *testing.T) {
	Convey("Testing compiling of templates", t, func() {
		Convey("Count the TOTAL number of templates", func() {
			parseTemplates()
			So(len(templates), ShouldEqual, 4)
		})

		yamlSetup := ReadYamlSetupFile("./.example.go-api.yml")

		var compiled string

		Convey("Test main.go.tpl", func() {
			compiled = compileTemplate("main.go.tpl", yamlSetup, true)
			So(compiled, ShouldContainSubstring, "package main")
			So(compiled, ShouldContainSubstring, `"flag"`)
			So(compiled, ShouldContainSubstring, `myflag1Flag = flag.String("myflag1", "", "The value of myflag1")`)
			So(compiled, ShouldContainSubstring, `awesome2Flag = flag.String("awesome2", "", "The value of awesome2")`)
		})

		Convey("Test service.go.tpl", func() {
			compiled = compileTemplate("service.go.tpl", yamlSetup, true)
			So(compiled, ShouldContainSubstring, "package main")
			So(compiled, ShouldContainSubstring, `NewServiceRunnerBuilder("MyService123", a).`)
		})
	})
}

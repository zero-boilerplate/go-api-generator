package {{.GoPackageName}}

{{if .IsService}}

import (
    "github.com/ayufan/golang-kardianos-service"
)

type app struct {
    logger service.Logger
}

func (a *app) Run(logger service.Logger) {
    a.logger = logger
    a.RunApp()
}

{{if .Features.Service.HandleStop}}
func (a *app) OnStop() {
    // Any work in Stop should be quick, usually a few seconds at most.
    a.StopApp()
}
{{end}}

{{end}}
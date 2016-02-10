package {{.GoPackageName}}

{{if .IsService}}

import (
    "github.com/zero-boilerplate/go-api-helpers/service"
)

func InitService() {
    a := &app{}
    service.
        NewServiceRunnerBuilder("{{.Features.Service.Name}}", a).
        {{if .Features.Service.HandleStop}} WithOnStopHandler(a). {{end}}
        Run()
}

{{end}}
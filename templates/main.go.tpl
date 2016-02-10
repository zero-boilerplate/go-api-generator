package {{.GoPackageName}}

import (
    "github.com/francoishill/golang-web-dry/errors/stacktraces/prettystacktrace"
    "log"
    {{if .HasFlags}}"flag"{{end}}
)

{{if .HasFlags}}var (
{{range $flag := .Flags}}
    {{$flag}}Flag = flag.String("{{$flag}}", "", "The value of {{$flag}}")
{{end}}
)
{{end}}

func handleStartupPanic() {
    if r := recover(); r != nil {
        log.Fatalf("ERROR: %s. Stack: %s\n", getErrorStringFromRecovery(r), prettystacktrace.GetPrettyStackTrace())
    }
}

func main() {
    defer handleStartupPanic()

    InitService()
}

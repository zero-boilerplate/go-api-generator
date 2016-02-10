package {{.GoPackageName}}

import (
    "fmt"
)

func getErrorStringFromRecovery(r interface{}) string {
    str := ""
    switch t := r.(type) {
    case error:
        str = t.Error()
        break
    default:
        str = fmt.Sprintf("%#v", r)
    }
    return str
}
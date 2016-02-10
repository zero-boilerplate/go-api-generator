# go-api-generator
A generator for easily generating new backend (api) golang apps based on an input config (yaml) file

**WORK IN PROGRESS**

## Goals

To have a simple YAML file to specify what our generator app must include. For example a YAML like this:

```
Name: My App
Services:
  - Mysql
  - Redis
  - Sendgrid
  - GoogleCloudStorage
Security:
  - JWT
Features:
  - Service # run as a service - github.com/ayufan/golang-kardianos-service
```

## Current status

A very rough implementation. Currently will generate an app according to the `-file` flag of the YAML file. Generates based on these properties, look at the `.example.go-api.yml` file.

- `GoPackageName` will be used in all the output files as `package ???`
- `Flags` will be used to generate string flag variables in the `main.go` file
- `Features.Service.Name` is used when your app is installed as a service with `-service install`

### Usage

So to make use of this tool (replace the DEST_DIR with your own):
```
go get -u github.com/zero-boilerplate/go-api-generator
cd "$GOPATH/src/github.com/zero-boilerplate/go-api-generator"
go-api-generator -file .example.go-api.yml -out DEST_DIR
```

Now before calling `go build` inside your DEST_DIR note that you might need to implement these two methods in a new file. The method `StopApp` will only be required if your YAML `Features.Service.HandleStop` is `true`.

For example lets create file `my_app.go` in your DEST_DIR:

```
package main

import (
    "time"
)

var mustExit bool = false

func (a *app) RunApp() {
    for !mustExit {
        sleepDuration := 5*time.Second
        a.logger.Infof("Continuous loop, will now sleep for %s", sleepDuration)
        time.Sleep(sleepDuration)
    }
}

func (a *app) StopApp() {
    mustExit = true
}
```

#### Install your app as a service


Using your DEST_DIR from above.

```
cd DEST_DIR
go build -o=main
main -service install
main -service start
```

Now check out your service logs (in Windows would be Event Viewer) for the logs printed out "Continuous loop, will now sleep for". To stop the service just call `main -service stop` and to uninstall `main -service uninstall`.
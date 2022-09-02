package main

import (
	"fileserver/pkg/controller"
	"fileserver/pkg/repository"
	"fileserver/pkg/service"
	"fmt"
	"github.com/fakefloordiv/indigo"
	methods "github.com/fakefloordiv/indigo/http/method"
	"github.com/fakefloordiv/indigo/http/status"
	"github.com/fakefloordiv/indigo/router/simple"
	"github.com/fakefloordiv/indigo/types"
	"html/template"
	"log"
)

var addr = "localhost:9090"

func RunFileServer(addr, root string) error {
	tmpl, err := template.ParseFiles("resource/dir.html", "resource/header.html")
	if err != nil {
		return err
	}
	fsRepo := repository.NewFileSystemRepo()
	fsService := service.NewFSService(fsRepo, tmpl)

	indigoController := controller.NewHTTPController(root, fsService)

	fmt.Println("Listening on", addr)

	return runHTTPServer(addr, indigoController)
}

func runHTTPServer(addr string, controller controller.HTTPController) error {
	r := simple.NewRouter(func(request *types.Request) types.Response {
		switch request.Method {
		case methods.GET:
			return controller.DisplayPage(request)
		default:
			return types.WithResponse.
				WithCode(status.MethodNotAllowed).
				WithBody(`<h1 align="center">405 Method Not Allowed</h1>`)
		}
	})

	app := indigo.NewApp(addr)

	return app.Serve(r)
}

func main() {
	// TODO: parse command line options to receive a root from there
	log.Fatal(RunFileServer(addr, "."))
}

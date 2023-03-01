package main

import (
	"fmt"
	"github.com/fakefloordiv/fileserver/pkg/controller"
	"github.com/fakefloordiv/fileserver/pkg/model"
	"github.com/fakefloordiv/fileserver/pkg/repository"
	"github.com/fakefloordiv/fileserver/pkg/service"
	"github.com/indigo-web/indigo/v2"
	"github.com/indigo-web/indigo/v2/http"
	methods "github.com/indigo-web/indigo/v2/http/method"
	"github.com/indigo-web/indigo/v2/http/status"
	"github.com/indigo-web/indigo/v2/router/simple"
	"html/template"
	"log"
)

var addr = "localhost:8080"

func RunFileServer(addr, root string) error {
	dirTemplate, err := template.ParseFiles("resource/dir.html", "resource/header.html")
	if err != nil {
		return err
	}
	fileTemplate, err := template.ParseFiles("resource/file.html")
	if err != nil {
		return err
	}
	templates := model.NewTemplates(dirTemplate, fileTemplate)
	fsRepo := repository.NewFileSystemRepo()
	fsService := service.NewFSService(fsRepo, templates)

	indigoController := controller.NewHTTPController(root, fsService)

	fmt.Println("Listening on", addr)

	return runHTTPServer(addr, indigoController)
}

func runHTTPServer(addr string, controller controller.HTTPController) error {
	r := simple.NewRouter(func(request *http.Request) http.Response {
		switch request.Method {
		case methods.GET:
			return controller.DisplayPage(request)
		default:
			return http.RespondTo(request).
				WithCode(status.MethodNotAllowed).
				WithBody(`<h1 align="center">405 Method Not Allowed</h1>`)
		}
	}, func(request *http.Request, err error) http.Response {
		return http.RespondTo(request).WithError(err)
	})

	app := indigo.NewApp(addr)

	return app.Serve(r)
}

func main() {
	// TODO: parse command line options to receive a root from there
	log.Fatal(RunFileServer(addr, "."))
}

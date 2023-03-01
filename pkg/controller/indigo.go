package controller

import (
	"github.com/fakefloordiv/fileserver/pkg/service"
	"github.com/indigo-web/indigo/v2/http"
	"github.com/indigo-web/indigo/v2/http/status"
	"io"
	"path"
)

type HTTPController interface {
	DisplayPage(request *http.Request) http.Response
}

type httpController struct {
	fsService service.FSService

	rootPath string
}

func NewHTTPController(root string, fsService service.FSService) HTTPController {
	return httpController{
		fsService: fsService,
		rootPath:  root,
	}
}

func (h httpController) DisplayPage(request *http.Request) http.Response {
	resp, err := http.RespondTo(request).WithWriter(func(writer io.Writer) error {
		return h.fsService.RenderPage(path.Join(h.rootPath, request.Path), writer)
	})

	if err != nil {
		return http.RespondTo(request).
			WithCode(status.BadRequest).
			WithBody(err.Error())
	}

	return resp
}

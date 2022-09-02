package controller

import (
	"github.com/fakefloordiv/fileserver/pkg/service"
	"github.com/fakefloordiv/indigo/http/status"
	"github.com/fakefloordiv/indigo/types"
	"path"
)

type HTTPController interface {
	DisplayPage(request *types.Request) types.Response
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

func (h httpController) DisplayPage(request *types.Request) types.Response {
	page, err := h.fsService.RenderPage(path.Join(h.rootPath, request.Path))
	if err != nil {
		return types.WithResponse.
			WithCode(status.BadRequest).
			WithBody(err.Error())
	}

	return types.WithResponse.WithBodyByte(page)
}

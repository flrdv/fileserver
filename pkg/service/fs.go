package service

import (
	"github.com/fakefloordiv/fileserver/pkg/model"
	"github.com/fakefloordiv/fileserver/pkg/repository"
	"io"
	"path"
)

type FSService interface {
	RenderPage(path string, writer io.Writer) error
}

type fsService struct {
	fsRepo    repository.FileSystemRepo
	templates model.Templates
	root      string
}

func NewFSService(
	fsRepo repository.FileSystemRepo, templates model.Templates, root string,
) FSService {
	return fsService{
		fsRepo:    fsRepo,
		templates: templates,
		root:      root,
	}
}

func (f fsService) RenderPage(relative string, writer io.Writer) error {
	absolute := path.Join(f.root, relative)
	isFile, err := f.fsRepo.IsFile(absolute)
	if err != nil {
		return err
	}

	if isFile {
		content, err := f.fsRepo.ReadFile(absolute)
		if err != nil {
			return err
		}

		filename := path.Base(absolute)

		return f.templates.File.Execute(writer, model.NewFile(filename, content))
	}

	entries, err := f.fsRepo.ListDir(f.root, relative)
	if err != nil {
		return err
	}

	return f.templates.Dir.Execute(writer, entries)
}

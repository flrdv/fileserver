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
}

func NewFSService(fsRepo repository.FileSystemRepo, templates model.Templates) FSService {
	return fsService{
		fsRepo:    fsRepo,
		templates: templates,
	}
}

func (f fsService) RenderPage(p string, writer io.Writer) error {
	isFile, err := f.fsRepo.IsFile(p)
	if err != nil {
		return err
	}

	if isFile {
		content, err := f.fsRepo.ReadFile(p)
		if err != nil {
			return err
		}

		filename := path.Base(p)

		return f.templates.File.Execute(writer, model.NewFile(filename, content))
	}

	entries, err := f.getEntries(p)
	if err != nil {
		return err
	}

	return f.templates.Dir.Execute(writer, entries)
}

func (f fsService) getEntries(path string) (model.FileSystemEntries, error) {
	objs, err := f.fsRepo.ListDir(path)

	return model.NewFileSystemEntry(f.fsRepo.GetParentDir(path), objs), err
}

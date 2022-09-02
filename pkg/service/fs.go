package service

import (
	"github.com/fakefloordiv/fileserver/internal"
	"github.com/fakefloordiv/fileserver/pkg/model"
	"github.com/fakefloordiv/fileserver/pkg/repository"
	"html/template"
)

type FSService interface {
	RenderPage(path string) (page []byte, err error)
}

type fsService struct {
	fsRepo   repository.FileSystemRepo
	template *template.Template
}

func NewFSService(fsRepo repository.FileSystemRepo, template *template.Template) FSService {
	return fsService{
		fsRepo:   fsRepo,
		template: template,
	}
}

func (f fsService) RenderPage(path string) (page []byte, err error) {
	isFile, err := f.fsRepo.IsFile(path)
	if err != nil {
		return nil, err
	}

	if isFile {
		return f.fsRepo.ReadFile(path)
	}

	entries, err := f.getEntries(path)
	if err != nil {
		return nil, err
	}

	wr := new(internal.Writer)
	err = f.template.Execute(wr, entries)

	return wr.Content(), err
}

func (f fsService) getEntries(path string) (model.FileSystemEntries, error) {
	objs, err := f.fsRepo.ListDir(path)

	return model.NewFileSystemEntry(f.fsRepo.GetParentDir(path), objs), err
}

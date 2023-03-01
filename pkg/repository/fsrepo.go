package repository

import (
	"github.com/fakefloordiv/fileserver/pkg/model"
	"os"
	"path"
)

type FileSystemRepo interface {
	ListDir(root, path string) ([]model.FileSystemObject, error)
	ReadFile(path string) (string, error)
	IsFile(path string) (bool, error)
	GetParentDir(path string) string
}

type fileSystemRepo struct {
}

func NewFileSystemRepo() FileSystemRepo {
	return fileSystemRepo{}
}

func (f fileSystemRepo) ListDir(root, relative string) ([]model.FileSystemObject, error) {
	entries, err := os.ReadDir(path.Join(root, relative))
	if err != nil {
		return nil, err
	}

	objects := make([]model.FileSystemObject, 0, len(entries))

	for _, entry := range entries {
		objects = append(objects, model.NewFileSystemObject(
			!entry.IsDir(), entry.Name(), path.Join(relative, entry.Name()),
		))
	}

	return objects, nil
}

func (f fileSystemRepo) ReadFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	return string(file), err
}

func (f fileSystemRepo) IsFile(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return !stat.IsDir(), nil
}

func (f fileSystemRepo) GetParentDir(child string) string {
	return path.Dir(child)
}

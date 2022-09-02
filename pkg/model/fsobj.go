package model

type FileSystemObject struct {
	isFile bool
	name   string
	path   string
}

func NewFileSystemObject(isFile bool, name, path string) FileSystemObject {
	return FileSystemObject{
		isFile: isFile,
		name:   name,
		path:   path,
	}
}

func (f FileSystemObject) IsFile() bool {
	return f.isFile
}

func (f FileSystemObject) Name() string {
	return f.name
}

func (f FileSystemObject) Path() string {
	return f.path
}

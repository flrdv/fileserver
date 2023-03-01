package model

type FileSystemObject struct {
	isFile     bool
	name, path string
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

type File struct {
	name, content string
}

func NewFile(name, content string) File {
	return File{
		name:    name,
		content: content,
	}
}

func (f File) Name() string {
	return f.name
}

func (f File) Content() string {
	return f.content
}

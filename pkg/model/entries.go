package model

type FileSystemEntries struct {
	Back    string
	Entries []FileSystemObject
}

func NewFileSystemEntry(back string, entries []FileSystemObject) FileSystemEntries {
	return FileSystemEntries{
		Back:    back,
		Entries: entries,
	}
}

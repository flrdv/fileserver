package model

import "html/template"

type Templates struct {
	Dir  *template.Template
	File *template.Template
}

func NewTemplates(dir, file *template.Template) Templates {
	return Templates{
		Dir:  dir,
		File: file,
	}
}

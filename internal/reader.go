package internal

type Writer struct {
	content []byte
}

func (w *Writer) Write(b []byte) (n int, err error) {
	w.content = append(w.content, b...)

	return len(b), nil
}

func (w Writer) Content() []byte {
	return w.content
}

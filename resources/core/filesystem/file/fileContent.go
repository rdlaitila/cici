package file

import "encoding/json"

type fileContent struct {
	Path     string
	Length   int64
	Content  string
	Encoding string
}

func (t *fileContent) ToJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

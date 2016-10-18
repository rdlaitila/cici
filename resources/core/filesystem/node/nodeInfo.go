package node

import (
	"cici/hateoas"
	"encoding/json"
	"time"
)

type nodeInfo struct {
	Path     string
	Name     string
	Size     int64
	ModTime  time.Time
	Mode     string
	IsDir    bool
	Children []*nodeInfo
	Links    []*hateoas.Link
}

func (t *nodeInfo) ToJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

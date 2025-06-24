package entities

import (
	"crypto/sha1"
	"encoding/base64"
)

type Link struct {
	data string
}

func NewLink(data string) *Link {
	return &Link{data: data}
}

func (l Link) Shorter() string {
	hasher := sha1.New()

	hasher.Write([]byte(l.data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (l Link) GetData() string {
	return l.data
}

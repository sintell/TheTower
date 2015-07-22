package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	NAME_BASE = "Character"
)

type Name struct {
	Text string
}

func NewName(key string) *Name {
	h := md5.New()
	h.Write([]byte(key))
	textHash := hex.EncodeToString(h.Sum(nil))

	return &Name{fmt.Sprintf("%s#%s", NAME_BASE, textHash)}
}

func (this *Name) String() string {
	return this.Text
}

package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	NAME_BASE = "Character"
)

type Name struct {
	Text string
}

func NewName(key string, nameBase ...string) *Name {
	h := md5.New()
	h.Write([]byte(key))
	h.Write([]byte(fmt.Sprintf("%d", time.Now().Nanosecond())))
	textHash := hex.EncodeToString(h.Sum(nil))

	namePrefix := NAME_BASE

	if len(nameBase) == 1 {
		namePrefix = nameBase[0]
	}

	return &Name{fmt.Sprintf("%s#%s", namePrefix, textHash)}
}

func (this *Name) String() string {
	return this.Text
}

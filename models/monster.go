package models

import (
	"github.com/jinzhu/gorm"
)

type Monster struct {
	gorm.Model
	Attributes
	Stats
	Level uint   `json:"level"`
	Name  string `json:"name"      sql:"unique"`
	Class string `json:"class"`
}

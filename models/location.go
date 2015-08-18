package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model `json:"-"`
	Name       string               `json:"name"`
	Key        string               `json:"key"`
	Type       string               `json:"type"`
	Creatures  map[string]*Creature `json:"-"`
}

type LocationInfo struct {
	Name           string
	Type           string
	CreaturesCount uint
}

const STARTING_LOCATION = "mirage_bay"

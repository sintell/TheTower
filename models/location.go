package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Name       string               `json:"name"`
	Type       string               `json:"type"`
	Creatures  map[string]*Creature `json:"-"`
	Characters map[uint]*Character  `json:"-"`
}

const STARTING_LOCATION = "mirage_bay"

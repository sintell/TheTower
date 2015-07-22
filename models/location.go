package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Creatures  []Creature `json:"creatures;omitempty"`
	Characters []string   `json:"characters;omitempty"`
}

package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Item struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Number      uint      `json:"number"`
	Value       uint      `json:"value"`
	AcquiredAt  time.Time `json:"acquiredAt;omitempty"`
}

package models

import (
	"time"
)

type Effect interface {
	GetAttributes() Attributes
	GetStats() Stats
}

type Buff struct {
	Attributes
	Stats
	Duration    time.Duration
	Description string
	Name        string
}

type Debuff struct {
	Attributes
	Stats
	Duration    time.Duration
	Description string
	Name        string
}

func (b *Buff) GetAttributes() Attributes {
	return b.Attributes
}

func (b *Buff) GetStats() Stats {
	return b.Stats
}

func (d *Debuff) GetAttributes() Attributes {
	return d.Attributes
}

func (d *Debuff) GetStats() Stats {
	return d.Stats
}

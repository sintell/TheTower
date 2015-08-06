package models

type Spell struct {
	Id      uint
	Name    string
	Trigger string
	Effect  string

	MpCost   float32
	Cooldown float32
	CastTime float32
}

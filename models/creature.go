package models

type Creature interface {
	ApplyEffect(...*Effect)
}

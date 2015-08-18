package models

type CreatureType uint8

const (
	PLAYER CreatureType = 1 << iota
	NPC
)

type Creature interface {
	Type() CreatureType
	ApplyEffect(...Effect)
	CreatureStats() *Stats
	CreatureAttributes() *Attributes
	CreatureUid() string
	CreatureLocation() string
	NextAbility() Effect
}

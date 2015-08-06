package models

import (
	"time"
)

// Effect activation time
const (
	EFFECT_APPLIES_AT_START EffectActivationTime = 1 << iota
	EFFECT_APPLIES_AT_END
)

// Effect type
const (
	EFFECT_MOVE EffectType = 1 << iota
	EFFECT_ATTRIBUTES
	EFFECT_DAMAGE
	EFFECT_HEAL
)

type EffectActivationTime uint8
type EffectType uint8

type Effect interface {
	Target() *Character
	Type() EffectType
	ActivationTime() EffectActivationTime
	Start() uint64
	Duration() time.Duration
	Ticks() uint
	Delta() map[string]interface{}
}

type Ability struct {
	target         *Character
	abilityType    EffectType
	activationTime EffectActivationTime
	duration       time.Duration
	castStart      uint64
	ticks          uint
	Description    string
	Name           string
}

func (this *Ability) Target() *Character {
	return this.target
}

func (this *Ability) Type() EffectType {
	return this.abilityType
}

func (this *Ability) ActivationTime() EffectActivationTime {
	return this.activationTime
}

func (this *Ability) Duration() time.Duration {
	return this.duration
}

func (this *Ability) Ticks() uint {
	return this.ticks
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

func HandleMoveEffect(effect *Effect) {

}

func HandleAttributesEffect(effect *Effect) {

}

func HandleDamageEffect(effect *Effect) {

}

func HandleHealEffect(effect *Effect) {

}

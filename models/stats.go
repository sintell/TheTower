package models

type Stats struct {
	Hp          float32 `json:"hp"`
	Mp          float32 `json:"mp"`
	MaxHp       float32 `json:"maxHp"`
	MaxMp       float32 `json:"maxMp"`
	HitChanse   float32 `json:"hitChanse"`
	CritChanse  float32 `json:"critChanse"`
	DodgeChanse float32 `json:"dodgeChance"`
	Lidership   int     `json:"lidership"`
	Tradnig     int     `json:"trading"`
}

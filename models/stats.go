package models

type Stats struct {
	HP          float32 `json:"hp"`
	MP          float32 `json:"mp"`
	MaxHP       float32 `json:"maxHp"`
	MaxMP       float32 `json:"maxMp"`
	HitChanse   float32 `json:"hitChanse"`
	CritChanse  float32 `json:"critChanse"`
	DodgeChanse float32 `json:"dodgeChance"`
	Lidership   int     `json:"lidership"`
	Tradnig     int     `json:"trading"`
}

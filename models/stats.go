package models

type Stats struct {
	HP          int     `json:"hp"`
	MP          int     `json:"mp"`
	MaxHP       int     `json:"maxHp"`
	MaxMP       int     `json:"maxMp"`
	HitChanse   float32 `json:"hitChanse"`
	CritChanse  float32 `json:"critChanse"`
	DodgeChanse float32 `json:"dodgeChance"`
	Lidership   int     `json:"lidership"`
	Tradnig     int     `json:"trading"`
}

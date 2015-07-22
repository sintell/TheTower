package models

type Attributes struct {
	Strength int `json:"str" sql:""`
	Agility  int `json:"agi" sql:""`
	Intelect int `json:"int" sql:""`
	Vitality int `json:"vit" sql:""`
	Luck     int `json:"luk" sql:""`
}

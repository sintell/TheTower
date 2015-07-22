package models

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/sintell/mmo-server/utils"
	"math"
	"strings"
)

type Character struct {
	gorm.Model
	Attributes
	Stats
	Level   uint   `json: "level"`
	Name    string `json:"name" sql:"unique"`
	Class   string `json:"class"`
	UserID  uint
	Effects []*Effect `sql:"-"`
	Owner   *User     `json:"-" sql:"-"`
}

type CClass uint8

var defaultAttributes map[string]Attributes
var defaultStats map[string]Stats

func init() {
	err := utils.LoadSetting(&defaultAttributes, "data/defaultAttributes.json")
	if err != nil {
		glog.Error("Can't load default attributes. Reason:", err.Error())
		panic(err.Error())
	}
	glog.Infof("Loaded default attributes: %i\n Total: %d", defaultAttributes, len(defaultAttributes))

	err = utils.LoadSetting(&defaultStats, "data/defaultStats.json")
	if err != nil {
		glog.Error("Can't load default stats. Reason:", err.Error())
		panic(err.Error())
	}
	glog.Infof("Loaded default stats: %i\n Total: %d", defaultStats, len(defaultStats))
}

func (this *Character) SetDefaults() {
	this.Stats = defaultStats[strings.ToLower(this.Class)]
	this.Attributes = defaultAttributes[strings.ToLower(this.Class)]

	this.HP = this.MaxHP
	this.MP = this.MaxMP

	this.RecalculateAttributes()

	glog.Infof("Setting default stats for character %i:\nS : %i\nA : %i", this, defaultStats[this.Class], defaultAttributes[this.Class])
}

func (this *Character) ApplyEffect(effects ...*Effect) Character {
	return Character{}
}

func (this *Character) RecalculateAttributes() {
	this.RecalculateHp()
	this.RecalculateMp()
}

func (this *Character) RecalculateHp() {
	difference := this.HP / this.MaxHP * 100

	this.MaxHP = this.Strength*(0.25*defaultAttrsibutes[strings.ToLower(this.Class)].Strength) +
		-20*math.Cos(this.Level/60*(math.Pi/2)) + 20

	this.HP = this.MaxHP * difference / 100
}

func (this *Character) RecalculateMp() {
	difference := this.MP / this.MaxMP * 100

	this.MaxMP = this.Intelect*(0.25*defaultAttrsibutes[strings.ToLower(this.Class)].Intelect) +
		-40*Math.cos(this.Level/60*(Math.PI/2)) + 40

	this.MP = this.MaxMP * difference / 100
}

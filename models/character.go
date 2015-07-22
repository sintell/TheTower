package models

import (
	"flag"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/sintell/mmo-server/utils"
)

type Character struct {
	gorm.Model
	Attributes
	Stats
	Level   uint   `json: "level"`
	Name    string `json:"name" 		sql:"unique"`
	Class   string `json:"class"`
	UserID  uint
	Effects []*Effect `sql:"-"`
}

type CClass uint8

var defaultAttributes map[string]Attributes
var defaultStats map[string]Stats

func init() {
	flag.Parse()

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
	this.Stats = defaultStats[this.Class]
	this.Attributes = defaultAttributes[this.Class]

	this.HP = this.MaxHP
	this.MP = this.MaxMP
}

func (c *Character) ApplyEffect(effects ...*Effect) Character {
	return Character{}
}

package models

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/sintell/mmo-server/utils"
	"math"
	"strings"
)

type Character struct {
	gorm.Model `json:"-"`
	Attributes
	Stats
	characterType   CreatureType
	Level           uint   `json: "level"`
	Name            string `json:"name" sql:"unique"`
	Class           string `json:"class"`
	CurrentLocation string `json:"homeLocation" sql:"default:'mirage_bay'"`
	UserID          uint
	OwnerUid        string
	AbilityQueue    []Effect `json:"-" sql:"-"`
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

func (this *Character) Type() CreatureType {
	return this.characterType
}

func (this *Character) CreatureStats() *Stats {
	return &this.Stats
}

func (this *Character) CreatureAttributes() *Attributes {
	return &this.Attributes
}

func (this *Character) CreatureUid() string {
	return this.OwnerUid
}

func (this *Character) CreatureLocation() string {
	return this.CurrentLocation
}

func (this *Character) NextAbility() Effect {
	if len(this.AbilityQueue) == 0 {
		return nil
	}
	effect := this.AbilityQueue[0]
	this.AbilityQueue = this.AbilityQueue[1:]
	return effect
}

func (this *Character) ShortData() map[string]interface{} {
	shortData := make(map[string]interface{})

	shortData["Name"] = this.Name
	shortData["Level"] = this.Level
	shortData["Class"] = this.Class

	return shortData
}

func (this *Character) SetDefaults() {
	this.Stats = defaultStats[strings.ToLower(this.Class)]
	this.Attributes = defaultAttributes[strings.ToLower(this.Class)]
	if this.CurrentLocation == "" {
		this.CurrentLocation = STARTING_LOCATION
	}

	this.Hp = this.MaxHp
	this.Mp = this.MaxMp

	this.RecalculateStats()

	glog.V(10).Infof("Setting default stats for character %i:\nS : %i\nA : %i", this, defaultStats[this.Class], defaultAttributes[this.Class])
}

func (this *Character) RecalculateStats() {
	this.RecalculateMaxHp()
	this.RecalculateMaxMp()
}

func (this *Character) RecalculateMaxHp() {
	difference := this.Hp / this.MaxHp * 100.0

	this.MaxHp = this.MaxHp + float32(this.Strength)*(0.25*float32(defaultAttributes[strings.ToLower(this.Class)].Strength)) +
		-20.0*float32(math.Cos(float64(this.Level)/60.0*(math.Pi/2))) + 20.0

	this.Hp = this.MaxHp * difference / 100.0
}

func (this *Character) RecalculateMaxMp() {
	difference := this.Mp / this.MaxMp * 100.0

	this.MaxMp = this.MaxMp + float32(this.Intelect)*(0.25*float32(defaultAttributes[strings.ToLower(this.Class)].Intelect)) +
		-40.0*float32(math.Cos(float64(this.Level)/60.0*(math.Pi/2))) + 40.0

	this.Mp = this.MaxMp * difference / 100.0
}

func (this *Character) ApplyEffect(effects ...Effect) {
}

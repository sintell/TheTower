package models

import (
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/utils"
	"time"
)

type World struct {
	Locations map[string]Location
	Name      string
	time      time.Time
	ticks     int64
}

func NewWorld() *World {
	world := &World{Name: NewName("World", "World").String()}
	world.Locations = make(map[string]Location)

	err := utils.LoadSetting(&world.Locations, "data/locations.json")
	if err != nil {
		panic(err.Error())
	}

	for _, loc := range world.Locations {
		loc.Characters = make(map[uint]*Character)
	}

	return world
}

func (this *World) Tick(t time.Time) {
	this.ticks += 1
	if this.ticks%500 == 0 {
		for _, loc := range this.Locations {
			if charNumbers := len(loc.Characters); charNumbers > 0 {
				glog.Infof("Tick#%d: there are %d characters online in location %s", this.ticks, charNumbers, loc.Name)
			} else {
				glog.Infof("Tick#%d: there are no characters online, %s", this.ticks, loc.Name)
			}
		}
	}
}

func (this *World) MoveCharacter(character *Character, locId string) {
	glog.Infof("Moving character to %s", locId)
	loc := this.Locations[locId]

	if loc.Characters == nil {
		loc.Characters = make(map[uint]*Character)
		this.Locations[locId] = loc
	}
	loc.Characters[character.UserID] = character
}

func (this *World) RemoveCharacter(userId uint) {
	for _, loc := range this.Locations {
		if _, exists := loc.Characters[userId]; exists {
			glog.Warningf("Removing character from world: %s", userId)
			delete(loc.Characters, userId)
		}

	}
}

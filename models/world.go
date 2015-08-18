package models

import (
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/utils"
	"time"
)

type World struct {
	Locations map[string]*Location
	Name      string
	Tickers   map[string]*time.Ticker
	time      time.Time
	ticks     int64
}

func NewWorld() *World {
	world := &World{Name: NewName("World", "World").String()}
	world.Locations = make(map[string]*Location)
	world.Tickers = make(map[string]*time.Ticker)

	err := utils.LoadSetting(&world.Locations, "data/locations.json")
	if err != nil {
		panic(err.Error())
	}

	for _, loc := range world.Locations {
		loc.Creatures = make(map[string]*Creature)
		world.Tickers[loc.Key] = time.NewTicker(time.Millisecond)
		go world.Encapsulate(loc.Key, world.Tickers[loc.Key])
	}

	return world
}

func (this *World) Tick(t time.Time) {
	this.ticks += 1
}

func (this *World) MoveCharacter(character *Creature, locId string) {
	glog.V(10).Infof("Moving character to %s", locId)
	loc := this.Locations[locId]

	if loc.Creatures == nil {
		loc.Creatures = make(map[string]*Creature)
		this.Locations[locId] = loc
	}
	loc.Creatures[(*character).CreatureUid()] = character
}

func (this *World) RemoveCharacter(ownerId string) {
	for _, loc := range this.Locations {
		if _, exists := loc.Creatures[ownerId]; exists {
			glog.Warningf("Removing character from world: %s", ownerId)
			delete(loc.Creatures, ownerId)
		}

	}
}

// Make this static data
func (this *World) LocationsInfo() (locInfo map[string]LocationInfo) {
	locInfo = make(map[string]LocationInfo)
	for id, location := range this.Locations {
		locInfo[id] = LocationInfo{location.Name, location.Type, uint(len(location.Creatures))}
	}
	return
}

func SpellCasts(t time.Time, creatures map[string]*Creature) {
	for _, character := range creatures {
		if ability := (*character).NextAbility(); ability != nil {
			if ability.ActivationTime() == EFFECT_APPLIES_AT_START {
				ability.Target().ApplyEffect(ability)
			} else {
				if t.UnixNano() <= int64(ability.Start()+uint64(ability.Duration().Nanoseconds())) {
					ability.Target().ApplyEffect(ability)
				}
			}
		}
	}
}

func (this *World) Encapsulate(locId string, timer *time.Ticker) {
	loc := this.Locations[locId]
	for {
		t := <-timer.C
		if charNumbers := len(loc.Creatures); charNumbers > 0 {
			if this.ticks%500 == 0 {
				glog.V(5).Infof("Tick#%d: there are %d characters online at %s", this.ticks, charNumbers, loc.Name)
			}
			SpellCasts(t, loc.Creatures)
		} else if this.ticks%500 == 0 {
			glog.V(5).Infof("Tick#%d: there are no characters online, %s", this.ticks, loc.Name)
		}
	}
}

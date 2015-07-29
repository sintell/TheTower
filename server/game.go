package server

import (
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/models"
	"github.com/sintell/mmo-server/utils"
)

type Game struct {
	Characters map[string]*models.Character
	Monsters   map[string]*models.Monster
	Locations  map[string]*models.Location
	Items      map[string]*models.Item
}

var game *Game

func init() {
	game = &Game{}
	game.Characters = make(map[string]*models.Character)
	game.Monsters = make(map[string]*models.Monster)
	game.Locations = make(map[string]*models.Location)
	game.Items = make(map[string]*models.Item)

	game.LoadAssets()
	glog.Info("Creating game instance")
}

func GameInstance() *Game {
	if game == nil {
		game = &Game{}
		game.LoadAssets()
	}

	return game
}

func (this *Game) LoadAssets() {
	err := utils.LoadSetting(&this.Monsters, "data/monsters.json")
	if err != nil {
		panic(err.Error())
	}
	err = utils.LoadSetting(&this.Locations, "data/locations.json")
	if err != nil {
		panic(err.Error())
	}
	err = utils.LoadSetting(&this.Items, "data/items.json")
	if err != nil {
		panic(err.Error())
	}
}

func (this *Game) LoginCharacter(uid string, character *models.Character) error {
	if _, exists := this.Characters[uid]; exists {
		glog.Warningf("There are already character in game for this player: %s", uid)
		this.Characters[uid] = character
	}

	this.Characters[uid] = character
	return nil
}

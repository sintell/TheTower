package server

import (
	"errors"
	"fmt"
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

func (g *Game) LoadAssets() {
	err := utils.LoadSetting(&g.Monsters, "data/monsters.json")
	if err != nil {
		panic(err.Error())
	}
	err = utils.LoadSetting(&g.Locations, "data/locations.json")
	if err != nil {
		panic(err.Error())
	}
	err = utils.LoadSetting(&g.Items, "data/items.json")
	if err != nil {
		panic(err.Error())
	}
}

func (g *Game) LoginCharacter(uid string, character *models.Character) error {
	if _, exists := g.Characters[uid]; exists {
		glog.Errorf("Can't login character with uid %sCharacter already exists", uid)
		return errors.New(fmt.Sprintf("Can't login character with uid %sCharacter already exists", uid))
	}
	g.Characters[uid] = character
	return nil
}

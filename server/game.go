package server

import (
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/models"
	"github.com/sintell/mmo-server/utils"
	"time"
)

type Game struct {
	Characters   map[string]*models.Character
	Monsters     map[string]*models.Monster
	Locations    map[string]*models.Location
	Items        map[string]*models.Item
	World        *models.World
	loginChannel chan *models.Character
	stopChannel  chan interface{}
}

var game *Game

func GameInit() {
	game = &Game{}
	game.Characters = make(map[string]*models.Character)
	game.Monsters = make(map[string]*models.Monster)
	game.Locations = make(map[string]*models.Location)
	game.Items = make(map[string]*models.Item)
	game.World = models.NewWorld()

	game.LoadAssets()
	glog.Info("Creating game instance")
	game.loginChannel = game.Loop()
}

func init() {
	GameInit()
}

func GameInstance() *Game {
	if game == nil {
		GameInit()
	}

	return game
}

func (this *Game) LoadAssets() {
	err := utils.LoadSetting(&this.Monsters, "data/monsters.json")
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
	}
	glog.Infof("%s,\n\n%i", uid, character)

	this.loginChannel <- character
	this.Characters[uid] = character
	return nil
}

func (this *Game) LogoutCharacter(uid string) {
	if _, exists := this.Characters[uid]; exists {
		glog.Infof("Logging out character: %s", uid)

		// this.World.RemoveCharacter(character.UserID)
		delete(this.Characters, uid)
	}
}

func (this *Game) Loop() chan *models.Character {
	ticker := time.NewTicker(time.Millisecond)
	logging := make(chan *models.Character)
	this.stopChannel = make(chan interface{})

	go func(this *Game) {
		for {
			timer := <-ticker.C
			this.World.Tick(timer)

			select {
			case character := <-logging:
				{
					this.World.MoveCharacter(character, character.CurrentLocation)
				}
			default:
				{
					continue
				}
			}
			select {
			case <-this.stopChannel:
				{
					glog.Infof("Stopping game loop for world: %s", this.World.Name)
					ticker.Stop()
					break
				}
			default:
				{
					continue
				}
			}
		}
	}(this)

	return logging
}

func (this *Game) Stop() {
	close(this.stopChannel)
}

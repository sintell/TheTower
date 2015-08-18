package server

import (
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/models"
	"github.com/sintell/mmo-server/utils"
)

type Game struct {
	Creatures    map[string]models.Creature
	Characters   map[string]*models.Character
	Locations    map[string]*models.Location
	Items        map[string]*models.Item
	World        *models.World
	loginChannel chan models.Creature
	stopChannel  chan struct{}
}

var game *Game

func GameInit() {
	game = &Game{}
	game.Creatures = make(map[string]models.Creature)
	game.Locations = make(map[string]*models.Location)
	game.Items = make(map[string]*models.Item)
	game.World = models.NewWorld()

	game.loginChannel = game.Loop()
	game.stopChannel = make(chan struct{})

	game.LoadAssets()
	glog.Info("Creating game instance")
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
	err := utils.LoadSetting(&this.Items, "data/items.json")
	if err != nil {
		panic(err.Error())
	}

	npcs := models.GetNpcs()
	for _, npc := range npcs {
		this.LoginCharacter(npc)
	}
}

func (this *Game) LoginCharacter(character models.Creature) error {
	uid := character.CreatureUid()
	if _, exists := this.Creatures[uid]; exists {
		glog.Warningf("There are already character in game for this player: %s", uid)
	}

	this.loginChannel <- character

	switch character.Type() {
	case models.PLAYER:
		this.Characters[uid] = character.(*models.Character)
	case models.NPC:
		this.Creatures[uid] = character
	}

	return nil
}

func (this *Game) LogoutCharacter(uid string) {
	if _, exists := this.Creatures[uid]; exists {
		glog.Infof("Logging out character: %s", uid)

		this.World.RemoveCharacter(uid)
		delete(this.Creatures, uid)
	}
}

func (this *Game) Loop() chan models.Creature {
	loginChannel := make(chan models.Creature, 100000)
	go func(this *Game) {
		for {
			select {
			case character := <-loginChannel:
				{
					this.World.MoveCharacter(&character, character.CreatureLocation())
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
					break
				}
			default:
				{
					continue
				}
			}
		}
	}(this)

	return loginChannel
}

func (this *Game) Stop() {
	close(this.stopChannel)
}

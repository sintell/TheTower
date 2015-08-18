package models

import (
	"fmt"
	"github.com/sintell/mmo-server/utils"
	"strings"
)

type Npc struct {
	Character
	Behavior *Behavior
}

type NpcInfo struct {
	Behavior  string
	Locations []map[string]int
	Class     string
}

func GetNpcs() (npcs []*Npc) {
	var data []NpcInfo
	err := utils.LoadSetting(&data, "data/npc.json")
	if err != nil {
		panic(err.Error())
	}

	for _, npcType := range data {
		for _, location := range npcType.Locations {
			for locName, count := range location {
				counter := 0
				npcClass := npcType.Class
				npcBehavior := npcType.Behavior
				for i := 0; i < count; i++ {
					counter++
					npc := &Npc{
						Character{
							Name:            fmt.Sprintf("%s#%d", strings.ToTitle(npcBehavior), counter),
							Class:           npcClass,
							CurrentLocation: locName,
							OwnerUid:        fmt.Sprintf("%s#%d", strings.ToTitle(npcBehavior), counter),
						},
						ai.Behavior(npcBehavior),
					}
					npc.SetDefaults()
					npcs = append(npcs, npc)

				}
			}

		}
	}

	return
}

func NewNpc() (npc *Npc) {
	npc = &Npc{Character{Name: "Teleporter #1"}, ai.Behavior("teleporter")}
	npc.SetDefaults()
	npc.CurrentLocation = "mirage_bay_travellers_house"
	return
}

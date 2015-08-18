package models

import (
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/utils"
)

var ai Ai

func init() {
	ai.behaviors = make(map[string]Behavior)
	utils.LoadSetting(&ai.behaviors, "data/behaviors.json")
	glog.Infof("%i", ai.behaviors)
}

type Ai struct {
	behaviors map[string]Behavior
}

type Behavior struct {
	name string
}

func (this *Ai) Behavior(name string) *Behavior {
	if behavior, exists := this.behaviors[name]; exists {
		return &behavior
	} else {
		return &Behavior{name: "default"}
	}
}

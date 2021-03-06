package models

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

type User struct {
	gorm.Model
	Name            string          `sql:"unique" json:"name"`
	Email           string          `sql:"unique" json:"email"`
	Age             uint8           `sql:"default:'18'" json:"age,omitempty"`
	Uid             string          `sql:"default:'null';unique" json:"uid"`
	Characters      []Character     `sql:"" json:"characters"`
	Connection      *websocket.Conn `sql:"-"	json:"-"`
	ActiveCharacter *Character      `sql:"-" json:"-"`
}

func NewUser() (*User, error) {
	user := User{}

	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d", time.Now().Nanosecond())))
	uid := fmt.Sprintf("%x", h.Sum(nil))
	glog.Infof("Setting uid value equal %s", string(uid))

	name := NewName(uid, "User").String()
	email := fmt.Sprintf("%s@%s.%s", name, "testmail", "com")

	user.Uid = uid
	user.Name = name
	user.Email = email

	db.Create(&user)
	if db.Error != nil {
		return nil, db.Error
	}

	return &user, nil
}

func (this *User) NewCharacter() error {
	name := NewName(this.Email)
	classes := []string{"MAGE", "WARRIOR", "PRIEST"}
	character := Character{
		Name:          name.String(),
		Class:         classes[rand.Int31n(int32(len(classes)-1))],
		OwnerUid:      this.Uid,
		characterType: PLAYER,
	}
	character.SetDefaults()

	this.Characters = append(this.Characters, character)

	if err := db.Save(this).Error; err != nil {
		glog.Errorf("Error on user save: %s", err.Error())
		return err
	}
	glog.Infof("Chreating new character for %s %s", this.Name, this.Uid, character)
	this.LoadCharacters()

	this.SetActiveCharacter(this.Characters[0].ID)

	return nil
}

func (this *User) Populate() error {
	glog.Infof("Checking for user existance")

	if db.Where(this).Preload("Characters").First(this).RecordNotFound() {
		glog.Errorf("Requested user has uid but no matching record in db: %s", this.Uid)
		return gorm.RecordNotFound
	} else {
		glog.Infof("Requested user found: %s:%s, %s", this.Name, this.Email, this.Uid)
		return nil
	}
}

func (this *User) LoadCharacters() {
	if err := db.Model(this).Related(&this.Characters).Error; err != nil {
		glog.Errorf("Error requesting characters information: %s", err.Error())
	} else {
		glog.Infof("Characters loaded: %i", this.Characters)
	}
}

func (this *User) SetActiveCharacter(characterId uint) *Character {
	var charLinc *Character
	for _, character := range this.Characters {
		if character.ID == characterId {
			charLinc = &character
			this.ActiveCharacter = charLinc
		}
	}

	if len(this.Characters) == 0 {
		this.LoadCharacters()
		this.SetActiveCharacter(characterId)
	}

	return this.ActiveCharacter
}

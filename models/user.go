package models

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string          `sql:"unique" json:"name"`
	Email      string          `sql:"unique" json:"email"`
	Age        uint8           `sql:"default:'18'" json:"age,omitempty"`
	Uid        string          `sql:"default:'null';unique" json:"uid"`
	Characters []Character     `sql:"" json:"characters"`
	Connection *websocket.Conn `sql:"-"	json:"-"`
}

func NewUser() (*User, error) {
	db := GetDB()
	user := User{}

	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d%s", user.ID, user.Email)))
	uid := fmt.Sprintf("%x", h.Sum(nil))
	glog.Infof("Setting uid value equal %s", string(uid))

	user.Uid = uid

	db.Create(&user)
	if db.Error != nil {
		return nil, db.Error
	}

	return &user, nil
}

func (this *User) NewCharacter() error {
	name := NewName(this.Email)
	class := "MAGE"

	character := Character{Name: name.String(), Class: class}
	character.SetDefaults()

	this.Characters = append(this.Characters, character)

	glog.Infof("Chreating new character for %i", this)
	db.Save(this)

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

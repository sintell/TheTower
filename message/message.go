package message

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/sintell/mmo-server/models"
	"io"
)

type Message struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
	Uid  string          `json:"uid"`
}

type MessageType uint

// Request  constants
const (
	MSG_ERROR MessageType = 1 + iota
	MSG_GET_USER
	MSG_CREATE_CHARACTER
	MSG_REMOVE_CHARACTER
	MSG_CHECK_CHARACTER
	MSG_LOGIN_CHARACTER

	MSG_USER_ACTION

	MSG_UA_LOAD_GAME_DATA
)

// Response constants
const (
	MSG_USER_DATA = 1 + iota
	MSG_CHARACTER_DATA
)

var db gorm.DB

func init() {
	db = models.GetDB()
}

// TODO: Move func from this module
func GetUser(r io.Reader) (*models.User, error) {
	dec := json.NewDecoder(r)
	data := Message{}
	user := &models.User{}

	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}

	glog.Infof("Checking for user existance")

	if data.Uid != "" {
		if db.Where(&models.User{Uid: data.Uid}).Preload("Characters").First(user).RecordNotFound() {
			glog.Errorf("Requested user has uid but no matching record in db: %s", data.Uid)
		} else {
			glog.Infof("Requested user found: %s:%s, %s", user.Name, user.Email, user.Uid)
			return user, nil
		}
	}
	glog.Infof("Connection from new user")
	user, err = models.NewUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func New(mType MessageType, uid string, data []uint8) Message {
	glog.Infof("Creating new message with type %d for %s", mType, uid)
	return Message{mType, json.RawMessage(data), uid}
}

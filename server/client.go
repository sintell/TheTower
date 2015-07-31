package server

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"github.com/sintell/mmo-server/message"
	"github.com/sintell/mmo-server/models"
	"io"
	"time"
)

type Client struct {
	Conn        *websocket.Conn
	ConnectedAt int64
	User        *models.User
}

func NewClient(conn *websocket.Conn, user *models.User) *Client {
	return &Client{conn, time.Now().Unix(), user}
}

func (this *Client) Send(data interface{}, dataType message.MessageType) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := message.New(dataType, this.User.Uid, byteData)
	this.Conn.WriteJSON(&msg)

	return nil
}

func (this *Client) Handle() {
	msg := &message.Message{}
	glog.Infof("Starting message loop for: %s", this.User.Uid)

	for {
		_, reader, err := this.Conn.NextReader()
		dec := json.NewDecoder(reader)

		if err != nil {
			if err != io.EOF {
				glog.Errorf("Error creating next reader: %s", err.Error())
			}

			return
		}

		err = dec.Decode(msg)
		if err != nil {
			glog.Errorf("Error decoding message: %s", err.Error())
			return
		}
		glog.Infof("Got message from %s with type %d", this.User.Uid, msg.Type)

		switch msg.Type {
		case message.MSG_ERROR:
			{

			}
		case message.MSG_GET_USER:
			{

			}
		case message.MSG_CREATE_CHARACTER:
			{
				glog.Info("Got new character request")
				err := this.User.NewCharacter()
				if err != nil {
					glog.Errorf("Error creating character: %s", err.Error())
					break
				}
				game.LoginCharacter(this.User.Uid, this.User.ActiveCharacter)
				this.Send(this.User.Characters, message.MSG_CHARACTER_DATA)

			}
		case message.MSG_REMOVE_CHARACTER:
			{

			}
		case message.MSG_CHECK_CHARACTER:
			{
				glog.Info("Requested characters")
				this.User.LoadCharacters()
				var data []map[string]interface{}

				for _, char := range this.User.Characters {
					data = append(data, char.ShortData())
				}

				this.Send(data, message.MSG_CHARACTER_DATA)
			}
		case message.MSG_USER_ACTION:
			{

			}
		}
	}
}

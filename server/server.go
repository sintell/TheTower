package server

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"github.com/sintell/mmo-server/message"
	"github.com/sintell/mmo-server/models"
	"github.com/sintell/mmo-server/utils"
	"net/http"
)

type Server struct {
	utils.Settings
	Users map[string]*Client
}

func init() {
	flag.Parse()
}

func Init() *Server {
	server := &Server{Users: make(map[string]*Client)}
	err := server.LoadArgs()

	if err != nil {
		panic(err.Error())
	}

	db := models.GetDB()
	if db.Error != nil {
		fmt.Errorf("%s", db.Error.Error())
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  server.WsReadBuffSize,
		WriteBufferSize: server.WsWriteBuffSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		defer conn.Close()

		if err != nil {
			glog.Errorf("Error creating ws handler: %s", err.Error())
			return
		}

		glog.Infof("Creating ws handler from %s", fmt.Sprintf("%s with %s", r.Referer(), r.UserAgent()))

		uid, err := server.NewConnection(conn)
		defer server.CloseConnection(uid)

		client := server.Users[uid]

		client.Handle()
	})

	glog.Infof("App started on %s\n", fmt.Sprintf("%s:%s", server.Ip, server.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", server.Ip, server.Port), nil)

	return server
}

func (this *Server) NewConnection(conn *websocket.Conn) (string, error) {
	_, reader, err := conn.NextReader()
	if err != nil {
		glog.Errorf("Error establishing new connection: %s", err.Error())
	}
	user, err := message.GetUser(reader)
	if err != nil {
		return "", err
	}

	glog.Infof("Register conenction with uid: %s", user.Uid)
	this.Users[user.Uid] = NewClient(conn, user)
	this.Users[user.Uid].Send(user, message.MSG_USER_DATA)

	return user.Uid, nil
}

func (this *Server) CloseConnection(uid string) {
	glog.Infof("Closing connection from: %s", uid)
	delete(this.Users, uid)
}

package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/sintell/mmo-server/server"
	"os"
	"os/signal"
)

// Depends on: gorm, gorilla-websockets
// Client -> Front -> Auth -> Back -> DB -> back -> Front -> Client

func init() {
	flag.Parse()
}

// start options:
// -logtostderr=true
// -log_dir="."
// -alsologtostderr=false

func main() {
	defer glog.Flush()
	glog.Info("Starting...")

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)

	go func() {
		<-osSignals
		glog.Info("Got SIGINT, terminating app")
		glog.Flush()
		os.Exit(0)
	}()
	server.Init()
}

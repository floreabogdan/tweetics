package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/floreabogdan/tweetics/core"
	"github.com/floreabogdan/tweetics/dashboard"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "gui")

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT)
	signal.Notify(stopChan, syscall.SIGTERM)

	core := core.New()
	go core.Run()

	dash := dashboard.New(core)
	go dash.Run()

	<-stopChan
	log.Info("Got stop signal. Finishing work.")
}

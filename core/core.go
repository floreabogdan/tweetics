package core

import "github.com/sirupsen/logrus"

var log = logrus.WithField("module", "core")

type Core struct{}

func New() *Core {
	return &Core{}
}

func (c *Core) Run() {

}

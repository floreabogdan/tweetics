package dashboard

import (
	"goweb/core"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "core")

type Dashboard struct {
	engine *gin.Engine

	core *core.Core
}

func New(core *core.Core) *Dashboard {
	return &Dashboard{
		core: core,
	}
}

func (d *Dashboard) Run() {
	d.engine = gin.Default()
	d.initRoutes()

	err := d.engine.Run()
	if err != nil {
		log.Fatal(err)
	}
}

package dashboard

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func (d *Dashboard) initRoutes() {
	d.engine.LoadHTMLGlob("static/views/**/**")
	d.engine.Use(static.Serve("/static/assets", static.LocalFile("static/assets", false)))

	d.engine.GET("/ping", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		c.HTML(http.StatusOK, "index", gin.H{})
	})
}

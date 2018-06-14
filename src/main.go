package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonnygoncalves/extraterrestres/src/db"
	"github.com/tonnygoncalves/extraterrestres/src/models"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "asdfasdf", nil)
	})
	router.GET("/iniprocess", handleProcess)
	router.GET("/clima", handleGetData)

	router.Run(":9990")

}

func handleProcess(c *gin.Context) {
	db.CrearTablas()
	var planet = models.GetPlanetInfo(1)
	fmt.Println(planet)
	//c.Done("Proceso funcionó correctamente")

}

func handleGetData(c *gin.Context) {
	dia := c.DefaultQuery("dia", "0")
	c.JSON(200, gin.H{
		"dia":   dia,
		"clima": "asdfasdf",
	})
}

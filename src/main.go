package main

import (
	"fmt"

	"github.com/tonnygoncalves/extraterrestres/src/db"
	"github.com/tonnygoncalves/extraterrestres/src/models"
)

func main() {
	db.CrearTablas()
	var planet = models.GetPlanetInfo(1)
	fmt.Println(planet)
}

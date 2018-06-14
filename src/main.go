package main

import (
	"fmt"

	"github.com/tonnygoncalves/extraterrestres/src/models"
)

func main() {
	var planet = models.GetPlanetInfo(1)
	fmt.Println(planet)
}

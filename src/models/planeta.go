package models

import "math"

//This planet model is to affect all planets from the solar system.

type Planet struct {
	Name            string
	AngularVelocity int
	Distancy        int
	Grado           int
	X               float64
	Y               float64
}

func GetPlanetInfo(planetNumber int) *Planet {
	p := new(Planet)
	switch planetNumber {
	case 1:
		p = &Planet{
			Name:            "Ferengi",
			AngularVelocity: 1,
			Distancy:        500,
			Grado:           0,
			X:               500,
			Y:               0,
		}
	case 2:
		p = &Planet{
			Name:            "Betasoide",
			AngularVelocity: 3,
			Distancy:        2000,
			Grado:           0,
			X:               2000,
			Y:               0,
		}
	case 3:
		p = &Planet{
			Name:            "Vulcano",
			AngularVelocity: -5,
			Distancy:        1000,
			Grado:           0,
			X:               1000,
			Y:               0,
		}
	}

	return p
}

// Calculate new x
func Calculate_X(d float64, v float64, i int) float64 {
	return d * math.Cos(((-v*math.Pi)/180)*float64(i))
}

// Calculate new y
func Calculate_Y_y(d float64, v float64, i int) float64 {
	return d * math.Sin(((-v*math.Pi)/180)*float64(i))
}

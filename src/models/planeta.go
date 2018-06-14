package models

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

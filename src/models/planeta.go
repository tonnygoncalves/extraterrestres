package models

//This planet model is to affect all planets from the solar system.

type Planet struct {
	Name            string
	AngularVelocity int
	Distancy        int
	x               float64
	y               float64
}

func GetPlanetInfo(planetNumber int) *Planet {
	p := new(Planet)
	switch planetNumber {
	case 1:
		p = &Planet{
			Name:            "Ferengi",
			AngularVelocity: 3,
			Distancy:        500,
		}
	case 2:
		p = &Planet{
			Name:            "Frangi",
			AngularVelocity: 3,
			Distancy:        500,
		}
	case 3:
		p = &Planet{
			Name:            "Frangi",
			AngularVelocity: 3,
			Distancy:        500,
		}
	}

	return p
}

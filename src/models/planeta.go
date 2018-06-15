package models

import (
	"fmt"
	"math"

	"github.com/tonnygoncalves/extraterrestres/src/db"
	"github.com/tonnygoncalves/extraterrestres/src/utils"
)

//Planet model is to affect all planets from the solar system.
type Planet struct {
	Name            string
	AngularVelocity int
	Distancy        int
	Grado           int
	X               float64
	Y               float64
}

//Struct for X and Y
type place struct {
	x, y float64
}

var (
	vTime          int //10 years
	maxperimetro   float64
	diamaximo      int
	contadorSequia int64 // contador para saber cuantos días de sequía hay en el periodo a evaluar
	contadorLluvia int64 // contador para saber cuantos días de lluvia hay en el periodo a evaluar
)

//GetPlanetInfo get info with a number planet
func GetPlanetInfo(planetNumber int) *Planet {
	p := new(Planet)
	switch planetNumber {
	case 0:
		p = &Planet{
			Name:            "The Sun",
			AngularVelocity: 0,
			Distancy:        0,
			Grado:           0,
			X:               0,
			Y:               0,
		}
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

// GetNewCoordenates get new coordenates for a new grade
func GetNewCoordenates(planetNumber int, grade int) *Planet {
	var planet = GetPlanetInfo(planetNumber)
	planet.X = equations.CalculateX(float64(planet.Distancy), float64(planet.AngularVelocity), grade)
	planet.Y = equations.CalculateY(float64(planet.Distancy), float64(planet.AngularVelocity), grade)
	return planet
}

// ApointInATriangle detect if some point is or not inside a triangle
func ApointInATriangle(p *Planet, p1 *Planet, p2 *Planet, p3 *Planet) bool {
	var qty int
	if p.Y > math.Min(p1.Y, p2.Y) {
		if p.Y <= math.Max(p1.Y, p2.Y) {
			if p.X <= math.Max(p1.X, p2.X) {
				if p1.Y != p2.Y {
					var xinters = (p.Y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y) + p1.X
					if p1.X == p2.X || p.X <= xinters {
						qty++
					}
				}
			}
		}
	}
	p1 = p2
	p2 = p3
	if p.Y > math.Min(p1.Y, p2.Y) {
		if p.Y <= math.Max(p1.Y, p2.Y) {
			if p.X <= math.Max(p1.X, p2.X) {
				if p1.Y != p2.Y {
					var xinters = (p.Y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y) + p1.X
					if p1.X == p2.X || p.X <= xinters {
						qty++
					}
				}
			}
		}
	}
	// if this is true point is outside triangle
	if qty%2 == 0 {
		return false // outside
	}
	return true // inside the triangle
}

func processPlanets() {
	vTime = 3600 //10 years
	db.CrearTablas()

	var i int
	//dias := []*dia{}
	for i = 0; i <= vTime; i++ {
		var tiempo = ""
		diamaximo = 0
		var p0 = GetNewCoordenates(0, 0)
		var p1 = GetNewCoordenates(1, i)
		var p2 = GetNewCoordenates(2, i)
		var p3 = GetNewCoordenates(3, i)
		s := ApointInATriangle(p0, p1, p2, p3)
		if s {
			tiempo = "Lluvia"
			contadorLluvia++
		}
		fmt.Printf("Lluvia: %s", tiempo)
		//Formula: P0P3 = P0P1 + P1P2 + P2P3
		//la formula de la recta dice, que los planetas estaran alineados cuando la suma le la distancia del planeta más cercano
		//en este caso sera el sol hasta el planeta más lejano, en este caso vulcano, sera igual a la suma de las distancias del sol
		//a planeta 1 (ferengi), planeta 1(ferengi) a planeta 2(vulcano), planeta 2(vulcano) a planeta 3(betasioide)
		/*var P0P3 = equations.DistanceBetween2points(px0, px2, py0, py2) // distancia desde el sol al planeta mas lejano Betasoide
		var P0P1 = equations.DistanceBetween2points(px0, px1, py0, py1) // distancia desde el sol al planeta mas cercano al sol
		var P1P2 = equations.DistanceBetween2points(px1, px3, py1, py3) // distancia del primer planeta al segundo planeta
		var P2P3 = equations.DistanceBetween2points(px3, px2, py3, py2) // distancia desde el segundo planeta al tercero más lejano
		var totalp = P0P1 + P1P2 + P2P3
		if totalp > maxperimetro {
			limpiarTodosLosMax()
			maxperimetro = totalp
			diamaximo = 1
		}*/ /*
			if totalp > P0P3-200 && totalp < P0P3+200 {
				tiempo = "Sequia"
				contadorSequia++
				//fmt.Printf("P0P3: %f total: %f P0P1: %f P1P2: %f P2P3: %f", P0P3, total, P0P1, P1P2, P2P3, )
			}
			fmt.Printf("Lluvia: %s Sequía: %s Day: %f contador: %f contador lluvia: %f PX1=%f PY1=%f PX2=%f PY2=%f PX3=%f PY3=%f \n", 0, 0, i, contadorSequia, contadorLluvia, px1, py1, px2, py2, px3, py3)

			nuevoDia(i, tiempo, diamaximo)*/
	}

	/*fmt.Printf("Lluvia: %s Sequía: %s Maxp %f Dia: %f \n", contadorLluvia, contadorSequia, maxperimetro, diamaximo)*/
}

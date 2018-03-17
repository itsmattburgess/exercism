// Package space calculates how old someone would be on planets in our solar system
package space

const (
	earthSeconds = 31557600
)

type Planet string

// How long is 1 earth second on each planet?
var relativeSeconds = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates how old someone is on another planet
func Age(age float64, planet Planet) float64 {
	return age / (earthSeconds * relativeSeconds[planet])
}

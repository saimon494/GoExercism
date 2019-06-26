package space

type Planet = string

func Age(seconds float64, planet Planet) (age float64) {
	const earthOrbitalPeriod = 31557600

	switch planet {
	case "Earth":
		age = seconds / earthOrbitalPeriod
	case "Mercury":
		age = seconds / (0.2408467 * earthOrbitalPeriod)
	case "Venus":
		age = seconds / (0.61519726 * earthOrbitalPeriod)
	case "Mars":
		age = seconds / (1.8808158 * earthOrbitalPeriod)
	case "Jupiter":
		age = seconds / (11.862615 * earthOrbitalPeriod)
	case "Saturn":
		age = seconds / (29.447498 * earthOrbitalPeriod)
	case "Uranus":
		age = seconds / (84.016846 * earthOrbitalPeriod)
	case "Neptune":
		age = seconds / (164.79132 * earthOrbitalPeriod)
	}

	return
}

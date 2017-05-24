package simulation

import "math"

type SimulationSettings struct {
	//The amount of simluation seconds between each update (in minutes)
	StepTime int

}

type Planet struct {
	Radius float64
	Rocket Rocket
}

func (p Planet) AddRocket(settings RocketSettings) Rocket {

	return Rocket{}
}

func (p Planet) HasCrashed(rocket Rocket) bool {
	// Rocket radius is 1, for simplicity's sake
	if math.Pow(rocket.Position.X-0, 2) + math.Pow(rocket.Position.Y-0, 2) < (p.Radius+1) {
		return true
	}
	return false
}

type Rocket struct {
	Position Vector
	Velocity Vector
	// IN KILOGRAMS
	FullMass float64
	EmptyMass float64
	RocketParts []RocketPart
}

func (r Rocket) GetCenterOfGravity() Vector {

	return Vector{0,0}
}

type RocketPart struct {

}

type RocketSettings struct {

}
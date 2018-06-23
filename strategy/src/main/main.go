package main

import (
	"travel"
)

func main() {
	var travelone travel.Vehicle
	travelone.Init("shanghai", "hongkong", new(travel.Plane))
	travelone.Move()

	travelone.Init("shanghai", "hongkong", new(travel.Car))
	travelone.Move()
}

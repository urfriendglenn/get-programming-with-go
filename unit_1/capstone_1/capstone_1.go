package main

import (
	"fmt"
	"math/rand"
)

var airlines = [3]string{"Virgin Galactic", "SpaceX", "Space Adventures"}
var tripTypes = [2]string{"One-way", "Round-trip"}
var airline = airlines[rand.Intn(3)]
var tripType = tripTypes[rand.Intn(2)]
var speed = rand.Intn(15) + 16
var days = (62100000 / speed) / 86400

func main() {
	//TODO: Create price logic, apply formatting, and loop 10 times
	fmt.Println(airline)
	fmt.Println(tripType)
	fmt.Println(days)
}

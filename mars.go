package main

import "fmt"

// func main() {

// 	var hours = 28 * 24     //hours
// 	var distance = 56000000 // km

// 	// var speed = hours / distance

// 	fmt.Println(hours/distance, " km/h is how fast a ship would need to travel to Malacandra")
// 	// fmt.Println(speed)
// }

func main() {
	const hoursPerDay = 24

	var days = 28
	var distance = 56000000

	var speed = hoursPerDay * days / distance
	fmt.Println(speed, " km/h is how fast a ship would need to travel to Malacandra")
}

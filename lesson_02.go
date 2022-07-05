package two

import "fmt"

func main() {
	const hoursPerDay = 24

	var days = 28
	var distance = 56000000

	var speed = (hoursPerDay * days) / distance

	// fmt.Println(hours/distance, " km/h is how fast a ship would need to travel to Malacandra")
	// Why doesn't the code above work?
	fmt.Println(speed, " km/h is how fast a ship would need to travel to Malacandra")
}

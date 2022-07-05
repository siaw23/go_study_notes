package three

import "fmt"

func main() {
	var man = "tired"

	switch {
	case man == "thirsty":
		fmt.Println("Go get something to drink.")
	case man == "tired":
		fmt.Println("Get to bed.")
		fallthrough
	case man == "hungry":
		fmt.Println("Go to the kitchen!")
	}

	// This prints nothing. I expect it to print:
	// Go to bed.
	// Go to the kitchen!
}

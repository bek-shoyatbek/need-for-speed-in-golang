package src

import "fmt"

func Restart() {
	var input string
	fmt.Println("Enter 1 to continue \t 0 to cancel")
	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}
	fmt.Println(input)
}

func conclude(inp string) {
	if inp == "1" {

	}
}

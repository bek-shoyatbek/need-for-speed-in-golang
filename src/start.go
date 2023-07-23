package src

import (
	"fmt"
)

func Start(s string) {
	Clean()
	line1 := makeWall(100)
	line2 := makeWall(43)
	line2 += " " + s + " "
	line2 += makeWall(43)
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line1)
}

func makeWall(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "#"
	}
	return str
}

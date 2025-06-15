package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter the city:")

	//to read newline/whitespace for New York usecase
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	switch city {
	case "New York":
		fmt.Println("city is", city)

	case "Singapore", "London":
		fmt.Println("Same Kind Of city:", city)

	default:
		fmt.Println("Unknown city:", city)
	}

}

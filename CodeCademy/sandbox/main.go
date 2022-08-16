package main

import "fmt"

func main() {
	queryDatabase("SELECT * FROM coolTable;")
	var numOfFlavors int
	var b = 32
	numOfFlavors = 57
	fmt.Println(numOfFlavors, b)
	// type of b
	fmt.Printf("%T\n", b)

	// Scoped Short Declaration Statement
	x := 8
	y := 9
	if product := x * y; product > 60 {
		fmt.Println(product, "is greater than 60")
	}

	switch season := "summer"; season {
	case "summer":
		fmt.Println("Go out and enjoy the sun!")
	}
}

func queryDatabase(query string) string {
	var result string
	connectDatabase()
	// Add deferred call to disconnectDatabase() here
	defer disconnectDatabase()

	if query == "SELECT * FROM coolTable;" {
		result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
	}
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}

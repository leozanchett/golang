package main

import "fmt"

func main() {
	// pointerForInt armazenará o endereço de uma variável que possui um tipo de dados int. Para detalhar ainda mais,
	//o operador * significa que esta variável armazenará um endereço e a parte int significa que o endereço contém um valor inteiro
	var pointerForInt *int
	minutes := 525600
	pointerForInt = &minutes
	fmt.Println(pointerForInt) // Prints 0xc000018038

	lyrics := "Moments so dear"
	pointerForStr := &lyrics
	*pointerForStr = "Journeys to plan"
	fmt.Println(lyrics)

	star := "Polaris"
	lua := "lua"
	starAddress := &star
	starAddress = &lua
	// Add your code below:
	*starAddress = "Sirius"
	fmt.Println("The actual value of star is", star, "and the value of lua is", lua)
}

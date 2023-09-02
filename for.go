package main

import "fmt"

func main() {

	nomor := 1

	for nomor <= 10 {
		fmt.Println(nomor)
		nomor++
	}

	//Cara ke-2
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//Slice for
	animal := []string{
		"ular",
		"komodo",
		"buaya",
	}

	for i := 0; i < len(animal); i++ {
		fmt.Println("slice for:", animal[i])
	}
}

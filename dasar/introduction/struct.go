package main
import "fmt"

type Customer struct {
	Name, Address string
	Age			int
}


func main(){
	// Cara 1
	var akbar Customer
	akbar.Name = "Akbar"
	akbar.Address = "Surabaya"
	akbar.Age = 14
	fmt.Println(akbar)

	// Cara 2
	lucu := Customer{
		Name: "Akbar lucu",
		Address: "Rahasia",
		Age: 80,
	}
	fmt.Println(lucu)

	// Cara 3
	akbarlucu := Customer{"Ganteng", "Kepo", 15}
	fmt.Println(akbarlucu)

	asda := Customer{
		
	}
	
}
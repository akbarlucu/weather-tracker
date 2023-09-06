// Tipe data slice memiliki 3 data yaitu pointer, length, dan capacity
// Pointer adalah petunjuk data pertama di array pada slice
// Length adalah panjang dari slice
// Capacity adalah kapasitas dari slice, length tidak boleh lebih dari capacity


package main

import "fmt"

func main(){
	var month = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// month [5] = "Diganti"
	// fmt.Println(slice1)

	// slice1[0] = "Mei juga"
	// fmt.Println(slice1)

	var slice2 = month [10:]
	fmt.Println(slice2)

	//---APPEND---
	// Add slice
	var slice3 = append(slice2, "Desemer lagi")
	fmt.Println(slice3)

	// Change slice
	slice3[1] = "bukan desember"
	fmt.Println(slice3)
}
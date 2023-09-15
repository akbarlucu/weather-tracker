package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address {"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := &address1 // Tambah & agar addres1 mengikuti address2

	address2.City = "Bandung"

	*address2 = Address {"Malang", "Jawa Barat", "Indonesia"} // Operator * untuk mengganti semua yg ada di memory address1

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)

	var address4 *Address = new(Address)
	fmt.Println(address4)
}
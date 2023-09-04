package main
import "fmt"

//Constant seperti variable tapi tidak bisa diubah datanya setelah dideklarasi

func main(){
	//constant tidak akan error meskipun tidak dicompile
	const firstName string = "rachmad"
	const lastName = "Rizqi"

/*	code dibawah akan error karena constant tidak bisa dirubah
	firstName = "akbar"
	lastName = "syah"
*/
	fmt.Println(firstName)
	fmt.Println(lastName)

	//Deklarasi multi constant
	const (
		namaDepan string = "rachmad"
		namaBelakang = "Rizqi"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
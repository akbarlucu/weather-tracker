// Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
// Saat membuat array, perlu menentukan jumlah data yang bisa ditampung oleh array tersebut
// Daya tampung array tidak bisa bertambah setelah array dibuat

package main
import "fmt"

func main(){
	// Cara satu membuat array
	var names [3]string
	names [0] = "Rachmad"
	names [1] = "Rizqi"
	names [2] = "Akbar"
	
	// Print output secara terpisah
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//Cara dua membuat array secara langsung/ simple
	var values = [3]int{
		90,
		80,
		75,
	}
	// Print output dalam kotak array
	fmt.Println(values)

	// Print untuk mengetahui jumlah panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
}

package main
import "fmt"


type Man struct {
	Name string
}

// Tambahkan * ke pointer agar tidak usah merubah func main agar bisa print Mr
func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
	//fmt.Println("Di methode", man.Name)
}

func main(){
	akbar := Man{"Akbar"}
	akbar.Married()
	fmt.Println(akbar.Name)

}
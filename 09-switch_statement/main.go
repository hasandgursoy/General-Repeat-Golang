package main

import (
	"fmt"
)

func main() {

	grade := 3

	switch grade {
	case 5:
		fmt.Println("Pekiyi")
	case 4:
		fmt.Println("İyi")
	case 3:
		fmt.Println("Orta")

	default:
		fmt.Println("Geçersiz not")
	
	}
	

}

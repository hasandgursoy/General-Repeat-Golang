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

	

	switch trade := 5; {
	case trade < 10:
		fmt.Println("Pekiyi")
		fallthrough
	case trade < 15:
		fmt.Println("İyi")
		fallthrough

	case trade < 20:
		fmt.Println("Orta")

	default:
		fmt.Println("Geçersiz not")
	
	}
	

}

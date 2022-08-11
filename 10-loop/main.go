package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		fmt.Println(i)
	}

	var numb int = 5
	var result int = 1

	// Faktoriyel hesaplama

	for i := 1; i <= numb; i++ {
		result *= i
	}

	fmt.Println(result)

	// Infinite Loop

	// for {
	// 	fmt.Println("Benim Adım Hasan")
	// }
	// for i := 1; true; i++ {
	// 	fmt.Println(i)
	// }

	for i := 1; i < 10; i++ {

		fmt.Println(i)

		if i%3 == 0 {
			continue
		}

		if i == 7 {
			break
		}

	}

	// 1-100 arasında ki asal sayıları bulalım.

	for i := 3; i <= 100; i++ {

		var bol bool = false
		for x := 2; x < i; x++ {

			if i%x == 0 {
				bol = true
			}

		}

		if !bol {
			fmt.Println(i, "sayısı asaldır.")
		}

	}

}

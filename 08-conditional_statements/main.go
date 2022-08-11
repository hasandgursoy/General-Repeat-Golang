package main

import (
	"fmt"
)

func main() {

	x := 27

	if x%2 == 0 {
		fmt.Println(x, " : çift sayıdır.")
	} else {
		fmt.Println(x, " : tek sayıdır.")
	}

	if !false {
		fmt.Println("Mesaj Gösterilecek")
	}

	if x > 10 {
		x = 25
		fmt.Println(x)
	}else if x == 10 {

		x = 12
		fmt.Println(x)

	}else{
		fmt.Println(x)
	}

	// İlginç bir örnek yapalım
	// if statement'ın koşulunda değişken tanımladık :DASDASDASD
	
	if s := 10; s <= 10 {
		fmt.Println(s,"10 dan küçüktür veya eşittir.")
	}


}

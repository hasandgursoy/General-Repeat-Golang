package main

import (
	"fmt"
)

func main()  {
	
	x := 10.0
	y:= 10


	fmt.Printf("%v %T \n", x,x)
	fmt.Printf("%v %T \n", y,y)

	fmt.Println(x + float64(y))
	fmt.Println(int(x)+y)

	// Int bir değeri string bir ifadeye çevirmek aşşağıdaki gibi çok kolay :D 
	var name string = fmt.Sprint(y)
	fmt.Printf("%v %T \n",name,name)

	// ASCI karakterlerini string'e değiştirme

	str1 := string(106)
	fmt.Println(str1) // j harfine karşılık gelecek.


}
package main

import (
	"fmt"
	"math"
)

func main() {

	const pi float64 = 3.14
	r := 3.0

	areaOfCircle := 3.14 * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)

	// Sabitlere tabikide başka değer atayamıyoruz
	// Constantlar compile zamanına ait yapılardır.

	const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T %v \n", x, x)
	fmt.Printf("%T %v \n", y, y)
	fmt.Printf("%T %v \n", z, z)
	fmt.Printf("%T %v \n", t, t)

	// const -----> compile time
	// var ------> run time

	// Fonksiyonlar run time da çalışır bu yüzden aşşağıdaki yapıyı constant olarak atayamayız.
	// Çünkü constant ifadeler compile time da haberdar olmak ister.
	//const numb = math.Pow(3,4)

	// Aynı şekilde şu şekilde bir örnek verelim
	// sss değişkeni run time değişkeni o yüzden compile time da ki bir yapıya atayamıyorsun.
	// sss := 15
	// const r =sss

	// fmt.Println(r)

}

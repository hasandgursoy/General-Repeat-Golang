package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// Moduler Programing
	// Readable Code
	var value int = sum(5, 10)
	fmt.Println(value)
	var helllo string = helloAge("hasan", 23)
	fmt.Println(helllo)
	fmt.Println(result(50))
	getTime()
	getExamResult()
	bolum,kalan := divider(104,5)
	fmt.Println(bolum,kalan)
}

func sum(x, y int) int {
	return x + y
}

// String interpolation'a örnek.
func helloAge(name string, age int) string {
	hi := fmt.Sprintf("Adım %s, yaşım %d", name, age)
	return hi
}

func result(grade int) string {
	if grade >= 50 {
		return "Geçtiniz"
	} else {
		return "kaldınız"
	}
}
// Time işlemleri
func getTime() {
	var moment time.Time = time.Now()
	fmt.Println(moment.UTC())
}

func getExamResult() {
	fmt.Println("Lütfen sınav sonucunu giriniz : ")
	reader := bufio.NewReader(os.Stdin)
	// _ blank idenditifier  (hatayı şuan boşa sonra hallederiz diyoruz.)
	value, _ := reader.ReadString('\n')
	fmt.Println(value)
}


func divider(bolunen, bolen int) (bolum,kalan int){

	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum,kalan

}  
	

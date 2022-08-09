package main

import "fmt"

// Paket düzeyinde tanımlama yaparken büyük harfle başlarsak bu değişken diğer paketlerden ulaşılabilir hale gelir.
var OO string = "OO"

func main() {

	// Print işlemleri ile ilgili "fmt" paketine bakmalıyız.

	x := 100
	y := 20
	z := 30

	// Değişkenin değeri %v
	// Değişkenin type'ı %T
	fmt.Printf("%v %T \n", x, x)

	// İkilik tabanda gösterimi
	fmt.Printf("%b \n", y)

	// 16 lık tabanda gösterimi
	fmt.Printf("%X \n", z)

	// Bir çok yapı var bunların ihtiyaç halinde paketten okunması çok daha makul.

	//https://pkg.go.dev/fmt
	// yukarıdaki link'de fmt paketinin adresi var gerekli olan herşey yazıyor

	// Camel Case kullanılıyor GO da

	var coinType string = "Etherium"

	fmt.Println(coinType)

	// Kısaltmalar Büyük harfle yazılır

	var URL string = "www.google.com"
	var HTTP string = "localhost:8080"

	fmt.Println(URL, HTTP)

}

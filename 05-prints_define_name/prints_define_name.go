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

	// Statement = Yapılan en ufak işlem dahi bir statement 'dır ve genelde bir sonucu olur.
	// Expression = Bir sonucu olmayabilir ve birşey dönmek zorunda değildir.

	// Bu ifade bir statement'dır.
	fmt.Println(5)
	
	// Bu ifade bir expressiondır.
	//5*5

	// Golang'de statement durumunda bir bir expression tanımlarsak. 
	// örnek : 

	fmt.Println(5*5)

	// Bu ifade artık expression statements olur 
	// Ve Go da statement durumunda ki bir yapıda Inc yada Dec işlemleri yapılamaz.
	// örnek : 

	var val int = 10
	val-- // burda hata vermez ama statement durumunda verecektir.
	//fmt.Println(val--)



}

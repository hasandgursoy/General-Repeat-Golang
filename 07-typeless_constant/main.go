package main

import "fmt"

func main() {

	// Const'a veri tipi verdik :D
	const x int16 = 5
	const y = 5
	var z int32 = 23

	fmt.Printf("%T %v", x, x)
	fmt.Printf("%T %v", y, y)

	// Veri tipi olmayan sabitlerde kullanılmadığı sürece typless kalabilir.
	// Ancak kullanınca işlemin çalışması için tipi kendi atıyor.

	fmt.Println(y + z)
	// normalde bu satırda hata almamız lazım ama sorun olmuyor go typless olana type conversion uyguluyor.

}

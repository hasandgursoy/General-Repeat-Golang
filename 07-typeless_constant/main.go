package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Const'a veri tipi verdik :D
	const x int16 = 5
	const y = 5
	var z int32 = 23

	fmt.Printf("%T %v \n ", x, x)
	fmt.Printf("%T %v \n ", y, y)

	// Veri tipi olmayan sabitlerde kullanılmadığı sürece typless kalabilir.
	// Ancak kullanınca işlemin çalışması için tipi kendi atıyor.

	fmt.Println(y + z)
	// normalde bu satırda hata almamız lazım ama sorun olmuyor go typless olana type conversion uyguluyor.

	var name bool = (reflect.TypeOf("name") == reflect.TypeOf("hasan"))

	fmt.Println(name)
}

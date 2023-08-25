package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	

	fmt.Println(nilai32) 
	fmt.Println(nilai64)
	fmt.Println(nilai8) // get wrong value, because 100000 bigger than int8 (127)

	var name = "Oscar Hermawan"
	var e byte = name[6] // get byte code.
	fmt.Println(e)
	fmt.Println(string(e)) // string(e) convert to string

}
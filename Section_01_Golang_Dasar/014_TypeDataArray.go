package main

import "fmt"

func main() {
	var names [2]string  // we need to calculate how long the array is
	names[0] = "Oscar"
	names[1] = "Hermawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names)
}
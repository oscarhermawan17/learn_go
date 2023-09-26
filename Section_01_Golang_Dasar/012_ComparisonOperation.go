package main

import "fmt"

func main() {
	var name1 = "Oscar"
	var name2 = "Oscar"
	var name3 = "Hermawan"

	var result bool = name1 == name2
	var result2 bool = name1 > name2
	var result3 bool = name1 > name3
	fmt.Println(result, result2, result3)
}
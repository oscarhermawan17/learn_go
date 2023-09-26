package main

import "fmt"

/*
array[low:high] // create slice from "low" index until before "high" array
array[low:] // create slice from "low" index until last array
array[:high] // create slice from index 0 until before "high" array
array[:] // create slice from index 0 until last array

*/

func main() {
	var months = [...]string {
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length of slice1 = 3
	fmt.Println(cap(slice1)) // capacity of slice1 = 8

	// months[4] = "Diubah"   // be carefull when change the array
	// fmt.Println(slice1)

	var slice2 = months[11:] // this will use new array, but try this [2:4] and see the difference
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)
}
package main

import "fmt"

func main() {
	type NoKtp string // create new data type called NoKTP as a string
	type Married bool // create new data type called NoKTP as a boolean

	var noKtpOscar NoKtp = "123123123"
	var marriedStatus Married = true

	fmt.Println(noKtpOscar)
	fmt.Println(marriedStatus)
}
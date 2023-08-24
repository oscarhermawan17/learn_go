package main

import "fmt"

func main() {
	var name string = "Oscar Hermawan"
	fmt.Println(name)

	name = "Oscar Aja"  // we can reassign value, is not like const in javascript
	fmt.Println(name)

	var test = "Test tanpa type data" // go will add automatic data type if we dont write it, for this case string
	fmt.Println(test)

	var testBaru = 100 // automatic assign data type, for this case number, int64 (depends on OS)
	fmt.Println(testBaru)

	newData := "New String" // we can create new variable without var, but we use :=
	fmt.Println(newData)

	var (
		firstName = "Oscar"
		lastName = "Hermawan"
	) // create many variable in single declaration

	fmt.Println(firstName)
	fmt.Println(lastName)
}  
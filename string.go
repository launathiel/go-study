package main

import "fmt"

func main()  {
	var message string = "this is string"
	fmt.Println("the message is", message)
	name := "Laurentius Nathaniel"
	e:= name[2]
	fmt.Println(string(e))

	panjang := len("Nathaniel")
	fmt.Println("Nathaniel memiliki", panjang, "Huruf")
}
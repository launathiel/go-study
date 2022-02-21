package main

import "fmt"

func main(){
	type Married bool

	var isNathanMarried Married = false

	fmt.Println("Nathan udah menikah itu",isNathanMarried)

	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)
}
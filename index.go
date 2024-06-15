package main

import "fmt"

func main() {

	var a float64
	var b float64

	value, err := fmt.Scanln(&a, &b)

	fmt.Println(a + b)
	p,q:=fmt.Println(value,err)
	fmt.Println(p,q)

}

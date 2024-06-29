package main

import "fmt"

type shape interface{
	func area();
}

type struct rectange{
	Length int
	Breadth int
}

func (r rectange) area(length float64,breadth float64) float64{
	return r.Length,r.Breadth
}

func main(){
	fmt.Println("Hello babes")
}
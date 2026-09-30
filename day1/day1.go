package main

import (
	"fmt"
    // "math"
	"strconv"
)

func main(){
	// fmt.Println("Hello, Vedant Shekhar")

	// var x uint8 = 2
	// x = 9
	// fmt.Println(x)

	// y := uint(3)

	// fmt.Printf("%T",y)

	// fmt.Println("Vedant", 21,"Golang")

	// x:=8.3
	// fmt.Printf("%v %T %b",x,x,x)

	// x:=8.34309439
	// fmt.Printf("%e",x)

	// x:=8.34309439
	// fmt.Printf("%f",x)

	// x:=338.34309439
	// fmt.Printf("\"%10.2f%%",x)

	// x:=338.34309439
	// y := fmt.Sprintf("\"%10.2f%%",x)
	// fmt.Println(y,y,y)

	// x := "Hello"
	// y := 2
	// z := x + fmt.Sprint(y)

	// fmt.Println(z)

	// fmt.Println(math.Min(4,5))
	// fmt.Println(math.Max(4,5))
	// fmt.Println(math.Pow(4,11))

	x := "10101001"
	y, err := strconv.ParseInt(x,2,0)
	fmt.Println(y, err)
 

 


}
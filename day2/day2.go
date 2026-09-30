package main

import "fmt"

func main(){
	// a := 10;
	// for a>=1{
	// 	fmt.Println(a)
	// 	a--
	// }

	// str := "hello World"

	// fmt.Println(string(str[0]))
	// str := "Vedant Shekhar"

	// for idx := 0;idx<len(str);idx++{
	// 	fmt.Printf("%c",str[idx]);

	// }

	// var arr[5000] bool

	// arr := [2][2]int{{1,2},{3,4}}
	// fmt.Println(arr)

	arr := [...][2]int{{1,2},{3,4},{5,6}}
	test(arr);
	fmt.Println(arr)
	
	// for _, value := range arr{
	// 	for _,value2 := range value{
	// 		fmt.Println(value2);

	// 	}
	// }

	// arr1 := [...]int{1,2,3}
	// fmt.Println(arr1)
}

func test(arr[3][2]int){
	arr[0] = [2] int {100,100};
}

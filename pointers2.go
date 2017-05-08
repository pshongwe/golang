package main

import "fmt"

func swap(x *int, y *int){
	var temp int = *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	fmt.Println(x)
	fmt.Println(y)
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

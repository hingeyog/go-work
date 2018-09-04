package main 

import "fmt"

func main() {
	x := 15
	a := &x // 'address of' x

	fmt.Println("address of a (&a): ", a) 
	fmt.Println("contents at that address (*a): ", *a)

	*a = 5
	fmt.Println("*a = 5 ") 
	fmt.Println("x: ", x) 
}

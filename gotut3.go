package main

// https://www.youtube.com/watch?v=ZCceigT2A6I&index=3&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX
// https://pythonprogramming.net/go/types-go-language-programming-tutorial/?completed=/go/syntax-go-language-programming-tutorial/

import "fmt"

// func add (x,y float64)  -->> MEANS THE SAME THING AS BELOW
func add (x float64, y float64) float64 {
	return x+y
}

func multiple(a,b string) (string, string){
	return a,b
}

func main() {
	// var num1, num2 float64 = 5.6, 9.1 -->> MEANS THE SAME THING AS BELOW
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	// num1, num2 := 5.6, 9.1 -->> MEANS THE SAME THING. GO COMPILER WILL FIGURE OUT THE TYPES AND ASSIGN ACCORDINGLY.
	// Go compiler will give 'float64' by default. So if 'float32' is used in the add(), compiler will throw an error.

	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1,w2))



	var a int = 64
	var b float64 = float64(a)
	x:= a // x will be of type int
	fmt.Println("a is: ", a)
	fmt.Println("b is: ", b)
	fmt.Println("x is: ", x)


}		
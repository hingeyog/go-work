package main

// https://www.youtube.com/watch?v=ZCceigT2A6I&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=3

// all the sub packages need to be inported
// e.g. math.rand.Intn() --> NOT ALLOWED 

// math is separate package and rand is separate package, it is the way of organizing things.
// These packages were in hierarchical, so that it would be easy for some-one looking for something, categorizing.
// When you import math package rand package will not be imported.
// A good example would be encoding package where there were several encoding's, like xml, json, csv, gob,...
// They all fall into encoding category, since you need encoding of some type like json or gob, it doesn't mean you are going to use other encodings like xml, csv,...
// So if you individual package whichever suits your need.﻿
import ("fmt"
		"math"
		"math/rand")

func foo() {
	fmt.Println("The square root of 4 is: ", math.Sqrt(4))	
}

func main() {
	foo()
	fmt.Println("A numberfrom 1-100 : ", rand.Intn(100))

	//square bracket means include. Round bracket means to not include. 
	// Therefore [0,n) means 0 is a valid output and n is not a valid output﻿
	// [0,n] is a closed interval with 0 and n included in that interval
	// [0,n) is a half open interval with 0 included and n not included in the interval
	// (0,n) is a open interval where 0 and n is not included in that interval﻿
}		
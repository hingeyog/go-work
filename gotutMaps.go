package main

import "fmt"

// part 14 - https://www.youtube.com/watch?v=Uv4ZLaV6a9g

func main() {

	// var grades map[string]float32

	// or 
	//  grades:= map[string]float32 

	// 'map' is a refernce type. If you want to have values, use Go's make 
	// make() will initialize
	grades:= make(map[string]float32)

	grades["Timmy"] = 42 
	grades["Jess"] = 92 
	grades["Sam"] = 67

	fmt.Println(grades)

	TimsGrade := grades["Timmy"] 
	fmt.Println(TimsGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	for k,v := range grades {
		fmt.Println(k, ":", v)
	}

}

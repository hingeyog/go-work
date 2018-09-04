package main

import "fmt"
// https://www.youtube.com/watch?v=i3o4G4bmqPc&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=7

/*
	2 types of methods in Go:
		1. Value receivers - calculations on values
		2. pointer receivers - to change value in the struts
*/

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // min 0 max 65,535
	break_pedal uint16
	steering_wheel int16 // -32k to + 32k
	top_speed_kmh float64
}

// 'c car' is used to associate this function to the car 'struct'. 
// This makes it a method. 
// This is a 'Value receiver method'
func (c car) kmph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax) 
}

// This is a 'Value receiver method'
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax/kmh_multiple) 
}

func main() {
	// a_car:= car {22341, 0, 12561, 225.0} -->> SAME AS BELOW
	a_car:= car {gas_pedal: 65000, 
				break_pedal: 0, 
				steering_wheel: 12561, 
				top_speed_kmh: 225.0}

	fmt.Println("Curent pedal position: ", a_car.gas_pedal)
	fmt.Println("top_speed_kmh: ", a_car.kmph())
	fmt.Println("top_speed_mph: ", a_car.mph())
}
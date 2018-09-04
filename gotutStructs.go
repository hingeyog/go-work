package main

import "fmt"
// https://www.youtube.com/watch?v=3fsvqo9pQyg&index=6&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX

type car struct {
	gas_pedal uint16 // min 0 max 65,535
	break_pedal uint16
	steering_wheel int16 // -32k to + 32k
	top_speed_kmh float64
}

func main() {
	// a_car:= car {22341, 0, 12561, 225.0} -->> SAME AS BELOW
	a_car:= car {gas_pedal: 22341, 
				break_pedal: 0, 
				steering_wheel: 12561, 
				top_speed_kmh: 225.0}

	fmt.Println(a_car.gas_pedal)
}
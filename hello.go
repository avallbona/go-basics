package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, world")
	var x int = 5
	var y int = 7
	var sum2 int = x + y
	fmt.Println(sum2)

	x1 := 19
	y1 := 38
	sum1 := x1 + y1
	fmt.Println(sum1)

	var a [5]int
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	c := []int{6, 7, 8, 9}
	c = append(c, 10)
	c = append(c, 11)
	fmt.Println(c)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	fmt.Println(vertices)

	delete(vertices, "square")
	fmt.Println(vertices)

	fmt.Println("Looping")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Looping and array")
	for index, value := range c {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Println("Looping a map")
	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	fmt.Println("calling sum function")
	result := sum(4, 5)
	fmt.Println(result)

	fmt.Println("calling sqrt function")
	result2, err := sqrt(47)
	if err != nil {
		fmt.Println("an error ocurred: ", err)
	} else {
		fmt.Println(result2)
	}

	fmt.Println("Using struct")
	usingStruct()

	fmt.Println("Using pointers")
	usingPointers()

}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

type person struct {
	name string
	age  int
	birthdate birthdate
}

type birthdate struct{
	year int
	month int
	day int
}

func usingStruct() {

	p1 := person{name: "Andreu", age: 48, birthdate: birthdate{year:1973, month: 5, day: 1}}
	p2 := person{name: "Anne", age: 45, birthdate: birthdate{year: 1976, month: 10, day: 29}}
	fmt.Println("p1: ", p1)
	fmt.Println("p1.name: ", p1.name, "p1.age:", p1.age, "birthday", p1.birthdate)
	fmt.Println("p2: ", p2)
	fmt.Println("p2.name: ", p1.name, "p2.age:", p2.age, "birthday", p2.birthdate)
}

func usingPointers() {
	i := 7
	fmt.Println("initial value of i:", i)
	inc(i)
	fmt.Println("value of i:", i)
	fmt.Println("address of i:", &i)

	j := 10
	fmt.Println("initial value of j:", j)
	inc_by_reference(&j)
	fmt.Println("value of j:", j)

}

func inc(x int) {
	x++
}

func inc_by_reference(x *int) {
	*x++
}

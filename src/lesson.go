package main

import (
	"fmt"
	"time"
	)

func add(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y
}

func calc(price, item int) (result int){
	result = price * item
	return
}

func incrementGenerator() (func() int){
	x := 0
	return func() int{
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64{
	return func(radius float64) float64{
		return pi * radius * radius
	}
}

func foo(params ...int){
	fmt.Println(len(params), params)
	for _, param := range params{
		fmt.Println(param)
	}
}

func by2(num int) string{
	if num % 2 == 0{
		return "ok"
	}else{
		return "no"
	}
}

func getOsName() string{
	return "mac"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mac")
	default:
		fmt.Println("default")		
	}

	t := time.Now()
	fmt.Println(t.Hour())

	// l := []string{"python", "go", "java"}

	// for i := 0; i < len(l); i++{
	// 	fmt.Println(i, l[i])
	// }

	// for _, v := range l{
	// 	fmt.Println(v)
	// }

	// m := map[string] int{"apple": 100, "banana": 200}

	// for k, v := range m{
	// 	fmt.Println(k, v)
	// }

	// sum := 1
	// for ; sum < 10; {
	// 	sum += sum
	// 	fmt.Println(sum)
	// }
	// fmt.Println(sum)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// if result := by2(10); result == "ok"{
	// 	fmt.Println("great")
	// }

	// num := 5
	// if num % 2 == 0 {
	// 	fmt.Println("by 2")
	// } else {
	// 	fmt.Println("else")
	// }

	// x, y := 10, 10
	// if x== 10 && y == 10{
	// 	fmt.Println("&")
	// }
	// f := 1.11
	// fmt.Println(int(f))

	// // 5, 6

	// m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	// fmt.Printf("%T %v", m, m)
}


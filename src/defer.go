package main

import	(
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {

	foo()

	// 関数の最後に実行される
	defer fmt.Println("world")

	fmt.Println("hello")

	// deferは最後のものから実行される
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")


	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 10000)
	file.Read(data)
	fmt.Println(string(data))
}
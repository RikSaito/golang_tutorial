package main

import "fmt"

type Vertex struct{
	x, y int
}

// 値レシーバー
func (v Vertex) Area() int {
	return v.x * v.y	
}

// ポインタレシーバー
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func Area(v Vertex) int {
	return v.x * v.y
}

type Vertex3D struct{
	Vertex
	z int
}

// 値レシーバー
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z	
}

// ポインタレシーバー
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// package名.New　でstructを返す（デザインパターン）
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}


// package名.New　でstructを返す（デザインパターン）
// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

func main() {
	v := New(3, 4, 5)
	// fmt.Println(Area(v))
	// v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}
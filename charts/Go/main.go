package main

import "fmt"

type NhanVien struct {
	Id   int
	Name string
	Age  int
}

func main() {
	employees := [...]NhanVien{
		{Id: 1, Name: "nguyen van a", Age: 18},
		{Id: 2, Name: "nguyen van b", Age: 19},
		{Id: 3, Name: "nguyen van c", Age: 20},
	}
	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i])
	}
}

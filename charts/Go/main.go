package main

import "fmt"

type giangvien struct {
	name   string
	gender int
	phone  string
}

func (g *giangvien) hienthithongtin() {
	fmt.Printf("Ho ten giang vien: %s \n", g.name)
	fmt.Printf("Gioi tinh: %d \n", g.gender)
	fmt.Printf("So dien thoai: %s \n", g.phone)
}

func (g *giangvien) clear() {
	g.name = ""
	g.gender = 0
	g.phone = ""
}

func main() {
	gv := giangvien{
		name:   "John",
		gender: 1,
		phone:  "123456789",
	}
	gv.hienthithongtin()
	fmt.Println()
	gv.clear()
	fmt.Println()
	gv.hienthithongtin()
}

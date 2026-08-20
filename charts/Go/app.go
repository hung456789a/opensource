package main

import (
	"bufio"
	"fmt"
	"os"
)

var address = "Ho Chi Minh City"

var cource = "golang"

var (
	totalScore int
	monHoc     string
)

func main() {
	// fmt.Println("Hello World")
	// randomUser()

	// var fullName = "Nguyen Minh Hung"
	// fullName = "Nguyen Van A"
	// fmt.Println("User:", fullName)

	// var age int
	// fmt.Println("Age:", age)

	// phone := "0981864443"
	// fmt.Println("Phone:", phone)

	// fmt.Println("Address:", address)
	// fmt.Println("Cource:", cource)
	// // var toan, tiengviet, anhvan int
	// // toan = 10
	// // tiengviet = 9
	// // anhvan = 8
	// toan, tiengviet, anhvan := 5, 6, 7

	// fmt.Println("Toan:", toan)
	// fmt.Println("Tieng viet:", tiengviet)
	// fmt.Println("Anh van:", anhvan)
	// var hoten string
	// fmt.Print("nhap ho ten: ")
	// fmt.Scan(&hoten)
	// fmt.Println("ho ten:", hoten)
	var hoten string
	fmt.Print("Nhap Ho ten: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		hoten = scanner.Text()
	}
	fmt.Println("ho ten:", hoten)
}

package main

import "fmt"

func main() {
	// Khai báo biến
	var name string = "Go"
	var age int = 10
	
	// Type inference
	var city = "San Francisco"
	var year = 2009
	
	// Short declaration (chỉ trong function)
	language := "Golang"
	version := 1.21
	
	// Khai báo nhiều biến
	var x, y int = 1, 2
	var a, b = "hello", true
	c, d := 3, 4
	
	// Khai báo nhóm
	var (
		firstName string = "Go"
		lastName  string = "Language"
		isActive  bool   = true
	)
	
	// Zero values
	var i int       // 0
	var f float64   // 0.0
	var s string    // ""
	var b2 bool     // false
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)
	fmt.Println("Language:", language)
	fmt.Println("Version:", version)
	fmt.Println("x, y:", x, y)
	fmt.Println("a, b:", a, b)
	fmt.Println("c, d:", c, d)
	fmt.Println("Full name:", firstName, lastName)
	fmt.Println("Is active:", isActive)
	fmt.Println("Zero values:", i, f, s, b2)
}


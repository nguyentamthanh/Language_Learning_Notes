package main

import "fmt"

func main() {
	// If statement
	x := 10
	if x > 5 {
		fmt.Println("x lớn hơn 5")
	}
	
	// If-else
	age := 18
	if age >= 18 {
		fmt.Println("Bạn đã trưởng thành")
	} else {
		fmt.Println("Bạn chưa trưởng thành")
	}
	
	// If với initialization
	if err := doSomething(); err != nil {
		fmt.Println("Lỗi:", err)
	}
	
	// Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Thứ Hai")
	case "Tuesday":
		fmt.Println("Thứ Ba")
	default:
		fmt.Println("Ngày khác")
	}
	
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	
	// For như while
	i := 0
	for i < 5 {
		fmt.Printf("i = %d\n", i)
		i++
	}
	
	// Range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	
	// Break và Continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func doSomething() error {
	return nil
}


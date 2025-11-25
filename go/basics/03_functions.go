package main

import "fmt"

// Basic function
func greet() {
	fmt.Println("Hello!")
}

// Function với parameters
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function với return value
func add(a, b int) int {
	return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("không thể chia cho 0")
	}
	return a / b, nil
}

// Named return values
func calculate(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	greet()
	greetPerson("Go")
	
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	divResult, err := divide(10, 2)
	if err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", divResult)
	}
	
	s, p := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)
	
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
}


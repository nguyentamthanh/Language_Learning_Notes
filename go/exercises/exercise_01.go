package main

import "fmt"

// TODO: Viết hàm tính tổng của một slice số nguyên
func sumSlice(numbers []int) int {
	// Your code here
	return 0
}

// TODO: Viết hàm tìm số lớn nhất trong slice
func findMax(numbers []int) int {
	// Your code here
	return 0
}

// TODO: Viết hàm đếm số chẵn trong slice
func countEven(numbers []int) int {
	// Your code here
	return 0
}

// TODO: Viết hàm kiểm tra số nguyên tố
func isPrime(n int) bool {
	// Your code here
	return false
}

func main() {
	// Test cases
	numbers := []int{1, 2, 3, 4, 5}
	
	if sumSlice(numbers) == 15 {
		fmt.Println("✓ sumSlice passed")
	} else {
		fmt.Println("✗ sumSlice failed")
	}
	
	if findMax(numbers) == 5 {
		fmt.Println("✓ findMax passed")
	} else {
		fmt.Println("✗ findMax failed")
	}
	
	if countEven(numbers) == 2 {
		fmt.Println("✓ countEven passed")
	} else {
		fmt.Println("✗ countEven failed")
	}
	
	if isPrime(7) && !isPrime(10) {
		fmt.Println("✓ isPrime passed")
	} else {
		fmt.Println("✗ isPrime failed")
	}
}


package main

import "fmt"

// Giải pháp: Tính tổng của một slice số nguyên
func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Giải pháp: Tìm số lớn nhất trong slice
func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Giải pháp: Đếm số chẵn trong slice
func countEven(numbers []int) int {
	count := 0
	for _, num := range numbers {
		if num%2 == 0 {
			count++
		}
	}
	return count
}

// Giải pháp: Kiểm tra số nguyên tố
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
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


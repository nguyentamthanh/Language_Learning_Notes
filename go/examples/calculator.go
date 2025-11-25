package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\n=== Máy tính đơn giản ===")
		fmt.Println("1. Cộng")
		fmt.Println("2. Trừ")
		fmt.Println("3. Nhân")
		fmt.Println("4. Chia")
		fmt.Println("5. Thoát")
		
		fmt.Print("\nChọn phép toán (1-5): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		if choice == "5" {
			fmt.Println("Tạm biệt!")
			break
		}
		
		if choice == "1" || choice == "2" || choice == "3" || choice == "4" {
			fmt.Print("Nhập số thứ nhất: ")
			num1Str, _ := reader.ReadString('\n')
			num1, err1 := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)
			
			fmt.Print("Nhập số thứ hai: ")
			num2Str, _ := reader.ReadString('\n')
			num2, err2 := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)
			
			if err1 != nil || err2 != nil {
				fmt.Println("Lỗi: Vui lòng nhập số hợp lệ!")
				continue
			}
			
			var result float64
			var operator string
			
			switch choice {
			case "1":
				result = num1 + num2
				operator = "+"
			case "2":
				result = num1 - num2
				operator = "-"
			case "3":
				result = num1 * num2
				operator = "*"
			case "4":
				if num2 == 0 {
					fmt.Println("Lỗi: Không thể chia cho 0!")
					continue
				}
				result = num1 / num2
				operator = "/"
			}
			
			fmt.Printf("Kết quả: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
		} else {
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
}


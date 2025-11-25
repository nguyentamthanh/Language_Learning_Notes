# Chương 7: Functions

## Function Declaration

### Basic function

```go
func greet() {
    fmt.Println("Hello!")
}
```

### Function với parameters

```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Function với return value

```go
func add(a int, b int) int {
    return a + b
}

// Type có thể viết gọn nếu cùng kiểu
func add(a, b int) int {
    return a + b
}
```

### Multiple return values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("không thể chia cho 0")
    }
    return a / b, nil
}

// Sử dụng
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Lỗi:", err)
} else {
    fmt.Println("Kết quả:", result)
}
```

### Named return values

```go
func calculate(x, y int) (sum int, product int) {
    sum = x + y      // Không cần :=
    product = x * y
    return           // Naked return
}

// Hoặc
func calculate(x, y int) (sum int, product int) {
    return x + y, x * y
}
```

## Variadic Functions

Functions nhận số lượng arguments không cố định:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Sử dụng
fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15

// Truyền slice
numbers := []int{1, 2, 3}
fmt.Println(sum(numbers...))     // Spread operator
```

## Function Types

Functions là first-class citizens:

```go
// Định nghĩa function type
type Operation func(int, int) int

// Sử dụng
var add Operation = func(a, b int) int {
    return a + b
}

var multiply Operation = func(a, b int) int {
    return a * b
}

// Truyền function như parameter
func calculate(a, b int, op Operation) int {
    return op(a, b)
}

result := calculate(5, 3, add)  // 8
```

## Anonymous Functions

```go
// Anonymous function
func() {
    fmt.Println("Anonymous function")
}()

// Gán vào biến
greet := func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
greet("Go")

// Return anonymous function
func makeMultiplier(n int) func(int) int {
    return func(x int) int {
        return x * n
    }
}

double := makeMultiplier(2)
fmt.Println(double(5))  // 10
```

## Closures

Closure là function có thể truy cập biến từ scope bên ngoài:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
fmt.Println(c())  // 3
```

## Recursion

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5))  // 120
```

## Defer

`defer` thực thi function sau khi function hiện tại return:

```go
func example() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
// Output: Hello
//         World

// Defer với multiple statements
func example() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
}
// Output: Third
//         Second
//         First (LIFO - Last In First Out)
```

### Defer với file operations

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Đảm bảo file được đóng
    
    // Đọc file...
    return nil
}
```

## Methods

Methods là functions với receiver:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Sử dụng
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())  // 50
rect.Scale(2)
fmt.Println(rect.Area())  // 200
```

## Function Best Practices

### 1. Keep functions small

```go
// Tốt: Function nhỏ, làm một việc
func calculateTotal(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Price
    }
    return total
}
```

### 2. Use meaningful names

```go
// Tốt
func calculateTax(price float64) float64

// Không tốt
func calc(p float64) float64
```

### 3. Handle errors properly

```go
func process(data []byte) error {
    if len(data) == 0 {
        return fmt.Errorf("data is empty")
    }
    // Process...
    return nil
}
```

## Ví dụ thực tế

```go
package main

import (
    "fmt"
    "math"
)

// Function tính khoảng cách giữa 2 điểm
func distance(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Sqrt(dx*dx + dy*dy)
}

// Variadic function tính trung bình
func average(numbers ...float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    return sum / float64(len(numbers))
}

func main() {
    dist := distance(0, 0, 3, 4)
    fmt.Printf("Khoảng cách: %.2f\n", dist)  // 5.00
    
    avg := average(10, 20, 30, 40, 50)
    fmt.Printf("Trung bình: %.2f\n", avg)  // 30.00
}
```

## Tóm tắt

- Functions có thể return nhiều giá trị
- Variadic functions nhận số lượng arguments không cố định
- Functions là first-class citizens
- Closures có thể truy cập biến từ scope ngoài
- Defer thực thi sau khi function return
- Methods là functions với receiver

**Tiếp theo**: [Chương 8: Arrays và Slices](./08-arrays-slices.md)


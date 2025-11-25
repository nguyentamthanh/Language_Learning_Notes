# Chương 6: Cấu trúc điều khiển

## If Statement

### Basic if

```go
x := 10

if x > 5 {
    fmt.Println("x lớn hơn 5")
}
```

### If-else

```go
age := 18

if age >= 18 {
    fmt.Println("Bạn đã trưởng thành")
} else {
    fmt.Println("Bạn chưa trưởng thành")
}
```

### If-else if-else

```go
score := 85

if score >= 90 {
    fmt.Println("Xuất sắc")
} else if score >= 80 {
    fmt.Println("Giỏi")
} else if score >= 70 {
    fmt.Println("Khá")
} else {
    fmt.Println("Cần cố gắng")
}
```

### If với initialization

```go
// Khai báo biến trong if
if err := doSomething(); err != nil {
    fmt.Println("Lỗi:", err)
}

// err chỉ tồn tại trong scope của if
```

## Switch Statement

### Basic switch

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Thứ Hai")
case "Tuesday":
    fmt.Println("Thứ Ba")
case "Wednesday":
    fmt.Println("Thứ Tư")
default:
    fmt.Println("Ngày khác")
}
```

### Switch với multiple values

```go
day := "Monday"

switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Ngày làm việc")
case "Saturday", "Sunday":
    fmt.Println("Cuối tuần")
default:
    fmt.Println("Không hợp lệ")
}
```

### Switch không có expression

```go
x := 10

switch {
case x < 0:
    fmt.Println("Âm")
case x == 0:
    fmt.Println("Không")
case x > 0:
    fmt.Println("Dương")
}
```

### Switch với fallthrough

```go
x := 2

switch x {
case 1:
    fmt.Println("Một")
    fallthrough  // Tiếp tục case tiếp theo
case 2:
    fmt.Println("Hai")
case 3:
    fmt.Println("Ba")
}
// Output: Hai
```

### Type switch

```go
var i interface{} = "hello"

switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type\n")
}
```

## For Loop

### Basic for loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### For như while

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

### Infinite loop

```go
for {
    fmt.Println("Vô hạn")
    // Cần break để thoát
    break
}
```

### For-range (iterate collections)

```go
// Array/Slice
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Chỉ lấy value
for _, value := range numbers {
    fmt.Println(value)
}

// Chỉ lấy index
for index := range numbers {
    fmt.Println(index)
}

// Map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}

// String (iterate runes)
s := "Hello"
for index, char := range s {
    fmt.Printf("Index: %d, Char: %c\n", index, char)
}

// Channel
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    close(ch)
}()
for value := range ch {
    fmt.Println(value)
}
```

## Break và Continue

### Break

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Thoát khỏi loop
    }
    fmt.Println(i)
}
```

### Break với label

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 2 {
            break outer  // Thoát cả hai loops
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

### Continue

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // Bỏ qua iteration hiện tại
    }
    fmt.Println(i)  // Chỉ in số lẻ
}
```

### Continue với label

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            continue outer  // Bỏ qua j loop, tiếp tục i loop
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

## Goto (Không khuyến nghị)

```go
i := 0
loop:
    if i < 5 {
        fmt.Println(i)
        i++
        goto loop
    }
```

## Ví dụ thực tế

```go
package main

import "fmt"

func main() {
    // Tìm số đầu tiên chia hết cho 7
    for i := 1; i <= 100; i++ {
        if i%7 == 0 {
            fmt.Printf("Số đầu tiên chia hết cho 7: %d\n", i)
            break
        }
    }
    
    // In số lẻ từ 1 đến 10
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println(i)
    }
    
    // Iterate map
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Charlie": 92,
    }
    
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }
}
```

## Tóm tắt

- `if` có thể khai báo biến trong điều kiện
- `switch` linh hoạt hơn các ngôn ngữ khác
- `for` là loop duy nhất trong Go
- `range` để iterate collections
- `break` và `continue` hỗ trợ labels

**Tiếp theo**: [Chương 7: Functions](./07-functions.md)


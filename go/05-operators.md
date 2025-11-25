# Chương 5: Toán tử

## Arithmetic Operators (Toán tử số học)

```go
a := 10
b := 3

fmt.Println(a + b)   // 13 - Cộng
fmt.Println(a - b)   // 7  - Trừ
fmt.Println(a * b)   // 30 - Nhân
fmt.Println(a / b)   // 3  - Chia (integer division)
fmt.Println(a % b)   // 1  - Chia lấy phần dư
```

**Lưu ý**: Integer division trong Go:
```go
fmt.Println(10 / 3)    // 3 (không phải 3.333)
fmt.Println(10.0 / 3) // 3.333...
```

## Assignment Operators (Toán tử gán)

```go
x := 10

x += 5   // x = x + 5  → 15
x -= 3   // x = x - 3  → 12
x *= 2   // x = x * 2  → 24
x /= 4   // x = x / 4  → 6
x %= 2   // x = x % 2  → 0

// Bitwise assignment
x &= 5   // x = x & 5
x |= 3   // x = x | 3
x ^= 2   // x = x ^ 2
x <<= 1  // x = x << 1
x >>= 1  // x = x >> 1
```

## Comparison Operators (Toán tử so sánh)

```go
a := 10
b := 20

fmt.Println(a == b)  // false - Bằng
fmt.Println(a != b)  // true  - Khác
fmt.Println(a < b)   // true  - Nhỏ hơn
fmt.Println(a > b)   // false - Lớn hơn
fmt.Println(a <= b)  // true  - Nhỏ hơn hoặc bằng
fmt.Println(a >= b)  // false - Lớn hơn hoặc bằng
```

## Logical Operators (Toán tử logic)

```go
x := true
y := false

fmt.Println(x && y)  // false - Và
fmt.Println(x || y)   // true  - Hoặc
fmt.Println(!x)       // false - Phủ định

// Short-circuit evaluation
func getValue() bool {
    fmt.Println("Called")
    return true
}

true || getValue()   // getValue() không được gọi
false && getValue()  // getValue() không được gọi
```

## Bitwise Operators (Toán tử bit)

```go
a := 5   // 101 in binary
b := 3   // 011 in binary

fmt.Println(a & b)   // 1  - AND
fmt.Println(a | b)   // 7  - OR
fmt.Println(a ^ b)   // 6  - XOR
fmt.Println(^a)      // -6 - NOT (two's complement)
fmt.Println(a << 1)  // 10 - Left shift
fmt.Println(a >> 1)  // 2  - Right shift
```

## Address Operators (Toán tử địa chỉ)

```go
x := 10

// Address of operator
ptr := &x        // Lấy địa chỉ của x
fmt.Println(ptr) // Địa chỉ memory

// Dereference operator
*ptr = 20        // Thay đổi giá trị qua pointer
fmt.Println(x)   // 20
```

## Increment và Decrement

```go
x := 10

x++      // x = x + 1 → 11
x--      // x = x - 1 → 10

// Lưu ý: ++ và -- là statements, không phải expressions
// y := x++  // Lỗi!
```

## Operator Precedence (Thứ tự ưu tiên)

Thứ tự từ cao đến thấp:

1. `*`, `/`, `%`, `<<`, `>>`, `&`, `&^`
2. `+`, `-`, `|`, `^`
3. `==`, `!=`, `<`, `<=`, `>`, `>=`
4. `&&`
5. `||`

```go
result := 2 + 3 * 4      // 14 (không phải 20)
result = (2 + 3) * 4    // 20
```

## Type-specific Operators

### String Operators

```go
s1 := "Hello"
s2 := " World"

// Concatenation
s3 := s1 + s2  // "Hello World"

// Comparison
fmt.Println(s1 < s2)  // So sánh lexicographically
```

### Channel Operators

```go
ch := make(chan int)

// Send
ch <- 10

// Receive
value := <-ch
```

### Slice Operators

```go
s := []int{1, 2, 3, 4, 5}

// Slicing
s[1:3]  // [2, 3]
s[:3]   // [1, 2, 3]
s[2:]   // [3, 4, 5]
```

## Ví dụ thực tế

```go
package main

import "fmt"

func main() {
    // Tính toán đơn giản
    price := 100.0
    discount := 0.1
    finalPrice := price * (1 - discount)
    fmt.Printf("Giá cuối: %.2f\n", finalPrice)
    
    // Kiểm tra số chẵn/lẻ
    number := 10
    if number%2 == 0 {
        fmt.Println("Số chẵn")
    } else {
        fmt.Println("Số lẻ")
    }
    
    // Bitwise operations
    flags := 5  // 101
    if flags&1 != 0 {
        fmt.Println("Bit đầu tiên được set")
    }
}
```

## Tóm tắt

- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Assignment: `=`, `+=`, `-=`, `*=`, `/=`, `%=`
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical: `&&`, `||`, `!`
- Bitwise: `&`, `|`, `^`, `<<`, `>>`
- Address: `&`, `*`
- Increment/Decrement: `++`, `--`

**Tiếp theo**: [Chương 6: Cấu trúc điều khiển](./06-control-flow.md)


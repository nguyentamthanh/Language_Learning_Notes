# Chương 4: Biến và Kiểu dữ liệu

## Variables (Biến)

### Khai báo biến

```go
// Cách 1: Khai báo rõ kiểu
var name string = "Go"
var age int = 10

// Cách 2: Type inference
var name = "Go"
var age = 10

// Cách 3: Short declaration (chỉ trong function)
name := "Go"
age := 10

// Cách 4: Khai báo nhiều biến
var x, y int = 1, 2
var a, b = "hello", true
x, y := 1, 2

// Cách 5: Khai báo nhóm
var (
    name string = "Go"
    age  int    = 10
    isActive bool = true
)
```

### Zero Values

Mỗi kiểu có giá trị mặc định khi khai báo:

```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
var p *int      // nil
var sl []int    // nil
var m map[string]int // nil
```

## Basic Types

### Numeric Types

```go
// Integers
var i8 int8    = 127           // -128 to 127
var i16 int16  = 32767         // -32768 to 32767
var i32 int32  = 2147483647    // -2^31 to 2^31-1
var i64 int64  = 9223372036854775807 // -2^63 to 2^63-1
var i int      // Platform dependent (32 or 64 bit)

// Unsigned integers
var u8 uint8   = 255           // 0 to 255
var u16 uint16 = 65535         // 0 to 65535
var u32 uint32 = 4294967295    // 0 to 2^32-1
var u64 uint64 = 18446744073709551615 // 0 to 2^64-1
var u uint     // Platform dependent

// Special integers
var byte uint8  // Alias for uint8
var rune int32  // Alias for int32 (Unicode code point)

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Complex numbers
var c64 complex64  = 1 + 2i
var c128 complex128 = 1 + 2i
```

### Boolean

```go
var isTrue bool = true
var isFalse bool = false
```

### String

```go
// String literals
var s1 string = "Hello"
var s2 string = `Multi-line
string`

// String là immutable
s := "Hello"
// s[0] = 'h' // Lỗi!

// String concatenation
s1 := "Hello"
s2 := " World"
s3 := s1 + s2  // "Hello World"

// String length
len("Hello")  // 5

// Access characters (byte)
s := "Hello"
s[0]  // 'H' (byte)
```

### Rune và Unicode

```go
// Rune = Unicode code point (int32)
var r rune = 'A'
var r2 rune = '中'

// String to runes
s := "Hello"
runes := []rune(s)

// Rune to string
r := 'A'
s := string(r)
```

## Type Conversion

Go yêu cầu explicit conversion:

```go
var i int = 42
var f float64 = float64(i)  // Phải explicit
var u uint = uint(f)

// String conversion
var s string = string(65)  // "A"
var num string = strconv.Itoa(42)  // "42"
var i int, _ = strconv.Atoi("42")  // 42
```

## Constants

```go
// Single constant
const pi = 3.14159

// Typed constant
const pi float64 = 3.14159

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)

// iota (auto-increment)
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
)

const (
    KB = 1024
    MB = KB * 1024
    GB = MB * 1024
)
```

## Type Assertion

```go
var i interface{} = "hello"

// Type assertion
s := i.(string)        // "hello"
s, ok := i.(string)    // "hello", true
f, ok := i.(float64)   // 0, false
```

## Type Switch

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

## Ví dụ thực tế

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Type conversion
    ageStr := "25"
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        fmt.Println("Lỗi chuyển đổi")
        return
    }
    fmt.Printf("Tuổi: %d\n", age)
    
    // String formatting
    name := "Go"
    message := fmt.Sprintf("Xin chào %s, bạn %d tuổi", name, age)
    fmt.Println(message)
}
```

## Tóm tắt

- Go có nhiều kiểu số nguyên và số thực
- String là immutable
- Rune đại diện cho Unicode code point
- Type conversion phải explicit
- Constants có thể dùng iota để tự động tăng

**Tiếp theo**: [Chương 5: Toán tử](./05-operators.md)


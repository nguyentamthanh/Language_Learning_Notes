# Chương 10: Structs

## Struct là gì?

Struct là kiểu dữ liệu tổng hợp chứa các fields có thể có kiểu khác nhau. Tương tự class trong các ngôn ngữ OOP nhưng không có inheritance.

## Định nghĩa Struct

```go
// Định nghĩa struct
type Person struct {
    Name string
    Age  int
    City string
}

// Khởi tạo struct
var p Person
p.Name = "Alice"
p.Age = 25
p.City = "New York"

// Khởi tạo với struct literal
p := Person{
    Name: "Bob",
    Age:  30,
    City: "San Francisco",
}

// Khởi tạo theo thứ tự (không khuyến nghị)
p := Person{"Charlie", 35, "London"}

// Khởi tạo một phần
p := Person{Name: "David"}
```

## Truy cập Fields

```go
p := Person{Name: "Alice", Age: 25}
fmt.Println(p.Name)                // "Alice"
fmt.Println(p.Age)                 // 25

// Gán giá trị
p.Age = 26
```

## Struct Embedding

Go không có inheritance nhưng có embedding:

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal                      // Embedded struct
    Breed string
}

// Sử dụng
dog := Dog{
    Animal: Animal{Name: "Buddy"},
    Breed:  "Golden Retriever",
}

fmt.Println(dog.Name)           // "Buddy" (truy cập qua embedding)
fmt.Println(dog.Breed)           // "Golden Retriever"
```

## Anonymous Structs

```go
// Struct không có tên
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  25,
}
```

## Struct với Tags

Tags dùng cho encoding/decoding (JSON, XML, etc.):

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city,omitempty"`
}
```

## Pointer to Struct

```go
p := &Person{Name: "Alice", Age: 25}

// Truy cập qua pointer (Go tự động dereference)
fmt.Println(p.Name)              // "Alice"
fmt.Println((*p).Name)           // Tương tự

// Thay đổi qua pointer
p.Age = 26
```

## Methods với Struct

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver (khuyến nghị khi cần modify)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Sử dụng
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())         // 50

rect.Scale(2)
fmt.Println(rect.Area())         // 200
```

## Value vs Pointer Receiver

### Value Receiver
- Copy struct khi gọi method
- Không thể modify struct gốc
- An toàn cho concurrent access

### Pointer Receiver
- Không copy, làm việc trực tiếp với struct
- Có thể modify struct gốc
- Hiệu quả hơn với struct lớn

```go
// Value receiver
func (r Rectangle) Area() float64 { ... }

// Pointer receiver
func (r *Rectangle) Scale(factor float64) { ... }
```

## Struct Comparison

Structs có thể so sánh nếu tất cả fields có thể so sánh:

```go
p1 := Person{Name: "Alice", Age: 25}
p2 := Person{Name: "Alice", Age: 25}

fmt.Println(p1 == p2)            // true

// Không thể so sánh nếu có slice, map, function
```

## Nested Structs

```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Age     int
    Address Address
}

p := Person{
    Name: "Alice",
    Age:  25,
    Address: Address{
        Street: "123 Main St",
        City:   "New York",
    },
}
```

## Ví dụ thực tế

```go
package main

import "fmt"

type BankAccount struct {
    AccountNumber string
    Balance      float64
    Owner         string
}

func (ba *BankAccount) Deposit(amount float64) {
    ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) bool {
    if ba.Balance >= amount {
        ba.Balance -= amount
        return true
    }
    return false
}

func (ba BankAccount) GetBalance() float64 {
    return ba.Balance
}

func main() {
    account := BankAccount{
        AccountNumber: "123456",
        Balance:      1000.0,
        Owner:        "Alice",
    }
    
    account.Deposit(500)
    fmt.Printf("Balance sau khi deposit: %.2f\n", account.GetBalance())
    
    success := account.Withdraw(200)
    if success {
        fmt.Printf("Balance sau khi withdraw: %.2f\n", account.GetBalance())
    }
}
```

## Tóm tắt

- Struct là kiểu dữ liệu tổng hợp
- Không có inheritance, dùng embedding
- Methods có thể có value hoặc pointer receiver
- Structs có thể so sánh nếu fields có thể so sánh
- Tags dùng cho encoding/decoding

**Tiếp theo**: [Chương 11: Pointers](./11-pointers.md)


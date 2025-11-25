# Chương 9: Maps

## Map là gì?

Map là cấu trúc dữ liệu key-value, tương tự dictionary trong Python hoặc HashMap trong Java.

## Khai báo Map

```go
// Khai báo map rỗng
var m map[string]int              // nil map

// Khởi tạo với make
m := make(map[string]int)

// Khởi tạo với giá trị
m := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Charlie": 35,
}

// Khởi tạo với make và capacity
m := make(map[string]int, 10)
```

## Thao tác với Map

### Thêm và cập nhật

```go
m := make(map[string]int)
m["Alice"] = 25                   // Thêm
m["Bob"] = 30
m["Alice"] = 26                   // Cập nhật
```

### Truy cập

```go
m := map[string]int{"Alice": 25}

// Truy cập trực tiếp (trả về zero value nếu không tồn tại)
age := m["Alice"]                 // 25
age2 := m["Bob"]                  // 0 (zero value)

// Kiểm tra tồn tại
age, exists := m["Alice"]
if exists {
    fmt.Println("Alice:", age)
}

// Chỉ kiểm tra tồn tại
_, exists := m["Bob"]
if !exists {
    fmt.Println("Bob không tồn tại")
}
```

### Xóa

```go
m := map[string]int{"Alice": 25, "Bob": 30}
delete(m, "Alice")                // Xóa key "Alice"
```

### Length

```go
m := map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(len(m))               // 2
```

## Iterate Map

```go
m := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Charlie": 35,
}

// Iterate với range
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Chỉ lấy keys
for key := range m {
    fmt.Println(key)
}

// Chỉ lấy values
for _, value := range m {
    fmt.Println(value)
}
```

**Lưu ý**: Thứ tự iteration không được đảm bảo!

## Map với các kiểu khác nhau

### Map of slices

```go
m := map[string][]int{
    "even": {2, 4, 6},
    "odd":  {1, 3, 5},
}
```

### Map of maps

```go
m := map[string]map[string]int{
    "Alice": {"age": 25, "score": 95},
    "Bob":   {"age": 30, "score": 87},
}
```

### Map với struct keys

```go
type Point struct {
    X, Y int
}

m := make(map[Point]string)
m[Point{1, 2}] = "First"
m[Point{3, 4}] = "Second"
```

## Nil Map

```go
var m map[string]int              // nil map

// Có thể đọc từ nil map (trả về zero value)
fmt.Println(m["key"])              // 0

// Nhưng không thể ghi vào nil map (panic!)
// m["key"] = 10                   // Panic!

// Phải khởi tạo trước
m = make(map[string]int)
m["key"] = 10                     // OK
```

## Common Patterns

### Đếm tần suất

```go
func countFrequency(words []string) map[string]int {
    freq := make(map[string]int)
    for _, word := range words {
        freq[word]++
    }
    return freq
}
```

### Group by

```go
func groupByAge(people []Person) map[int][]Person {
    groups := make(map[int][]Person)
    for _, person := range people {
        groups[person.Age] = append(groups[person.Age], person)
    }
    return groups
}
```

### Set (sử dụng map với bool)

```go
// Tạo set
set := make(map[string]bool)
set["apple"] = true
set["banana"] = true

// Kiểm tra membership
if set["apple"] {
    fmt.Println("apple exists")
}

// Xóa
delete(set, "apple")
```

## Map Performance

- **Lookup**: O(1) average case
- **Insert**: O(1) average case
- **Delete**: O(1) average case
- **Iterate**: O(n)

## Best Practices

### 1. Kiểm tra tồn tại khi đọc

```go
// Tốt
if value, ok := m["key"]; ok {
    // Sử dụng value
}

// Không tốt (nếu zero value là giá trị hợp lệ)
value := m["key"]  // Không biết key có tồn tại không
```

### 2. Pre-allocate khi biết size

```go
// Tốt
m := make(map[string]int, 100)

// Không tốt
m := make(map[string]int)  // Nhiều lần rehash
```

### 3. Không dựa vào thứ tự

```go
// Thứ tự iteration không được đảm bảo
for key, value := range m {
    // Không giả định thứ tự
}
```

## Ví dụ thực tế

```go
package main

import "fmt"

func main() {
    // Tạo map điểm số
    scores := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
    }
    
    // Thêm điểm mới
    scores["David"] = 88
    
    // Tìm điểm cao nhất
    maxScore := 0
    topStudent := ""
    for name, score := range scores {
        if score > maxScore {
            maxScore = score
            topStudent = name
        }
    }
    
    fmt.Printf("Học sinh điểm cao nhất: %s với %d điểm\n", topStudent, maxScore)
    
    // Đếm số học sinh đạt điểm >= 90
    count := 0
    for _, score := range scores {
        if score >= 90 {
            count++
        }
    }
    fmt.Printf("Số học sinh đạt >= 90 điểm: %d\n", count)
}
```

## Tóm tắt

- Map là cấu trúc key-value
- Zero value của map là `nil`
- Phải khởi tạo map trước khi ghi
- Sử dụng `ok` để kiểm tra tồn tại
- Thứ tự iteration không được đảm bảo

**Tiếp theo**: [Chương 10: Structs](./10-structs.md)


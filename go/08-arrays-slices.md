# Chương 8: Arrays và Slices

## Arrays

Array là tập hợp các phần tử cùng kiểu có kích thước cố định:

```go
// Khai báo array
var numbers [5]int                    // [0 0 0 0 0]
var names [3]string = [3]string{"Alice", "Bob", "Charlie"}

// Short declaration
numbers := [5]int{1, 2, 3, 4, 5}
numbers := [...]int{1, 2, 3}          // Compiler tự xác định size

// Truy cập phần tử
numbers[0] = 10
fmt.Println(numbers[0])               // 10

// Length
fmt.Println(len(numbers))              // 5

// Iterate
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}

for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

**Lưu ý**: Arrays trong Go là value types, không phải reference types.

## Slices

Slice là reference đến một phần của array, linh hoạt hơn array:

```go
// Tạo slice từ array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]                     // [2 3 4]

// Tạo slice trực tiếp
slice := []int{1, 2, 3, 4, 5}

// Slice rỗng
var slice []int                       // nil slice
slice := make([]int, 5)               // [0 0 0 0 0]
slice := make([]int, 0, 10)           // Length 0, capacity 10
```

## Slice Operations

### Length và Capacity

```go
slice := []int{1, 2, 3, 4, 5}
fmt.Println(len(slice))               // 5
fmt.Println(cap(slice))               // 5

// Slice với capacity lớn hơn
slice := make([]int, 3, 10)
fmt.Println(len(slice))               // 3
fmt.Println(cap(slice))               // 10
```

### Append

```go
slice := []int{1, 2, 3}
slice = append(slice, 4)              // [1 2 3 4]
slice = append(slice, 5, 6, 7)       // [1 2 3 4 5 6 7]

// Append slice vào slice
slice1 := []int{1, 2}
slice2 := []int{3, 4}
slice1 = append(slice1, slice2...)   // [1 2 3 4]
```

### Copy

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, len(src))
copy(dst, src)                        // Copy src vào dst
```

### Slicing

```go
slice := []int{1, 2, 3, 4, 5}

slice[1:3]    // [2 3] - từ index 1 đến 3 (không bao gồm 3)
slice[:3]     // [1 2 3] - từ đầu đến index 3
slice[2:]     // [3 4 5] - từ index 2 đến cuối
slice[:]      // [1 2 3 4 5] - toàn bộ slice
```

## Slice Internals

Slice gồm 3 thành phần:
- Pointer: Trỏ đến phần tử đầu tiên
- Length: Số phần tử
- Capacity: Dung lượng tối đa

```go
slice := make([]int, 3, 10)
// Pointer → [0 0 0 _ _ _ _ _ _ _]
// Length: 3
// Capacity: 10
```

## Common Patterns

### Filter

```go
func filterEven(numbers []int) []int {
    result := []int{}
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}
```

### Map

```go
func double(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * 2
    }
    return result
}
```

### Reduce

```go
func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

## 2D Slices

```go
// 2D slice
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Truy cập
fmt.Println(matrix[1][2])            // 6

// Tạo 2D slice động
rows, cols := 3, 4
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols)
}
```

## Best Practices

### 1. Pre-allocate khi biết size

```go
// Tốt
result := make([]int, 0, 100)
for i := 0; i < 100; i++ {
    result = append(result, i)
}

// Không tốt
var result []int
for i := 0; i < 100; i++ {
    result = append(result, i)       // Nhiều lần reallocate
}
```

### 2. Kiểm tra nil slice

```go
var slice []int
if slice == nil {
    fmt.Println("Slice is nil")
}
```

### 3. Sử dụng copy để tránh aliasing

```go
original := []int{1, 2, 3}
copy := make([]int, len(original))
copy(copy, original)
```

## Ví dụ thực tế

```go
package main

import "fmt"

func main() {
    // Tạo và sử dụng slice
    numbers := []int{1, 2, 3, 4, 5}
    
    // Thêm phần tử
    numbers = append(numbers, 6, 7, 8)
    
    // Filter số chẵn
    evens := []int{}
    for _, num := range numbers {
        if num%2 == 0 {
            evens = append(evens, num)
        }
    }
    
    fmt.Println("Evens:", evens)      // [2 4 6 8]
    
    // Slice operations
    fmt.Println(numbers[2:5])         // [3 4 5]
    fmt.Println(numbers[:3])         // [1 2 3]
    fmt.Println(numbers[5:])         // [6 7 8]
}
```

## Tóm tắt

- Arrays có kích thước cố định, là value types
- Slices linh hoạt hơn, là reference types
- `append()` để thêm phần tử vào slice
- `copy()` để copy slice
- Slicing với `[start:end]` syntax

**Tiếp theo**: [Chương 9: Maps](./09-maps.md)


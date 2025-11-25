# Chương 3: Cú pháp cơ bản

## Package Declaration

Mọi file Go phải bắt đầu với package declaration:

```go
package main  // Package chính, có thể chạy được

package utils // Package thông thường
```

## Imports

```go
// Single import
import "fmt"

// Multiple imports
import (
    "fmt"
    "math"
    "strings"
)

// Import với alias
import f "fmt"
f.Println("Hello")

// Import nhưng không dùng (để side effects)
import _ "database/sql/driver"
```

## Comments

```go
// Single line comment

/*
Multi-line comment
*/

// Package-level comment
// Package main provides ...
package main
```

## Statements và Expressions

```go
// Statement: Thực hiện hành động
fmt.Println("Hello")

// Expression: Tính toán giá trị
x := 10 + 5
```

## Semicolons

Go tự động thêm semicolons, nhưng có thể viết tường minh:

```go
// Đúng
x := 10

// Cũng đúng
x := 10;

// Sai (semicolon sau if)
if x > 10; {  // Lỗi!
}
```

## Naming Conventions

```go
// Exported (public): Bắt đầu bằng chữ hoa
func PublicFunction() {}
var PublicVariable = 10

// Unexported (private): Bắt đầu bằng chữ thường
func privateFunction() {}
var privateVariable = 10

// Constants: UPPER_SNAKE_CASE
const MAX_SIZE = 100

// Interfaces: Thường kết thúc bằng -er
type Reader interface {}
type Writer interface {}
```

## Code Organization

### Một file Go đơn giản

```go
package main

import "fmt"

const greeting = "Hello"

func main() {
    fmt.Println(greeting)
}
```

### Thứ tự trong file Go

1. Package declaration
2. Imports
3. Constants
4. Variables
5. Types
6. Functions

## Ví dụ đầy đủ

```go
package main

import (
    "fmt"
    "math"
)

const pi = 3.14159

var radius float64 = 5.0

func main() {
    area := calculateArea(radius)
    fmt.Printf("Diện tích hình tròn: %.2f\n", area)
}

func calculateArea(r float64) float64 {
    return math.Pi * r * r
}
```

## Tóm tắt

- Package declaration là bắt buộc
- Imports có thể single hoặc grouped
- Comments giống C/C++
- Naming conventions quan trọng cho visibility
- Code được tổ chức theo thứ tự nhất định

**Tiếp theo**: [Chương 4: Biến và Kiểu dữ liệu](./04-variables-types.md)


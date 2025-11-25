# Chương 1: Giới thiệu Go

## Go là gì?

Go (hay Golang) là ngôn ngữ lập trình được phát triển bởi Google vào năm 2009. Được thiết kế bởi Robert Griesemer, Rob Pike, và Ken Thompson. Go kết hợp hiệu năng của C/C++ với sự đơn giản của Python.

## Tại sao học Go?

### ✅ Ưu điểm

1. **Đơn giản**: Cú pháp gọn gàng, dễ học
2. **Hiệu năng cao**: Compile thành binary, chạy nhanh
3. **Concurrency**: Goroutines và Channels tích hợp sẵn
4. **Tooling tốt**: gofmt, go test, go build
5. **Static typing**: Phát hiện lỗi sớm
6. **Garbage collection**: Tự động quản lý memory

### 📊 Ứng dụng của Go

- **Backend Services**: API servers, Microservices
- **Cloud Native**: Docker, Kubernetes, Terraform
- **CLI Tools**: Fast command-line applications
- **Web Development**: Gin, Echo, Fiber frameworks
- **DevOps Tools**: Monitoring, logging tools
- **Distributed Systems**: High-performance systems

## Đặc điểm của Go

### 1. Compiled Language
Go được compile thành machine code, không cần runtime như Python hay Node.js.

### 2. Statically Typed
Kiểu dữ liệu được kiểm tra tại compile time, giảm lỗi runtime.

### 3. Garbage Collected
Tự động quản lý memory, không cần manual memory management như C/C++.

### 4. Concurrency Built-in
Goroutines và Channels là first-class citizens trong Go.

### 5. Simple Syntax
Ít keywords (chỉ 25), cú pháp rõ ràng, dễ đọc.

## Hello World đầu tiên

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Xin chào Go!")
}
```

### Giải thích:
- `package main`: Định nghĩa package, `main` là entry point
- `import "fmt"`: Import package fmt để in ra màn hình
- `func main()`: Entry point của chương trình
- `fmt.Println()`: In ra màn hình và xuống dòng

## Go vs Các ngôn ngữ khác

### Go vs Python
- **Go**: Compiled, nhanh hơn, static typing
- **Python**: Interpreted, dễ học hơn, dynamic typing

### Go vs Java
- **Go**: Đơn giản hơn, compile nhanh hơn, không có JVM
- **Java**: OOP thuần túy, ecosystem lớn hơn

### Go vs C++
- **Go**: Garbage collected, đơn giản hơn, concurrency tốt
- **C++**: Kiểm soát memory tốt hơn, hiệu năng cao hơn

## Cộng đồng và Ecosystem

- **GitHub**: Hơn 100k repositories
- **Companies**: Google, Uber, Dropbox, Docker, Kubernetes
- **Frameworks**: Gin, Echo, Fiber, Beego
- **Tools**: Docker, Kubernetes, Terraform, Prometheus

## Tóm tắt

- Go là ngôn ngữ hiện đại, đơn giản và mạnh mẽ
- Được thiết kế cho concurrency và performance
- Phù hợp cho backend services và cloud native
- Cộng đồng đang phát triển mạnh

**Tiếp theo**: [Chương 2: Cài đặt và Môi trường](./02-setup.md)


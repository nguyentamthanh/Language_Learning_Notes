# Chương 19: Goroutines

## Goroutine là gì?

Goroutine là lightweight thread được quản lý bởi Go runtime. Goroutines rất nhẹ (chỉ vài KB stack), có thể tạo hàng nghìn goroutines.

## Tạo Goroutine

Sử dụng keyword `go` trước function call:

```go
// Function thông thường
func sayHello() {
    fmt.Println("Hello")
}

// Chạy trong goroutine
go sayHello()

// Anonymous function trong goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()

// Function với parameters
go func(name string) {
    fmt.Printf("Hello, %s\n", name)
}("Go")
```

## Ví dụ cơ bản

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Chạy trong goroutine
    go sayHello()
    
    // Main goroutine tiếp tục
    fmt.Println("Main function")
    time.Sleep(1 * time.Second)  // Đợi goroutine hoàn thành
}

func sayHello() {
    fmt.Println("Hello from goroutine")
}
```

## Goroutines và Concurrency

```go
func main() {
    for i := 0; i < 5; i++ {
        go func(n int) {
            fmt.Printf("Goroutine %d\n", n)
        }(i)
    }
    
    time.Sleep(1 * time.Second)
}
```

**Lưu ý**: Goroutines chạy concurrently, thứ tự không được đảm bảo!

## WaitGroup

`sync.WaitGroup` để đợi nhiều goroutines hoàn thành:

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)  // Tăng counter
        go func(n int) {
            defer wg.Done()  // Giảm counter khi xong
            fmt.Printf("Goroutine %d\n", n)
        }(i)
    }
    
    wg.Wait()  // Đợi tất cả goroutines hoàn thành
    fmt.Println("All goroutines finished")
}
```

## Goroutine với Channels

```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello"
    }()
    
    msg := <-ch
    fmt.Println(msg)
}
```

## Best Practices

### 1. Luôn có cách để goroutine kết thúc

```go
// Tốt: Sử dụng context hoặc channel để signal
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Do work
        }
    }
}()
```

### 2. Tránh goroutine leaks

```go
// Không tốt: Goroutine không bao giờ kết thúc
go func() {
    for {
        // Infinite loop
    }
}()

// Tốt: Có điều kiện dừng
done := make(chan bool)
go func() {
    for {
        select {
        case <-done:
            return
        default:
            // Do work
        }
    }
}()
```

### 3. Sử dụng buffered channels khi cần

```go
// Buffered channel để không block
ch := make(chan int, 10)
```

## Ví dụ thực tế

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
}
```

## Tóm tắt

- Goroutines là lightweight threads
- Tạo với keyword `go`
- Sử dụng WaitGroup để đợi nhiều goroutines
- Cần có cách để goroutine kết thúc
- Tránh goroutine leaks

**Tiếp theo**: [Chương 20: Channels](./20-channels.md)


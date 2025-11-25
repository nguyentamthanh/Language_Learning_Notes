# Chương 20: Channels

## Channel là gì?

Channel là cách để goroutines giao tiếp với nhau. Channel cho phép gửi và nhận giá trị giữa các goroutines một cách an toàn.

## Tạo Channel

```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel (capacity 10)
ch := make(chan int, 10)

// Channel với type cụ thể
ch := make(chan string)
```

## Gửi và Nhận

```go
ch := make(chan int)

// Gửi (send)
ch <- 10

// Nhận (receive)
value := <-ch

// Nhận và kiểm tra channel đóng
value, ok := <-ch
if !ok {
    fmt.Println("Channel đã đóng")
}
```

## Unbuffered vs Buffered Channels

### Unbuffered Channel

```go
ch := make(chan int)

// Block cho đến khi có receiver
ch <- 10  // Block ở đây nếu không có receiver

// Block cho đến khi có sender
value := <-ch  // Block ở đây nếu không có sender
```

### Buffered Channel

```go
ch := make(chan int, 3)  // Buffer size 3

// Không block nếu buffer chưa đầy
ch <- 1
ch <- 2
ch <- 3
ch <- 4  // Block ở đây vì buffer đầy

// Nhận từ buffer
value := <-ch
```

## Đóng Channel

```go
ch := make(chan int)

// Đóng channel
close(ch)

// Kiểm tra channel đã đóng
value, ok := <-ch
if !ok {
    fmt.Println("Channel đã đóng")
}
```

**Lưu ý**: 
- Chỉ sender mới đóng channel
- Gửi vào channel đã đóng sẽ panic
- Nhận từ channel đã đóng sẽ nhận zero value và ok = false

## Range over Channel

```go
ch := make(chan int)

go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}()

// Range tự động dừng khi channel đóng
for value := range ch {
    fmt.Println(value)
}
```

## Select Statement

`select` cho phép chờ nhiều channel operations:

```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case ch1 <- "hello":
    fmt.Println("Sent to ch1")
default:
    fmt.Println("No communication")
}
```

### Select với timeout

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
}
```

### Select với context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-ctx.Done():
    fmt.Println("Context cancelled")
}
```

## Channel Directions

```go
// Send-only channel
func sendOnly(ch chan<- int) {
    ch <- 10
    // value := <-ch  // Lỗi!
}

// Receive-only channel
func receiveOnly(ch <-chan int) {
    value := <-ch
    // ch <- 10  // Lỗi!
}

// Bidirectional (mặc định)
func bidirectional(ch chan int) {
    ch <- 10
    value := <-ch
}
```

## Common Patterns

### Worker Pool

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 5; r++ {
        fmt.Println(<-results)
    }
}
```

### Fan-out/Fan-in

```go
// Fan-out: Một channel → nhiều goroutines
func fanOut(input <-chan int, outputs []chan<- int) {
    for value := range input {
        for _, out := range outputs {
            out <- value
        }
    }
}

// Fan-in: Nhiều channels → một channel
func fanIn(inputs []<-chan int, output chan<- int) {
    var wg sync.WaitGroup
    for _, in := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for value := range ch {
                output <- value
            }
        }(in)
    }
    go func() {
        wg.Wait()
        close(output)
    }()
}
```

## Ví dụ thực tế

```go
package main

import (
    "fmt"
    "time"
)

func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Printf("Sent: %d\n", i)
        time.Sleep(100 * time.Millisecond)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    ch := make(chan int, 2)  // Buffered channel
    
    go producer(ch)
    consumer(ch)
    
    fmt.Println("Done")
}
```

## Tóm tắt

- Channels để giao tiếp giữa goroutines
- Unbuffered channels block
- Buffered channels có capacity
- `close()` để đóng channel
- `select` để chờ nhiều channels
- Channel directions: `chan<-` (send), `<-chan` (receive)

**Tiếp theo**: [Chương 21: Select Statement](./21-select.md)


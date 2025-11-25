# Chương 13: Error Handling

## Rust Error Handling Philosophy

Rust không có exceptions. Thay vào đó, Rust sử dụng:
- `Result<T, E>` cho recoverable errors
- `panic!` cho unrecoverable errors

## panic!

`panic!` dừng program khi gặp lỗi không thể xử lý:

```rust
panic!("Something went wrong!");

// Hoặc từ code
let v = vec![1, 2, 3];
v[99];  // Panic! Index out of bounds
```

## Result<T, E>

`Result` enum cho phép xử lý errors:

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

## Sử dụng Result

```rust
use std::fs::File;

fn main() {
    let f = File::open("hello.txt");
    
    let f = match f {
        Ok(file) => file,
        Err(error) => {
            panic!("Problem opening file: {:?}", error);
        }
    };
}
```

## Matching Different Errors

```rust
use std::fs::File;
use std::io::ErrorKind;

fn main() {
    let f = File::open("hello.txt");
    
    let f = match f {
        Ok(file) => file,
        Err(error) => match error.kind() {
            ErrorKind::NotFound => match File::create("hello.txt") {
                Ok(fc) => fc,
                Err(e) => panic!("Problem creating file: {:?}", e),
            },
            other_error => {
                panic!("Problem opening file: {:?}", other_error);
            }
        }
    };
}
```

## Shortcuts: unwrap và expect

```rust
// unwrap: Panic nếu Err
let f = File::open("hello.txt").unwrap();

// expect: Panic với message tùy chỉnh
let f = File::open("hello.txt")
    .expect("Failed to open hello.txt");
```

## Propagating Errors

Sử dụng `?` operator để propagate errors:

```rust
use std::fs::File;
use std::io;
use std::io::Read;

fn read_username_from_file() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

// Hoặc ngắn gọn hơn
fn read_username_from_file() -> Result<String, io::Error> {
    let mut s = String::new();
    File::open("hello.txt")?.read_to_string(&mut s)?;
    Ok(s)
}
```

## ? Operator

`?` operator:
- Nếu `Ok`, unwrap và return value
- Nếu `Err`, return error early

```rust
fn main() -> Result<(), Box<dyn std::error::Error>> {
    let f = File::open("hello.txt")?;
    Ok(())
}
```

## Custom Error Types

```rust
use std::fmt;

#[derive(Debug)]
struct MyError {
    message: String,
}

impl fmt::Display for MyError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.message)
    }
}

impl std::error::Error for MyError {}
```

## Result và Option

```rust
// Option methods
let x = Some(5);
x.map(|v| v * 2);           // Some(10)
x.and_then(|v| Some(v * 2)); // Some(10)

// Result methods
let x: Result<i32, &str> = Ok(5);
x.map(|v| v * 2);           // Ok(10)
x.map_err(|e| format!("Error: {}", e));
```

## Ví dụ thực tế

```rust
use std::fs::File;
use std::io::{self, Read};

fn read_file(filename: &str) -> Result<String, io::Error> {
    let mut file = File::open(filename)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}

fn main() {
    match read_file("hello.txt") {
        Ok(contents) => println!("File contents: {}", contents),
        Err(e) => eprintln!("Error reading file: {}", e),
    }
}
```

## Tóm tắt

- `panic!` cho unrecoverable errors
- `Result<T, E>` cho recoverable errors
- `unwrap()` và `expect()` để panic
- `?` operator để propagate errors
- Có thể tạo custom error types

**Tiếp theo**: [Chương 14: Generics](./14-generics.md)


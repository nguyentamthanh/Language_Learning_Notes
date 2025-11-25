# Chương 9: Enums và Pattern Matching

## Enums

Enum cho phép định nghĩa một type có thể là một trong nhiều variants:

```rust
enum IpAddrKind {
    V4,
    V6,
}

// Sử dụng
let four = IpAddrKind::V4;
let six = IpAddrKind::V6;
```

## Enums với Data

Enums có thể chứa data:

```rust
enum IpAddr {
    V4(String),
    V6(String),
}

let home = IpAddr::V4(String::from("127.0.0.1"));
let loopback = IpAddr::V6(String::from("::1"));
```

## Enums với Different Types

Mỗi variant có thể có type khác nhau:

```rust
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}
```

## Option Enum

`Option<T>` là enum quan trọng trong Rust:

```rust
enum Option<T> {
    Some(T),
    None,
}

// Sử dụng
let some_number = Some(5);
let some_string = Some("hello");
let absent_number: Option<i32> = None;
```

## Match Expression

`match` là cách xử lý enums:

```rust
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}
```

## Match với Data

```rust
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
}

fn handle_message(msg: Message) {
    match msg {
        Message::Quit => {
            println!("Quit");
        }
        Message::Move { x, y } => {
            println!("Move to ({}, {})", x, y);
        }
        Message::Write(s) => {
            println!("Write: {}", s);
        }
    }
}
```

## Match với Option

```rust
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

let five = Some(5);
let six = plus_one(five);
let none = plus_one(None);
```

## Exhaustive Matching

Match phải cover tất cả cases:

```rust
// Lỗi! Thiếu case None
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        Some(i) => Some(i + 1),
        // Thiếu None case
    }
}
```

## Catch-all Patterns

Sử dụng `_` để catch tất cả cases còn lại:

```rust
let dice_roll = 9;
match dice_roll {
    3 => add_fancy_hat(),
    7 => remove_fancy_hat(),
    other => move_player(other),  // Catch-all
}

// Hoặc bỏ qua giá trị
match dice_roll {
    3 => add_fancy_hat(),
    7 => remove_fancy_hat(),
    _ => reroll(),  // Bỏ qua giá trị
}
```

## if let

`if let` là syntax sugar cho match một pattern:

```rust
let some_value = Some(3);

// Với match
match some_value {
    Some(3) => println!("three"),
    _ => (),
}

// Với if let (ngắn gọn hơn)
if let Some(3) = some_value {
    println!("three");
}
```

## while let

Tương tự `if let` nhưng cho loops:

```rust
let mut stack = Vec::new();
stack.push(1);
stack.push(2);
stack.push(3);

while let Some(top) = stack.pop() {
    println!("{}", top);
}
```

## Ví dụ thực tế

```rust
enum UsState {
    Alabama,
    Alaska,
    // ...
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        }
    }
}
```

## Tóm tắt

- Enums định nghĩa type có thể là một trong nhiều variants
- Enums có thể chứa data
- `Option<T>` là enum quan trọng cho nullable values
- `match` để xử lý tất cả cases
- `if let` và `while let` là syntax sugar

**Tiếp theo**: [Chương 10: Functions](./10-functions.md)


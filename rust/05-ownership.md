# Chương 5: Ownership

## Ownership là gì?

Ownership là hệ thống quản lý memory độc đáo của Rust. Mỗi giá trị có một owner duy nhất. Khi owner ra khỏi scope, giá trị được giải phóng tự động.

## Ownership Rules

1. Mỗi giá trị có một owner duy nhất
2. Chỉ có một owner tại một thời điểm
3. Khi owner ra khỏi scope, giá trị bị drop

## Variable Scope

```rust
{
    let s = String::from("hello");  // s vào scope
    // s có thể sử dụng ở đây
}  // s ra khỏi scope, memory được giải phóng
```

## Move Semantics

Với types có ownership (như String, Vec), assignment sẽ move giá trị:

```rust
let s1 = String::from("hello");
let s2 = s1;              // s1 bị move vào s2
// println!("{}", s1);    // Lỗi! s1 không còn valid

// Copy types (i32, bool, char, tuples của copy types)
let x = 5;
let y = x;                // Copy, không move
println!("{}", x);        // OK, x vẫn valid
```

## Copy vs Move

### Copy Types
Types implement `Copy` trait được copy thay vì move:

```rust
// Copy types
let x = 5;
let y = x;                // Copy
println!("{}", x);        // OK

// Copy types: integers, bool, char, floats, tuples của copy types
```

### Move Types
Types không implement `Copy` bị move:

```rust
// Move types
let s1 = String::from("hello");
let s2 = s1;              // Move
// println!("{}", s1);    // Lỗi!
```

## Ownership và Functions

```rust
fn main() {
    let s = String::from("hello");
    takes_ownership(s);    // s bị move vào function
    // println!("{}", s);  // Lỗi! s không còn valid
    
    let x = 5;
    makes_copy(x);        // x được copy
    println!("{}", x);    // OK, x vẫn valid
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}  // some_string ra khỏi scope, memory được giải phóng

fn makes_copy(some_integer: i32) {
    println!("{}", some_integer);
}
```

## Return Values và Ownership

```rust
fn main() {
    let s1 = gives_ownership();        // s1 nhận ownership
    
    let s2 = String::from("hello");
    let s3 = takes_and_gives_back(s2); // s2 bị move, s3 nhận ownership
}

fn gives_ownership() -> String {
    let some_string = String::from("hello");
    some_string                         // Move ownership ra ngoài
}

fn takes_and_gives_back(a_string: String) -> String {
    a_string                            // Move ownership ra ngoài
}
```

## Multiple Return Values

```rust
fn main() {
    let s1 = String::from("hello");
    let (s2, len) = calculate_length(s1);
    println!("'{}' có độ dài {}", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();
    (s, length)                        // Return ownership
}
```

## References và Borrowing

Thay vì transfer ownership, có thể borrow:

```rust
fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1);   // Borrow, không move
    println!("'{}' có độ dài {}", s1, len);  // s1 vẫn valid
}

fn calculate_length(s: &String) -> usize {
    s.len()
}  // s ra khỏi scope nhưng không drop vì không có ownership
```

## Mutable References

```rust
fn main() {
    let mut s = String::from("hello");
    change(&mut s);
    println!("{}", s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```

## Rules của References

1. Tại một thời điểm, có thể có:
   - Một mutable reference, HOẶC
   - Nhiều immutable references
2. References phải luôn valid

```rust
let mut s = String::from("hello");

let r1 = &s;              // OK
let r2 = &s;              // OK
// let r3 = &mut s;        // Lỗi! Không thể có mutable và immutable cùng lúc

let r3 = &mut s;          // OK sau khi r1, r2 không còn dùng
```

## Dangling References

Rust ngăn chặn dangling references tại compile time:

```rust
// Lỗi! Không compile được
fn dangle() -> &String {
    let s = String::from("hello");
    &s                      // Lỗi! s sẽ bị drop
}
```

## Slices

Slices là reference đến một phần của collection:

```rust
let s = String::from("hello world");
let hello = &s[0..5];     // &str slice
let world = &s[6..11];
```

## Ví dụ thực tế

```rust
fn main() {
    let s = String::from("hello world");
    let word = first_word(&s);
    println!("First word: {}", word);
}

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();
    
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    
    &s[..]
}
```

## Tóm tắt

- Mỗi giá trị có một owner duy nhất
- Assignment với non-Copy types sẽ move
- References cho phép borrow mà không transfer ownership
- Chỉ có một mutable reference hoặc nhiều immutable references
- Rust ngăn chặn dangling references tại compile time

**Tiếp theo**: [Chương 6: Borrowing và References](./06-borrowing.md)


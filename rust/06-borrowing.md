# Chương 6: Borrowing và References

## Borrowing là gì?

Borrowing cho phép bạn sử dụng giá trị mà không transfer ownership. References cho phép truy cập giá trị mà không sở hữu nó.

## References

### Immutable References

```rust
fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1);  // &s1 tạo reference
    println!("'{}' có độ dài {}", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}  // s ra khỏi scope nhưng không drop vì không có ownership
```

### Mutable References

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

1. **Tại một thời điểm**, bạn có thể có:
   - Một mutable reference, HOẶC
   - Nhiều immutable references
2. References phải luôn valid (không được dangling)

```rust
let mut s = String::from("hello");

let r1 = &s;        // OK
let r2 = &s;        // OK
// let r3 = &mut s;  // Lỗi! Không thể có mutable và immutable cùng lúc

let r3 = &mut s;    // OK sau khi r1, r2 không còn dùng
```

## Dangling References

Rust compiler ngăn chặn dangling references:

```rust
// Lỗi! Không compile được
fn dangle() -> &String {
    let s = String::from("hello");
    &s  // Lỗi! s sẽ bị drop khi function return
}

// Đúng: Return owned value
fn no_dangle() -> String {
    let s = String::from("hello");
    s  // Move ownership ra ngoài
}
```

## String Slices

String slice `&str` là reference đến một phần của String:

```rust
let s = String::from("hello world");
let hello = &s[0..5];      // &str
let world = &s[6..11];
let slice = &s[..];        // Toàn bộ string
```

## Ví dụ thực tế

```rust
fn main() {
    let mut s = String::from("hello world");
    let word = first_word(&s);
    println!("First word: {}", word);
    
    s.clear();  // Lỗi! Không thể clear khi có immutable reference
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

- References cho phép borrow mà không transfer ownership
- Immutable references: `&T`
- Mutable references: `&mut T`
- Chỉ một mutable reference hoặc nhiều immutable references
- Rust ngăn chặn dangling references tại compile time

**Tiếp theo**: [Chương 7: Slices](./07-slices.md)


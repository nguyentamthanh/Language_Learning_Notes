# Chương 3: Cú pháp cơ bản

## Main Function

Mọi Rust program phải có function `main`:

```rust
fn main() {
    println!("Hello, World!");
}
```

## Comments (Ghi chú)

```rust
// Single line comment

/* Multi-line comment
   có thể viết nhiều dòng */

/// Documentation comment cho items
/// Sử dụng cho `cargo doc`

//! Documentation comment cho module/crate
```

## Print Macros

```rust
fn main() {
    // println! - In và xuống dòng
    println!("Hello, World!");
    
    // print! - In không xuống dòng
    print!("Hello");
    print!(" World");
    
    // eprintln! - In ra stderr
    eprintln!("Error message");
    
    // Format string
    let name = "Rust";
    println!("Hello, {}!", name);
    
    // Multiple arguments
    println!("{} + {} = {}", 5, 3, 5 + 3);
    
    // Positional arguments
    println!("{0} is {1}, {0} is cool", "Rust", "fast");
    
    // Named arguments
    println!("{name} is {age} years old", name = "Rust", age = 10);
    
    // Formatting
    println!("Number: {:05}", 42);        // 00042
    println!("Hex: {:x}", 255);           // ff
    println!("Binary: {:b}", 15);        // 1111
}
```

## Semicolons

Rust sử dụng semicolons để kết thúc statements:

```rust
let x = 5;        // Statement - cần semicolon
let y = x + 3;    // Statement

// Expression không cần semicolon (trong function return)
fn add(a: i32, b: i32) -> i32 {
    a + b         // Expression - không có semicolon
}
```

## Naming Conventions

```rust
// Variables và functions: snake_case
let my_variable = 10;
fn my_function() {}

// Types và structs: PascalCase
struct MyStruct {}
enum MyEnum {}

// Constants: SCREAMING_SNAKE_CASE
const MAX_SIZE: i32 = 100;

// Modules: snake_case
mod my_module {}

// Lifetimes: lowercase với apostrophe
fn foo<'a>(x: &'a str) {}
```

## Keywords

Rust có các keywords không thể dùng làm identifiers:

```rust
// Keywords quan trọng
as, async, await, break, const, continue, crate, dyn, else, 
enum, extern, false, fn, for, if, impl, in, let, loop, match, 
mod, move, mut, pub, ref, return, Self, self, static, struct, 
super, trait, true, type, unsafe, use, where, while

// Reserved keywords (chưa dùng nhưng dự trữ)
abstract, become, box, do, final, macro, override, priv, try, 
typeof, unsized, virtual, yield
```

## Expressions vs Statements

### Statements
Thực hiện hành động, không trả về giá trị:

```rust
let x = 5;        // Statement
x + 3;            // Expression statement (kết quả bị discard)
```

### Expressions
Tính toán giá trị:

```rust
5 + 3             // Expression
if x > 0 { 1 } else { 0 }  // Expression
```

## Blocks

Block là expression, có thể trả về giá trị:

```rust
let x = {
    let y = 5;
    y + 3         // Giá trị của block
};                // x = 8

// If expression
let result = if x > 0 {
    1
} else {
    0
};
```

## Macros

Macros kết thúc bằng `!`:

```rust
println!("Hello");        // Macro
vec![1, 2, 3];           // Macro
format!("{}", x);         // Macro
```

## Modules

```rust
// Định nghĩa module
mod my_module {
    pub fn public_function() {}
    fn private_function() {}
}

// Sử dụng module
use my_module::public_function;
```

## Attributes

Attributes cung cấp metadata:

```rust
#[derive(Debug)]          // Derive macro
#[allow(dead_code)]       // Cho phép unused code
#[cfg(test)]              // Chỉ compile khi test
```

## Ví dụ đầy đủ

```rust
// Module-level documentation
//! This is a simple Rust program

/// Function tính tổng hai số
fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn main() {
    let x = 10;
    let y = 20;
    
    let sum = add(x, y);
    println!("{} + {} = {}", x, y, sum);
    
    // Block expression
    let result = {
        let temp = x * y;
        temp / 2
    };
    println!("Result: {}", result);
}
```

## Tóm tắt

- `fn main()` là entry point
- Comments: `//`, `/* */`, `///`, `//!`
- `println!` và `print!` để in ra màn hình
- Semicolons kết thúc statements
- Expressions có thể trả về giá trị
- Macros kết thúc bằng `!`
- Naming conventions: snake_case, PascalCase, SCREAMING_SNAKE_CASE

**Tiếp theo**: [Chương 4: Biến và Kiểu dữ liệu](./04-variables-types.md)


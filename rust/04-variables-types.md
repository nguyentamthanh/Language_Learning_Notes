# Chương 4: Biến và Kiểu dữ liệu

## Variables (Biến)

### Khai báo biến

```rust
// Immutable variable (mặc định)
let x = 5;
let name = "Rust";

// Mutable variable
let mut y = 10;
y = 20;  // OK vì mut

// Type annotation
let x: i32 = 5;
let name: &str = "Rust";

// Multiple variables
let (x, y) = (5, 10);
let (a, b, c) = (1, 2, 3);
```

### Shadowing

Có thể khai báo lại biến với cùng tên:

```rust
let x = 5;
let x = x + 1;        // Shadowing
let x = x * 2;        // Shadowing lại

// Có thể thay đổi type
let spaces = "   ";
let spaces = spaces.len();  // spaces giờ là usize
```

## Scalar Types

### Integers

```rust
// Signed integers
let i8: i8 = 127;           // -128 to 127
let i16: i16 = 32767;       // -32768 to 32767
let i32: i32 = 2147483647;  // -2^31 to 2^31-1
let i64: i64 = 9223372036854775807; // -2^63 to 2^63-1
let i128: i128 = ...;       // -2^127 to 2^127-1
let isize: isize = ...;     // Architecture dependent

// Unsigned integers
let u8: u8 = 255;           // 0 to 255
let u16: u16 = 65535;       // 0 to 65535
let u32: u32 = 4294967295; // 0 to 2^32-1
let u64: u64 = ...;         // 0 to 2^64-1
let u128: u128 = ...;       // 0 to 2^128-1
let usize: usize = ...;     // Architecture dependent

// Literals
let decimal = 98_222;       // 98222
let hex = 0xff;            // 255
let octal = 0o77;          // 63
let binary = 0b1111_0000;   // 240
let byte = b'A';           // u8: 65
```

### Floating Point

```rust
let f32: f32 = 3.14;
let f64: f64 = 3.141592653589793;  // Default

// Scientific notation
let x = 1.5e2;              // 150.0
```

### Boolean

```rust
let t = true;
let f: bool = false;
```

### Character

```rust
let c = 'z';
let z = 'ℤ';
let heart_eyed_cat = '😻';

// Character là Unicode scalar value (4 bytes)
```

## Compound Types

### Tuples

```rust
// Khai báo tuple
let tup: (i32, f64, u8) = (500, 6.4, 1);

// Destructuring
let (x, y, z) = tup;
println!("x: {}, y: {}, z: {}", x, y, z);

// Truy cập bằng index
let five_hundred = tup.0;
let six_point_four = tup.1;
let one = tup.2;

// Tuple rỗng
let empty = ();
```

### Arrays

```rust
// Array có kích thước cố định
let a = [1, 2, 3, 4, 5];
let a: [i32; 5] = [1, 2, 3, 4, 5];

// Khởi tạo với giá trị giống nhau
let a = [3; 5];            // [3, 3, 3, 3, 3]

// Truy cập
let first = a[0];
let second = a[1];

// Length
let len = a.len();
```

## String Types

### &str (String Slice)

```rust
// String literal
let s = "Hello";           // &str
let s: &str = "Hello";

// String slice từ String
let s = String::from("Hello");
let slice: &str = &s[0..2];  // "He"
```

### String (Owned String)

```rust
// Tạo String
let s = String::from("Hello");
let s = "Hello".to_string();

// Mutate String
let mut s = String::from("Hello");
s.push_str(" World");
s.push('!');

// Concatenation
let s1 = String::from("Hello");
let s2 = String::from(" World");
let s3 = s1 + &s2;         // s1 bị move

// Format
let s = format!("{} {}", "Hello", "World");
```

## Type Conversion

```rust
// Explicit conversion
let x = 5;
let y = x as f64;          // 5.0

// Parse từ string
let num: i32 = "42".parse().expect("Not a number!");
let num = "42".parse::<i32>().unwrap();

// From/Into traits
let s = String::from("Hello");
let s: String = "Hello".into();
```

## Constants

```rust
// Constants phải có type annotation
const MAX_POINTS: u32 = 100_000;

// Constants có thể ở global scope
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
```

## Type Inference

Rust có type inference mạnh:

```rust
let x = 5;                 // i32 (default)
let y = 5.0;               // f64 (default)
let z = true;              // bool

// Với context, compiler có thể infer
let mut vec = Vec::new();
vec.push(1);               // vec: Vec<i32>
```

## Ví dụ thực tế

```rust
fn main() {
    // Tuples
    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("x: {}, y: {}, z: {}", x, y, z);
    
    // Arrays
    let months = ["January", "February", "March"];
    println!("First month: {}", months[0]);
    
    // Strings
    let mut s = String::from("Hello");
    s.push_str(" Rust");
    println!("{}", s);
    
    // Type conversion
    let num_str = "42";
    let num: i32 = num_str.parse().unwrap();
    println!("Parsed number: {}", num);
}
```

## Tóm tắt

- Variables mặc định là immutable, dùng `mut` để mutable
- Shadowing cho phép khai báo lại biến
- Integers: i8-i128, u8-u128, isize, usize
- Floating point: f32, f64
- Tuples và Arrays là compound types
- String vs &str: owned vs borrowed
- Constants phải có type annotation

**Tiếp theo**: [Chương 5: Ownership](./05-ownership.md)


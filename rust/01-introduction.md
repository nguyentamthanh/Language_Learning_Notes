# Chương 1: Giới thiệu Rust

## Rust là gì?

Rust là ngôn ngữ lập trình hệ thống được phát triển bởi Mozilla Research. Được thiết kế để cung cấp hiệu năng của C/C++ với sự an toàn về memory mà không cần garbage collector.

## Tại sao học Rust?

### ✅ Ưu điểm

1. **Memory Safety**: Không có null pointer, use-after-free, data races
2. **Hiệu năng cao**: Gần với C/C++, không có runtime overhead
3. **Concurrency**: An toàn về concurrency, không có data races
4. **Zero-cost Abstractions**: Abstractions không làm giảm hiệu năng
5. **Modern Language**: Pattern matching, traits, async/await
6. **Excellent Tooling**: Cargo, rustfmt, clippy, rust-analyzer

### 📊 Ứng dụng của Rust

- **Systems Programming**: Operating systems, Drivers
- **Web Assembly**: High-performance web apps
- **Blockchain**: Solana, Polkadot
- **Web Servers**: Actix-web, Rocket, Axum
- **CLI Tools**: ripgrep, fd, bat
- **Game Engines**: Bevy, Amethyst
- **Embedded Systems**: IoT, Microcontrollers

## Rust vs C/C++

### Rust Advantages
- Memory safety tại compile time
- Không có undefined behavior
- Concurrency an toàn
- Modern tooling (Cargo)

### C/C++ Advantages
- Mature ecosystem
- More control over memory
- Wider adoption

## Hello World đầu tiên

```rust
fn main() {
    println!("Hello, World!");
    println!("Xin chào Rust!");
}
```

### Giải thích:
- `fn main()`: Entry point của chương trình
- `println!()`: Macro để in ra màn hình
- `!` sau tên function nghĩa là macro, không phải function

## Các khái niệm cốt lõi

### 1. Ownership
Mỗi giá trị có một owner duy nhất. Khi owner ra khỏi scope, giá trị được giải phóng.

### 2. Borrowing
Thay vì transfer ownership, có thể "mượn" giá trị với references.

### 3. Lifetimes
Đảm bảo references luôn hợp lệ.

### 4. Zero-cost Abstractions
Các abstraction như iterators compile thành code hiệu quả như hand-written loops.

## Rust Ecosystem

### Cargo
Package manager và build tool:
```bash
cargo new my_project
cargo build
cargo run
cargo test
```

### Crates.io
Registry của các packages (crates).

### Rustup
Toolchain installer và version manager.

## Rust Editions

Rust sử dụng editions để giới thiệu breaking changes:

- **2015**: Edition đầu tiên
- **2018**: Module system mới, async/await preview
- **2021**: Current stable, improved macros
- **2024**: Đang phát triển

## Cộng đồng Rust

- **GitHub**: Hơn 100k repositories
- **Companies**: Microsoft, Google, Amazon, Facebook
- **Projects**: Firefox, Dropbox, Cloudflare, Discord
- **Community**: r/rust, Rust Discord, Rust Users Forum

## Learning Curve

Rust có learning curve dốc hơn các ngôn ngữ khác vì:
- Ownership system độc đáo
- Lifetimes phức tạp
- Compiler strict

Nhưng một khi hiểu được, bạn sẽ viết code an toàn và hiệu quả hơn!

## Tóm tắt

- Rust là ngôn ngữ hệ thống an toàn và hiệu năng cao
- Ownership và Borrowing là concepts độc đáo
- Compiler giúp phát hiện lỗi sớm
- Cộng đồng đang phát triển mạnh
- Được sử dụng bởi nhiều công ty lớn

**Tiếp theo**: [Chương 2: Cài đặt và Môi trường](./02-setup.md)


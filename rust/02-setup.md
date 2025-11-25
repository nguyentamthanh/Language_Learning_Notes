# Chương 2: Cài đặt và Môi trường

## Cài đặt Rust

### Windows

1. Tải và chạy [rustup-init.exe](https://rustup.rs/)
2. Hoặc chạy trong PowerShell:
```powershell
# Tải và chạy installer
Invoke-WebRequest https://win.rustup.rs/x86_64 -OutFile rustup-init.exe
.\rustup-init.exe
```

### macOS và Linux

```bash
# Cài đặt rustup
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Hoặc với wget
wget https://sh.rustup.rs -O - | sh
```

Sau khi cài đặt, restart terminal hoặc chạy:
```bash
source $HOME/.cargo/env
```

## Kiểm tra cài đặt

```bash
# Kiểm tra Rust version
rustc --version

# Kiểm tra Cargo version
cargo --version

# Xem thông tin toolchain
rustup show
```

## Cập nhật Rust

```bash
# Cập nhật Rust lên phiên bản mới nhất
rustup update

# Cập nhật lên nightly (nếu cần)
rustup update nightly
```

## Cấu hình Rust

### Toolchains

```bash
# Xem toolchains có sẵn
rustup toolchain list

# Cài đặt toolchain cụ thể
rustup toolchain install stable
rustup toolchain install nightly

# Set default toolchain
rustup default stable

# Sử dụng toolchain cho project cụ thể
rustup override set nightly
```

### Components

```bash
# Cài đặt components
rustup component add rustfmt      # Code formatter
rustup component add clippy       # Linter
rustup component add rust-analyzer # LSP

# Cài đặt cho target cụ thể
rustup target add wasm32-unknown-unknown
rustup target add x86_64-pc-windows-gnu
```

## IDE và Editor

### 1. VS Code (Khuyến nghị)

Cài extensions:
- **rust-analyzer**: Language server
- **CodeLLDB**: Debugger
- **Better TOML**: TOML syntax highlighting

### 2. IntelliJ IDEA / CLion

Cài plugin: **Rust**

### 3. Vim/Neovim

Cài plugin: **rust.vim** hoặc **rust-tools.nvim**

### 4. Emacs

Cài package: **rust-mode**, **lsp-mode**

## Tạo project đầu tiên

### Với Cargo (Khuyến nghị)

```bash
# Tạo binary project
cargo new my_project

# Tạo library project
cargo new --lib my_lib

# Tạo project trong thư mục hiện tại
cargo init

# Cấu trúc project
my_project/
├── Cargo.toml          # Manifest file
├── Cargo.lock          # Lock file (tự động tạo)
└── src/
    └── main.rs         # Entry point
```

### Cargo.toml

```toml
[package]
name = "my_project"
version = "0.1.0"
edition = "2021"

[dependencies]
# Thêm dependencies ở đây
```

## Các lệnh Cargo cơ bản

```bash
# Build project
cargo build

# Build và chạy
cargo run

# Build release (optimized)
cargo build --release

# Chạy tests
cargo test

# Format code
cargo fmt

# Lint code
cargo clippy

# Xem documentation
cargo doc --open

# Thêm dependency
cargo add serde

# Xem dependencies
cargo tree
```

## Cấu trúc project Rust

```
my_project/
├── Cargo.toml              # Manifest
├── Cargo.lock              # Lock file
├── src/
│   ├── main.rs             # Binary entry point
│   ├── lib.rs              # Library entry point
│   └── bin/                # Additional binaries
│       └── other_bin.rs
├── tests/                  # Integration tests
│   └── integration_test.rs
├── examples/               # Example binaries
│   └── example.rs
└── benches/                # Benchmark tests
    └── benchmark.rs
```

## Environment Variables

```bash
# Rust toolchain directory
export RUSTUP_HOME=$HOME/.rustup

# Cargo home directory
export CARGO_HOME=$HOME/.cargo

# Add to PATH
export PATH="$CARGO_HOME/bin:$PATH"
```

## Rust Playground

Không cần cài đặt, chạy code online:
- [play.rust-lang.org](https://play.rust-lang.org/)

## Tóm tắt

- Rustup là cách cài đặt Rust chính thức
- Cargo là package manager và build tool
- rust-analyzer là LSP tốt nhất cho Rust
- VS Code với rust-analyzer là lựa chọn tốt
- Cargo.toml quản lý dependencies

**Tiếp theo**: [Chương 3: Cú pháp cơ bản](./03-basic-syntax.md)


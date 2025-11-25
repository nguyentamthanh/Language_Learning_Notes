# Chương 1: Giới thiệu C++

## C++ là gì?

C++ là ngôn ngữ lập trình đa mục đích, được phát triển bởi Bjarne Stroustrup tại Bell Labs vào năm 1979. Ban đầu được gọi là "C with Classes", sau đổi tên thành C++.

## Tại sao học C++?

### ✅ Ưu điểm

1. **Hiệu năng cao**: Gần với hardware, không có overhead
2. **Kiểm soát Memory**: Manual memory management
3. **Đa mô hình**: OOP, Procedural, Generic programming
4. **STL**: Standard Template Library mạnh mẽ
5. **Ứng dụng rộng**: System programming, Games, Embedded

### 📊 Ứng dụng của C++

- **System Programming**: Operating systems, Drivers
- **Game Development**: Unreal Engine, Game engines
- **Embedded Systems**: IoT, Microcontrollers
- **High-Performance Computing**: Scientific computing
- **Desktop Applications**: Qt, Windows applications
- **Database Systems**: MySQL, MongoDB

## C++ Standards

- **C++98/03**: Phiên bản đầu tiên chuẩn hóa
- **C++11**: Modern C++ với nhiều tính năng mới
- **C++14**: Cải thiện C++11
- **C++17**: Filesystem, Parallel algorithms
- **C++20**: Concepts, Ranges, Coroutines
- **C++23**: Đang phát triển

## Hello World đầu tiên

```cpp
#include <iostream>
using namespace std;

int main() {
    cout << "Hello, World!" << endl;
    cout << "Xin chào C++!" << endl;
    return 0;
}
```

### Giải thích:
- `#include <iostream>`: Include thư viện input/output
- `using namespace std`: Sử dụng namespace std
- `int main()`: Entry point của program
- `cout`: Standard output stream
- `return 0`: Trả về 0 (success)

## Các khái niệm cơ bản

### 1. Compiled Language
C++ được compile thành machine code, chạy trực tiếp trên hardware.

### 2. Manual Memory Management
Lập trình viên tự quản lý memory (có thể dùng smart pointers).

### 3. Multi-paradigm
Hỗ trợ nhiều mô hình lập trình: OOP, Procedural, Generic.

### 4. Zero-cost Abstractions
Các abstraction không làm giảm hiệu năng.

## C++ vs C

- **C++** là superset của C với OOP và nhiều tính năng khác
- **C** là ngôn ngữ procedural đơn giản hơn
- C++ có thể sử dụng code C

## Tóm tắt

- C++ là ngôn ngữ hiệu năng cao
- Cho phép kiểm soát memory và hardware
- Được sử dụng trong nhiều lĩnh vực đòi hỏi hiệu suất
- Có STL mạnh mẽ và cộng đồng lớn

**Tiếp theo**: [Chương 2: Cài đặt Compiler](./02-setup.md)


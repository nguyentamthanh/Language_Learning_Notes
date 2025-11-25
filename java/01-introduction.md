# Chương 1: Giới thiệu Java

## Java là gì?

Java là ngôn ngữ lập trình hướng đối tượng, được phát triển bởi Sun Microsystems (hiện thuộc Oracle) vào năm 1995. Khẩu hiệu: "Write Once, Run Anywhere" (WORA).

## Tại sao học Java?

### ✅ Ưu điểm

1. **Platform Independent**: Chạy trên mọi hệ điều hành nhờ JVM
2. **OOP**: Lập trình hướng đối tượng thuần túy
3. **Enterprise**: Được sử dụng rộng rãi trong doanh nghiệp
4. **Android**: Ngôn ngữ chính cho Android development
5. **Cộng đồng lớn**: Nhiều thư viện và framework

### 📊 Ứng dụng của Java

- **Enterprise Applications**: Banking, E-commerce
- **Android Development**: Mobile apps
- **Web Applications**: Spring, Java EE
- **Big Data**: Hadoop, Spark
- **IoT**: Embedded systems
- **Game Development**: LibGDX, jMonkeyEngine

## Java Platform

### JVM (Java Virtual Machine)
- Thực thi bytecode
- Quản lý memory
- Platform independent

### JRE (Java Runtime Environment)
- JVM + Libraries
- Cần để chạy Java applications

### JDK (Java Development Kit)
- JRE + Development tools (compiler, debugger)
- Cần để phát triển Java applications

## Phiên bản Java

- **Java 8** (LTS): Phổ biến nhất
- **Java 11** (LTS): Long-term support
- **Java 17** (LTS): Phiên bản LTS mới nhất
- **Java 21+**: Các tính năng mới

## Hello World đầu tiên

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        System.out.println("Xin chào Java!");
    }
}
```

### Giải thích:
- `public class HelloWorld`: Định nghĩa class công khai
- `public static void main(String[] args)`: Entry point
- `System.out.println()`: In ra màn hình

## Các khái niệm cơ bản

### 1. Object-Oriented
Mọi thứ trong Java đều là object (trừ primitive types).

### 2. Compiled Language
Java code được compile thành bytecode, sau đó chạy trên JVM.

### 3. Strongly Typed
Phải khai báo kiểu dữ liệu rõ ràng.

### 4. Garbage Collection
Tự động quản lý memory, không cần manual memory management.

## Tóm tắt

- Java là ngôn ngữ OOP mạnh mẽ
- Platform independent nhờ JVM
- Được sử dụng rộng rãi trong enterprise
- Cộng đồng lớn và hỗ trợ tốt

**Tiếp theo**: [Chương 2: Cài đặt JDK](./02-setup.md)


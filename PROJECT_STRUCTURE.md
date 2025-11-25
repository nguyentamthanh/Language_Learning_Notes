# 📂 Cấu trúc dự án

## Tổng quan

Dự án được tổ chức thành các thư mục ngôn ngữ, mỗi ngôn ngữ là một "quyển sách" hoàn chỉnh.

## Cấu trúc thư mục

```
Language_Learning_Notes/
│
├── README.md                 # README chính của dự án
├── CONTRIBUTING.md          # Hướng dẫn đóng góp
├── GITHUB_SETUP.md          # Hướng dẫn đẩy lên GitHub
├── LICENSE                  # Giấy phép MIT
├── PROJECT_STRUCTURE.md     # File này
│
├── python/                  # 🐍 Python
│   ├── README.md            # Mục lục Python
│   ├── 01-introduction.md   # Chương 1: Giới thiệu
│   ├── 02-setup.md          # Chương 2: Cài đặt
│   ├── 03-basic-syntax.md   # Chương 3: Cú pháp cơ bản
│   ├── 04-variables-types.md
│   ├── 05-operators.md
│   ├── basics/              # Code examples cơ bản
│   ├── examples/            # Ví dụ thực tế
│   └── exercises/           # Bài tập
│
├── javascript/              # 🟨 JavaScript
│   ├── README.md
│   ├── 01-introduction.md
│   ├── basics/
│   ├── examples/
│   └── exercises/
│
├── java/                    # ☕ Java
│   ├── README.md
│   ├── 01-introduction.md
│   ├── basics/
│   ├── examples/
│   └── exercises/
│
├── cpp/                     # ⚙️ C++
│   ├── README.md
│   ├── 01-introduction.md
│   ├── basics/
│   ├── examples/
│   └── exercises/
│
├── typescript/              # 🔷 TypeScript
│   ├── README.md
│   └── basics/
│
├── go/                      # 🐹 Go
│   ├── README.md
│   └── basics/
│
├── rust/                    # 🦀 Rust
│   ├── README.md
│   └── basics/
│
├── csharp/                  # 🔷 C#
│   └── README.md
│
├── ruby/                    # 💎 Ruby
│   └── README.md
│
├── php/                     # 🐘 PHP
│   └── README.md
│
├── swift/                   # 🦉 Swift
│   └── README.md
│
├── kotlin/                  # 🔷 Kotlin
│   └── README.md
│
└── scripts/                 # Scripts hỗ trợ
    └── setup-git.sh         # Script setup Git
```

## Cấu trúc mỗi ngôn ngữ

Mỗi ngôn ngữ được tổ chức như một quyển sách:

### 1. README.md
- Tổng quan về ngôn ngữ
- Mục lục đầy đủ các chương
- Hướng dẫn bắt đầu
- Tài liệu tham khảo

### 2. Các chương (01-, 02-, ...)
- Được đánh số theo thứ tự
- Mỗi chương là một file markdown độc lập
- Có ví dụ code và giải thích chi tiết
- Link đến chương trước/sau

### 3. Thư mục con

#### basics/
- Code examples cơ bản
- Các file .py, .js, .java, .cpp, etc.
- Được comment rõ ràng

#### examples/
- Ví dụ thực tế
- Ứng dụng hoàn chỉnh
- Best practices

#### exercises/
- Bài tập thực hành
- File bài tập và file solution
- Test cases

#### projects/ (sẽ thêm sau)
- Dự án lớn hơn
- Áp dụng nhiều kiến thức

## Quy ước đặt tên

- **Files markdown**: `01-introduction.md`, `02-setup.md`
- **Code files**: `snake_case.py`, `camelCase.js`, `PascalCase.java`
- **Thư mục**: `lowercase-with-dashes`

## Nội dung đã hoàn thành

### ✅ Python
- [x] README với mục lục đầy đủ
- [x] 5 chương đầu tiên (Introduction → Operators)
- [x] Code examples trong basics/
- [x] Ví dụ thực tế (Calculator, Todo List)
- [x] Bài tập và solutions

### ✅ JavaScript
- [x] README với mục lục
- [x] Chương 1: Introduction
- [x] Code examples cơ bản
- [x] Ví dụ Todo App (HTML/CSS/JS)

### ✅ Java
- [x] README với mục lục
- [x] Chương 1: Introduction
- [x] Code examples cơ bản

### ✅ C++
- [x] README với mục lục
- [x] Chương 1: Introduction
- [x] Code examples cơ bản

### ✅ Các ngôn ngữ khác
- [x] README cho mỗi ngôn ngữ
- [ ] Nội dung chi tiết (sẽ bổ sung)

## Cách sử dụng

1. **Đọc README.md chính** để hiểu tổng quan
2. **Chọn ngôn ngữ** bạn muốn học
3. **Đọc README.md** của ngôn ngữ đó để xem mục lục
4. **Bắt đầu từ chương 1** và đọc tuần tự
5. **Thực hành** với code examples
6. **Làm bài tập** để củng cố kiến thức

## Đóng góp

Xem [CONTRIBUTING.md](./CONTRIBUTING.md) để biết cách đóng góp nội dung mới hoặc cải thiện nội dung hiện có.

## Lưu ý

- Mỗi ngôn ngữ là độc lập, có thể học riêng lẻ
- Nội dung được viết bằng tiếng Việt để dễ hiểu
- Code examples có comment giải thích
- Các chương được link với nhau để dễ điều hướng

---

**Cập nhật lần cuối**: 2024


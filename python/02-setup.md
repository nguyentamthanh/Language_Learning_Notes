# Chương 2: Cài đặt và Môi trường

## Cài đặt Python

### Windows

1. Tải Python từ [python.org](https://www.python.org/downloads/)
2. Chạy installer và **chọn "Add Python to PATH"**
3. Kiểm tra cài đặt:
```bash
python --version
```

### macOS

```bash
# Sử dụng Homebrew
brew install python3

# Hoặc tải từ python.org
```

### Linux

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install python3 python3-pip

# Fedora
sudo dnf install python3 python3-pip
```

## Kiểm tra cài đặt

```bash
# Kiểm tra version
python --version
# hoặc
python3 --version

# Kiểm tra pip (package manager)
pip --version
# hoặc
pip3 --version
```

## IDE và Editor

### 1. VS Code (Khuyến nghị)
- Tải từ [code.visualstudio.com](https://code.visualstudio.com/)
- Cài extension: Python

### 2. PyCharm
- Community Edition miễn phí
- Professional Edition (trả phí)

### 3. Jupyter Notebook
- Tốt cho Data Science
- Cài đặt: `pip install jupyter`

### 4. IDLE
- Đi kèm với Python
- Đơn giản cho người mới bắt đầu

## Virtual Environment

Tạo môi trường ảo để quản lý packages:

```bash
# Tạo virtual environment
python -m venv myenv

# Kích hoạt (Windows)
myenv\Scripts\activate

# Kích hoạt (macOS/Linux)
source myenv/bin/activate

# Deactivate
deactivate
```

## Cài đặt packages

```bash
# Cài package
pip install package_name

# Cài từ requirements.txt
pip install -r requirements.txt

# Liệt kê packages đã cài
pip list

# Tạo requirements.txt
pip freeze > requirements.txt
```

## Chạy Python code

### 1. Interactive Mode
```bash
python
>>> print("Hello")
Hello
>>> exit()
```

### 2. Chạy file
```bash
python script.py
```

### 3. Chạy module
```bash
python -m module_name
```

## Cấu trúc project Python

```
my_project/
├── src/
│   ├── __init__.py
│   └── main.py
├── tests/
│   └── test_main.py
├── requirements.txt
├── README.md
└── .gitignore
```

## Tóm tắt

- Python có thể cài đặt trên mọi hệ điều hành
- Sử dụng IDE phù hợp để code hiệu quả hơn
- Virtual environment giúp quản lý dependencies
- Pip là công cụ quản lý packages chính

**Tiếp theo**: [Chương 3: Cú pháp cơ bản](./03-basic-syntax.md)


# Chương 4: Biến và Kiểu dữ liệu

## Variables (Biến)

### Khai báo biến

```python
# Khai báo và gán giá trị
name = "Python"
age = 30
height = 1.75
is_student = True

# Gán lại giá trị
age = 31

# Xóa biến
del age
```

### Quy tắc đặt tên biến

- Bắt đầu bằng chữ cái hoặc dấu gạch dưới
- Chỉ chứa chữ cái, số và dấu gạch dưới
- Phân biệt chữ hoa/thường
- Không được là keyword

```python
# Hợp lệ
my_variable = 10
_variable = 20
variable123 = 30

# Không hợp lệ
# 123variable = 10  # Lỗi: bắt đầu bằng số
# my-variable = 10  # Lỗi: có dấu gạch ngang
# class = 10        # Lỗi: là keyword
```

## Data Types (Kiểu dữ liệu)

### 1. Numbers

```python
# Integer
age = 30
negative = -10
large = 1000000

# Float
price = 99.99
scientific = 1.5e3  # 1500.0

# Complex
complex_num = 3 + 4j

# Type checking
print(type(age))      # <class 'int'>
print(type(price))    # <class 'float'>
```

### 2. Strings

```python
# Single quotes
text1 = 'Hello'

# Double quotes
text2 = "World"

# Triple quotes (multi-line)
text3 = """
Multi-line
string
"""

# F-strings (Python 3.6+)
name = "Python"
greeting = f"Hello, {name}!"

# String methods
text = "  Hello World  "
print(text.strip())      # "Hello World"
print(text.upper())      # "  HELLO WORLD  "
print(text.lower())      # "  hello world  "
print(text.replace("World", "Python"))
```

### 3. Boolean

```python
is_true = True
is_false = False

# Boolean operations
result = True and False  # False
result = True or False   # True
result = not True        # False
```

### 4. None

```python
value = None
print(value)  # None
print(type(value))  # <class 'NoneType'>
```

## Type Conversion

```python
# String to Integer
str_num = "123"
int_num = int(str_num)

# Integer to String
num = 42
str_num = str(num)

# Float to Integer
float_num = 3.14
int_num = int(float_num)  # 3 (truncate)

# String to Float
str_float = "3.14"
float_num = float(str_float)

# Boolean conversion
bool(1)      # True
bool(0)      # False
bool("")     # False
bool("text") # True
```

## Type Checking

```python
# isinstance() - kiểm tra kiểu
age = 30
print(isinstance(age, int))      # True
print(isinstance(age, str))      # False

# type() - lấy kiểu
print(type(age))                 # <class 'int'>
print(type(age) == int)          # True
```

## Dynamic Typing

Python là dynamically typed - kiểu dữ liệu được xác định khi runtime:

```python
# Cùng một biến có thể chứa các kiểu khác nhau
variable = 10        # int
variable = "Hello"   # str
variable = [1, 2, 3] # list
```

## Type Hints (Python 3.5+)

```python
# Type hints (không bắt buộc nhưng khuyến nghị)
def add(a: int, b: int) -> int:
    return a + b

# Variables với type hints
name: str = "Python"
age: int = 30
```

## Ví dụ thực tế

```python
# User input và type conversion
user_input = input("Nhập tuổi của bạn: ")
age = int(user_input)

if age >= 18:
    print("Bạn đã trưởng thành")
else:
    print("Bạn chưa trưởng thành")

# Formatting numbers
price = 99.99
print(f"Giá: ${price:.2f}")  # Giá: $99.99

# String formatting
name = "John"
age = 25
print(f"{name} năm nay {age} tuổi")
```

## Tóm tắt

- Biến lưu trữ dữ liệu
- Python có nhiều kiểu dữ liệu: int, float, str, bool, None
- Có thể chuyển đổi giữa các kiểu
- Type hints giúp code rõ ràng hơn

**Tiếp theo**: [Chương 5: Toán tử](./05-operators.md)


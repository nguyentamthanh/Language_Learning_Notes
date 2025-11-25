# Chương 3: Cú pháp cơ bản

## Comments (Ghi chú)

```python
# Single line comment

"""
Multi-line comment
hoặc docstring
"""

'''
Cũng có thể dùng single quotes
'''
```

## Print Statement

```python
# In ra màn hình
print("Hello, World!")

# In nhiều giá trị
print("Xin chào", "Python", 2024)

# Format string
name = "Python"
print(f"Hello, {name}!")
print("Hello, {}!".format(name))
```

## Indentation (Thụt lề)

Python sử dụng indentation để xác định khối code:

```python
if True:
    print("Đúng")      # Thụt lề 4 spaces
    print("Vẫn trong if")
print("Ngoài if")      # Không thụt lề
```

**Quan trọng**: 
- Sử dụng 4 spaces (khuyến nghị) hoặc 1 tab
- Không trộn spaces và tabs
- Indentation phải nhất quán

## Naming Conventions

```python
# Variables: snake_case
my_variable = 10
user_name = "John"

# Constants: UPPER_SNAKE_CASE
MAX_SIZE = 100
PI = 3.14159

# Classes: PascalCase
class MyClass:
    pass

# Functions: snake_case
def my_function():
    pass

# Private: _leading_underscore
_private_var = "private"
```

## Keywords

Python có các từ khóa không thể dùng làm tên biến:

```python
# Các keywords quan trọng
and, as, assert, break, class, continue, def, del, elif, else, 
except, False, finally, for, from, global, if, import, in, is, 
lambda, None, nonlocal, not, or, pass, raise, return, True, 
try, while, with, yield
```

## Multi-line Statements

```python
# Implicit line continuation
total = 1 + 2 + 3 + \
        4 + 5 + 6

# Explicit line continuation với parentheses
total = (1 + 2 + 3 +
         4 + 5 + 6)

# Multiple statements trên một dòng
a = 1; b = 2; c = 3
```

## Docstrings

```python
def my_function():
    """
    Đây là docstring mô tả hàm.
    
    Args:
        None
        
    Returns:
        None
    """
    pass

# Truy cập docstring
print(my_function.__doc__)
```

## Code Examples

### Ví dụ 1: Basic Output
```python
print("=" * 50)
print("Chào mừng đến với Python!")
print("=" * 50)
```

### Ví dụ 2: Multiple Assignments
```python
# Gán nhiều biến cùng lúc
x, y, z = 1, 2, 3
a = b = c = 0

# Swap values
x, y = y, x
```

### Ví dụ 3: String Operations
```python
# Nối chuỗi
greeting = "Hello" + " " + "World"

# Nhân chuỗi
separator = "-" * 20

# Multi-line string
text = """
Dòng 1
Dòng 2
Dòng 3
"""
```

## Tóm tắt

- Comments giúp code dễ hiểu hơn
- Indentation là cú pháp bắt buộc trong Python
- Tuân thủ naming conventions
- Docstrings mô tả chức năng của code

**Tiếp theo**: [Chương 4: Biến và Kiểu dữ liệu](./04-variables-types.md)


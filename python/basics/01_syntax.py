"""
Python Basics - Cú pháp cơ bản
"""

# 1. Comments (Ghi chú)
# Single line comment
"""
Multi-line comment
hoặc docstring
"""

# 2. Print statement
print("Hello, World!")
print("Xin chào", "Python")

# 3. Variables (Biến)
name = "Python"
age = 30
height = 1.75
is_student = True

# 4. Data Types (Kiểu dữ liệu)
# String
text = "Hello"
# Integer
number = 42
# Float
decimal = 3.14
# Boolean
is_true = True
is_false = False
# None
nothing = None

# 5. Type checking
print(type(name))      # <class 'str'>
print(type(age))       # <class 'int'>
print(type(height))    # <class 'float'>
print(type(is_student)) # <class 'bool'>

# 6. Type conversion
str_number = "123"
int_number = int(str_number)
float_number = float(str_number)

# 7. String formatting
name = "Python"
print(f"Hello, {name}!")  # f-string (Python 3.6+)
print("Hello, {}!".format(name))  # .format()
print("Hello, %s!" % name)  # % formatting

# 8. Input from user
# user_input = input("Nhập tên của bạn: ")
# print(f"Xin chào, {user_input}!")


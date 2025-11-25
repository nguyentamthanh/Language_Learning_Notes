# Chương 5: Toán tử

## Arithmetic Operators (Toán tử số học)

```python
a = 10
b = 3

print(a + b)   # 13 - Cộng
print(a - b)   # 7  - Trừ
print(a * b)   # 30 - Nhân
print(a / b)   # 3.333... - Chia
print(a // b)  # 3  - Chia lấy phần nguyên
print(a % b)   # 1  - Chia lấy phần dư
print(a ** b)  # 1000 - Lũy thừa
```

## Assignment Operators (Toán tử gán)

```python
x = 10

x += 5   # x = x + 5  → 15
x -= 3   # x = x - 3  → 12
x *= 2   # x = x * 2  → 24
x /= 4   # x = x / 4  → 6.0
x //= 2  # x = x // 2 → 3
x %= 2   # x = x % 2  → 1
x **= 3  # x = x ** 3 → 1
```

## Comparison Operators (Toán tử so sánh)

```python
a = 10
b = 20

print(a == b)  # False - Bằng
print(a != b)  # True  - Khác
print(a < b)   # True  - Nhỏ hơn
print(a > b)   # False - Lớn hơn
print(a <= b)  # True  - Nhỏ hơn hoặc bằng
print(a >= b)  # False - Lớn hơn hoặc bằng

# Chaining comparisons
print(5 < 10 < 15)  # True
```

## Logical Operators (Toán tử logic)

```python
x = True
y = False

print(x and y)  # False - Và
print(x or y)   # True  - Hoặc
print(not x)    # False - Phủ định

# Short-circuit evaluation
def func():
    print("Called")
    return True

True or func()   # func() không được gọi
False and func() # func() không được gọi
```

## Identity Operators (Toán tử đồng nhất)

```python
a = [1, 2, 3]
b = [1, 2, 3]
c = a

print(a is b)     # False - Khác object
print(a is c)     # True  - Cùng object
print(a == b)     # True  - Giá trị giống nhau

print(a is not b) # True
```

## Membership Operators (Toán tử thành viên)

```python
numbers = [1, 2, 3, 4, 5]

print(3 in numbers)      # True
print(10 in numbers)     # False
print(3 not in numbers)  # False

# Với strings
text = "Hello"
print("H" in text)      # True
print("World" in text)  # False
```

## Bitwise Operators (Toán tử bit)

```python
a = 5   # 101 in binary
b = 3   # 011 in binary

print(a & b)   # 1  - AND
print(a | b)   # 7  - OR
print(a ^ b)   # 6  - XOR
print(~a)      # -6 - NOT
print(a << 1)  # 10 - Left shift
print(a >> 1)  # 2  - Right shift
```

## Operator Precedence (Thứ tự ưu tiên)

Thứ tự từ cao đến thấp:

1. `()` - Parentheses
2. `**` - Exponentiation
3. `*`, `/`, `//`, `%` - Multiplication, Division
4. `+`, `-` - Addition, Subtraction
5. `<`, `>`, `<=`, `>=`, `==`, `!=` - Comparison
6. `not` - Logical NOT
7. `and` - Logical AND
8. `or` - Logical OR

```python
result = 2 + 3 * 4      # 14 (không phải 20)
result = (2 + 3) * 4    # 20
result = 2 ** 3 ** 2    # 512 (right-associative)
```

## Ví dụ thực tế

```python
# Kiểm tra số chẵn/lẻ
number = 10
if number % 2 == 0:
    print("Số chẵn")
else:
    print("Số lẻ")

# Tính toán đơn giản
price = 100
discount = 0.1
final_price = price * (1 - discount)
print(f"Giá cuối: {final_price}")

# Kiểm tra trong danh sách
fruits = ["apple", "banana", "orange"]
if "apple" in fruits:
    print("Có táo trong danh sách")

# So sánh chuỗi
name1 = "Alice"
name2 = "Bob"
print(name1 < name2)  # True (so sánh theo alphabet)
```

## Tóm tắt

- Arithmetic: `+`, `-`, `*`, `/`, `//`, `%`, `**`
- Assignment: `=`, `+=`, `-=`, `*=`, `/=`, etc.
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical: `and`, `or`, `not`
- Identity: `is`, `is not`
- Membership: `in`, `not in`
- Bitwise: `&`, `|`, `^`, `~`, `<<`, `>>`

**Tiếp theo**: [Chương 6: Cấu trúc điều khiển](./06-control-flow.md)


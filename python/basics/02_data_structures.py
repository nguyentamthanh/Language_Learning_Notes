"""
Python Basics - Cấu trúc dữ liệu
"""

# 1. Lists (Danh sách)
fruits = ["apple", "banana", "orange"]
numbers = [1, 2, 3, 4, 5]
mixed = [1, "hello", 3.14, True]

# List operations
fruits.append("grape")        # Thêm phần tử
fruits.insert(1, "mango")     # Chèn tại vị trí
fruits.remove("banana")       # Xóa phần tử
fruits.pop()                  # Xóa phần tử cuối
fruits.sort()                 # Sắp xếp
fruits.reverse()              # Đảo ngược

# List slicing
print(fruits[0])              # Phần tử đầu tiên
print(fruits[-1])             # Phần tử cuối cùng
print(fruits[1:3])            # Từ vị trí 1 đến 3
print(fruits[:2])             # Từ đầu đến vị trí 2
print(fruits[2:])             # Từ vị trí 2 đến cuối

# 2. Tuples (Bộ dữ liệu) - Immutable
coordinates = (10, 20)
point = (x, y) = (5, 10)

# 3. Dictionaries (Từ điển)
person = {
    "name": "Python",
    "age": 30,
    "city": "San Francisco"
}

# Dictionary operations
person["email"] = "python@example.com"  # Thêm key-value
person.get("name")                       # Lấy giá trị
person.keys()                            # Tất cả keys
person.values()                          # Tất cả values
person.items()                           # Tất cả items

# 4. Sets (Tập hợp) - Unique elements
unique_numbers = {1, 2, 3, 4, 5}
unique_numbers.add(6)                    # Thêm phần tử
unique_numbers.remove(3)                 # Xóa phần tử
unique_numbers.discard(7)                # Xóa nếu tồn tại

# Set operations
set1 = {1, 2, 3}
set2 = {3, 4, 5}
union = set1 | set2                      # Hợp
intersection = set1 & set2               # Giao
difference = set1 - set2                 # Hiệu

# 5. List Comprehensions
squares = [x**2 for x in range(10)]
even_squares = [x**2 for x in range(10) if x % 2 == 0]

# 6. Dictionary Comprehensions
square_dict = {x: x**2 for x in range(5)}


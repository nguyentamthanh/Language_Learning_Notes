"""
Python Basics - Cấu trúc điều khiển
"""

# 1. If/Else statements
age = 18

if age >= 18:
    print("Bạn đã trưởng thành")
elif age >= 13:
    print("Bạn là thiếu niên")
else:
    print("Bạn là trẻ em")

# Ternary operator
status = "adult" if age >= 18 else "minor"

# 2. For loops
# Loop through list
fruits = ["apple", "banana", "orange"]
for fruit in fruits:
    print(fruit)

# Loop with index
for index, fruit in enumerate(fruits):
    print(f"{index}: {fruit}")

# Loop through range
for i in range(5):
    print(i)

for i in range(1, 10, 2):  # Start, end, step
    print(i)

# Loop through dictionary
person = {"name": "Python", "age": 30}
for key, value in person.items():
    print(f"{key}: {value}")

# 3. While loops
count = 0
while count < 5:
    print(count)
    count += 1

# 4. Break and Continue
for i in range(10):
    if i == 5:
        break      # Thoát khỏi loop
    if i % 2 == 0:
        continue   # Bỏ qua iteration hiện tại
    print(i)

# 5. Nested loops
for i in range(3):
    for j in range(3):
        print(f"({i}, {j})")

# 6. Try/Except (Error handling)
try:
    result = 10 / 0
except ZeroDivisionError:
    print("Không thể chia cho 0")
except Exception as e:
    print(f"Lỗi: {e}")
else:
    print("Không có lỗi")
finally:
    print("Luôn chạy")

# 7. Match/Case (Python 3.10+)
def get_day_type(day):
    match day:
        case "Monday" | "Tuesday" | "Wednesday" | "Thursday" | "Friday":
            return "Weekday"
        case "Saturday" | "Sunday":
            return "Weekend"
        case _:
            return "Unknown"


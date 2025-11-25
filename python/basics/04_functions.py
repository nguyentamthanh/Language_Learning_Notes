"""
Python Basics - Hàm (Functions)
"""

# 1. Basic function
def greet():
    print("Hello, World!")

greet()

# 2. Function with parameters
def greet_person(name):
    print(f"Hello, {name}!")

greet_person("Python")

# 3. Function with default parameters
def greet_with_default(name="Guest"):
    print(f"Hello, {name}!")

greet_with_default()
greet_with_default("Python")

# 4. Function with return value
def add(a, b):
    return a + b

result = add(5, 3)
print(result)

# 5. Function with multiple return values
def get_name_and_age():
    return "Python", 30

name, age = get_name_and_age()

# 6. Function with *args (variable arguments)
def sum_all(*args):
    total = 0
    for num in args:
        total += num
    return total

print(sum_all(1, 2, 3, 4, 5))

# 7. Function with **kwargs (keyword arguments)
def print_info(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")

print_info(name="Python", age=30, city="SF")

# 8. Lambda functions (Anonymous functions)
square = lambda x: x ** 2
print(square(5))

# Lambda với map
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, numbers))

# Lambda với filter
even_numbers = list(filter(lambda x: x % 2 == 0, numbers))

# 9. Recursive function
def factorial(n):
    if n == 0 or n == 1:
        return 1
    return n * factorial(n - 1)

print(factorial(5))

# 10. Function decorators
def my_decorator(func):
    def wrapper():
        print("Before function call")
        func()
        print("After function call")
    return wrapper

@my_decorator
def say_hello():
    print("Hello!")

say_hello()

# 11. Type hints (Python 3.5+)
def add_numbers(a: int, b: int) -> int:
    return a + b


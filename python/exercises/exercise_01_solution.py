"""
Bài tập 1: Giải pháp
"""

def sum_list(numbers):
    """Tính tổng của một danh sách số"""
    total = 0
    for num in numbers:
        total += num
    return total
    # Hoặc: return sum(numbers)

def find_max(numbers):
    """Tìm số lớn nhất trong danh sách"""
    if not numbers:
        return None
    max_num = numbers[0]
    for num in numbers:
        if num > max_num:
            max_num = num
    return max_num
    # Hoặc: return max(numbers)

def count_even(numbers):
    """Đếm số chẵn trong danh sách"""
    count = 0
    for num in numbers:
        if num % 2 == 0:
            count += 1
    return count
    # Hoặc: return len([n for n in numbers if n % 2 == 0])

def is_prime(n):
    """Kiểm tra số nguyên tố"""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

# Test cases
if __name__ == "__main__":
    assert sum_list([1, 2, 3, 4, 5]) == 15
    assert find_max([1, 5, 3, 9, 2]) == 9
    assert count_even([1, 2, 3, 4, 5, 6]) == 3
    assert is_prime(7) == True
    assert is_prime(10) == False
    
    print("Tất cả test cases đã pass!")


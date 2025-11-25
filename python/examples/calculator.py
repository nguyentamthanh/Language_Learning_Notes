"""
Ví dụ: Máy tính đơn giản
"""

def calculator():
    """Máy tính đơn giản với các phép toán cơ bản"""
    
    print("=== Máy tính đơn giản ===")
    print("1. Cộng")
    print("2. Trừ")
    print("3. Nhân")
    print("4. Chia")
    print("5. Thoát")
    
    while True:
        choice = input("\nChọn phép toán (1-5): ")
        
        if choice == '5':
            print("Tạm biệt!")
            break
        
        if choice in ['1', '2', '3', '4']:
            try:
                num1 = float(input("Nhập số thứ nhất: "))
                num2 = float(input("Nhập số thứ hai: "))
                
                if choice == '1':
                    result = num1 + num2
                    print(f"Kết quả: {num1} + {num2} = {result}")
                elif choice == '2':
                    result = num1 - num2
                    print(f"Kết quả: {num1} - {num2} = {result}")
                elif choice == '3':
                    result = num1 * num2
                    print(f"Kết quả: {num1} * {num2} = {result}")
                elif choice == '4':
                    if num2 == 0:
                        print("Lỗi: Không thể chia cho 0!")
                    else:
                        result = num1 / num2
                        print(f"Kết quả: {num1} / {num2} = {result}")
            except ValueError:
                print("Lỗi: Vui lòng nhập số hợp lệ!")
        else:
            print("Lựa chọn không hợp lệ!")

if __name__ == "__main__":
    calculator()


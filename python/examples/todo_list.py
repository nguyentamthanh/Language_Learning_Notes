"""
Ví dụ: Ứng dụng Todo List
"""

class TodoList:
    """Lớp quản lý danh sách công việc"""
    
    def __init__(self):
        self.tasks = []
    
    def add_task(self, task):
        """Thêm công việc mới"""
        self.tasks.append({"task": task, "completed": False})
        print(f"Đã thêm: {task}")
    
    def remove_task(self, index):
        """Xóa công việc"""
        if 0 <= index < len(self.tasks):
            removed = self.tasks.pop(index)
            print(f"Đã xóa: {removed['task']}")
        else:
            print("Vị trí không hợp lệ!")
    
    def complete_task(self, index):
        """Đánh dấu công việc đã hoàn thành"""
        if 0 <= index < len(self.tasks):
            self.tasks[index]["completed"] = True
            print(f"Đã hoàn thành: {self.tasks[index]['task']}")
        else:
            print("Vị trí không hợp lệ!")
    
    def show_tasks(self):
        """Hiển thị danh sách công việc"""
        if not self.tasks:
            print("Danh sách trống!")
            return
        
        print("\n=== Danh sách công việc ===")
        for i, task in enumerate(self.tasks):
            status = "✓" if task["completed"] else "○"
            print(f"{i}. [{status}] {task['task']}")
        print()

def main():
    todo = TodoList()
    
    while True:
        print("\n1. Thêm công việc")
        print("2. Xóa công việc")
        print("3. Đánh dấu hoàn thành")
        print("4. Hiển thị danh sách")
        print("5. Thoát")
        
        choice = input("\nChọn hành động (1-5): ")
        
        if choice == '1':
            task = input("Nhập công việc: ")
            todo.add_task(task)
        elif choice == '2':
            todo.show_tasks()
            try:
                index = int(input("Nhập số thứ tự cần xóa: "))
                todo.remove_task(index)
            except ValueError:
                print("Vui lòng nhập số hợp lệ!")
        elif choice == '3':
            todo.show_tasks()
            try:
                index = int(input("Nhập số thứ tự đã hoàn thành: "))
                todo.complete_task(index)
            except ValueError:
                print("Vui lòng nhập số hợp lệ!")
        elif choice == '4':
            todo.show_tasks()
        elif choice == '5':
            print("Tạm biệt!")
            break
        else:
            print("Lựa chọn không hợp lệ!")

if __name__ == "__main__":
    main()


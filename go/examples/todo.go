package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Text      string
	Completed bool
}

type TodoList struct {
	tasks []Task
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (tl *TodoList) AddTask(text string) {
	task := Task{
		ID:        tl.nextID,
		Text:      text,
		Completed: false,
	}
	tl.tasks = append(tl.tasks, task)
	tl.nextID++
	fmt.Printf("Đã thêm: %s\n", text)
}

func (tl *TodoList) RemoveTask(id int) {
	for i, task := range tl.tasks {
		if task.ID == id {
			tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
			fmt.Printf("Đã xóa: %s\n", task.Text)
			return
		}
	}
	fmt.Println("Không tìm thấy task!")
}

func (tl *TodoList) CompleteTask(id int) {
	for i := range tl.tasks {
		if tl.tasks[i].ID == id {
			tl.tasks[i].Completed = true
			fmt.Printf("Đã hoàn thành: %s\n", tl.tasks[i].Text)
			return
		}
	}
	fmt.Println("Không tìm thấy task!")
}

func (tl *TodoList) ShowTasks() {
	if len(tl.tasks) == 0 {
		fmt.Println("Danh sách trống!")
		return
	}
	
	fmt.Println("\n=== Danh sách công việc ===")
	for _, task := range tl.tasks {
		status := "○"
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
	}
	fmt.Println()
}

func main() {
	todo := NewTodoList()
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\n1. Thêm công việc")
		fmt.Println("2. Xóa công việc")
		fmt.Println("3. Đánh dấu hoàn thành")
		fmt.Println("4. Hiển thị danh sách")
		fmt.Println("5. Thoát")
		
		fmt.Print("\nChọn hành động (1-5): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			fmt.Print("Nhập công việc: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			todo.AddTask(text)
			
		case "2":
			todo.ShowTasks()
			fmt.Print("Nhập ID cần xóa: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Vui lòng nhập số hợp lệ!")
				continue
			}
			todo.RemoveTask(id)
			
		case "3":
			todo.ShowTasks()
			fmt.Print("Nhập ID đã hoàn thành: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Vui lòng nhập số hợp lệ!")
				continue
			}
			todo.CompleteTask(id)
			
		case "4":
			todo.ShowTasks()
			
		case "5":
			fmt.Println("Tạm biệt!")
			return
			
		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
}


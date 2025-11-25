use std::io;

struct Task {
    id: usize,
    text: String,
    completed: bool,
}

struct TodoList {
    tasks: Vec<Task>,
    next_id: usize,
}

impl TodoList {
    fn new() -> TodoList {
        TodoList {
            tasks: Vec::new(),
            next_id: 1,
        }
    }
    
    fn add_task(&mut self, text: String) {
        let task = Task {
            id: self.next_id,
            text,
            completed: false,
        };
        self.tasks.push(task);
        self.next_id += 1;
        println!("Đã thêm công việc!");
    }
    
    fn remove_task(&mut self, id: usize) {
        if let Some(pos) = self.tasks.iter().position(|t| t.id == id) {
            let task = self.tasks.remove(pos);
            println!("Đã xóa: {}", task.text);
        } else {
            println!("Không tìm thấy task với ID: {}", id);
        }
    }
    
    fn complete_task(&mut self, id: usize) {
        if let Some(task) = self.tasks.iter_mut().find(|t| t.id == id) {
            task.completed = true;
            println!("Đã hoàn thành: {}", task.text);
        } else {
            println!("Không tìm thấy task với ID: {}", id);
        }
    }
    
    fn show_tasks(&self) {
        if self.tasks.is_empty() {
            println!("Danh sách trống!");
            return;
        }
        
        println!("\n=== Danh sách công việc ===");
        for task in &self.tasks {
            let status = if task.completed { "✓" } else { "○" };
            println!("{}. [{}] {}", task.id, status, task.text);
        }
        println!();
    }
}

fn main() {
    let mut todo = TodoList::new();
    
    loop {
        println!("\n1. Thêm công việc");
        println!("2. Xóa công việc");
        println!("3. Đánh dấu hoàn thành");
        println!("4. Hiển thị danh sách");
        println!("5. Thoát");
        
        print!("\nChọn hành động (1-5): ");
        io::Write::flush(&mut io::stdout()).unwrap();
        
        let mut choice = String::new();
        io::stdin()
            .read_line(&mut choice)
            .expect("Failed to read line");
        
        match choice.trim() {
            "1" => {
                print!("Nhập công việc: ");
                io::Write::flush(&mut io::stdout()).unwrap();
                let mut text = String::new();
                io::stdin()
                    .read_line(&mut text)
                    .expect("Failed to read line");
                todo.add_task(text.trim().to_string());
            }
            "2" => {
                todo.show_tasks();
                print!("Nhập ID cần xóa: ");
                io::Write::flush(&mut io::stdout()).unwrap();
                let mut id_str = String::new();
                io::stdin()
                    .read_line(&mut id_str)
                    .expect("Failed to read line");
                if let Ok(id) = id_str.trim().parse() {
                    todo.remove_task(id);
                } else {
                    println!("Vui lòng nhập số hợp lệ!");
                }
            }
            "3" => {
                todo.show_tasks();
                print!("Nhập ID đã hoàn thành: ");
                io::Write::flush(&mut io::stdout()).unwrap();
                let mut id_str = String::new();
                io::stdin()
                    .read_line(&mut id_str)
                    .expect("Failed to read line");
                if let Ok(id) = id_str.trim().parse() {
                    todo.complete_task(id);
                } else {
                    println!("Vui lòng nhập số hợp lệ!");
                }
            }
            "4" => {
                todo.show_tasks();
            }
            "5" => {
                println!("Tạm biệt!");
                break;
            }
            _ => {
                println!("Lựa chọn không hợp lệ!");
            }
        }
    }
}


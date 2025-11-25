use std::io;

fn main() {
    loop {
        println!("\n=== Máy tính đơn giản ===");
        println!("1. Cộng");
        println!("2. Trừ");
        println!("3. Nhân");
        println!("4. Chia");
        println!("5. Thoát");
        
        print!("\nChọn phép toán (1-5): ");
        io::Write::flush(&mut io::stdout()).unwrap();
        
        let mut choice = String::new();
        io::stdin()
            .read_line(&mut choice)
            .expect("Failed to read line");
        
        let choice = choice.trim();
        
        if choice == "5" {
            println!("Tạm biệt!");
            break;
        }
        
        if choice == "1" || choice == "2" || choice == "3" || choice == "4" {
            print!("Nhập số thứ nhất: ");
            io::Write::flush(&mut io::stdout()).unwrap();
            let mut num1_str = String::new();
            io::stdin()
                .read_line(&mut num1_str)
                .expect("Failed to read line");
            
            print!("Nhập số thứ hai: ");
            io::Write::flush(&mut io::stdout()).unwrap();
            let mut num2_str = String::new();
            io::stdin()
                .read_line(&mut num2_str)
                .expect("Failed to read line");
            
            let num1: f64 = match num1_str.trim().parse() {
                Ok(num) => num,
                Err(_) => {
                    println!("Lỗi: Vui lòng nhập số hợp lệ!");
                    continue;
                }
            };
            
            let num2: f64 = match num2_str.trim().parse() {
                Ok(num) => num,
                Err(_) => {
                    println!("Lỗi: Vui lòng nhập số hợp lệ!");
                    continue;
                }
            };
            
            match choice {
                "1" => {
                    let result = num1 + num2;
                    println!("Kết quả: {} + {} = {}", num1, num2, result);
                }
                "2" => {
                    let result = num1 - num2;
                    println!("Kết quả: {} - {} = {}", num1, num2, result);
                }
                "3" => {
                    let result = num1 * num2;
                    println!("Kết quả: {} * {} = {}", num1, num2, result);
                }
                "4" => {
                    if num2 == 0.0 {
                        println!("Lỗi: Không thể chia cho 0!");
                    } else {
                        let result = num1 / num2;
                        println!("Kết quả: {} / {} = {:.2}", num1, num2, result);
                    }
                }
                _ => {}
            }
        } else {
            println!("Lựa chọn không hợp lệ!");
        }
    }
}


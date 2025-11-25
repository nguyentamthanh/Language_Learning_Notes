fn main() {
    // Ownership example
    let s1 = String::from("hello");
    let s2 = s1;  // s1 bị move vào s2
    // println!("{}", s1);  // Lỗi! s1 không còn valid
    println!("s2: {}", s2);
    
    // Copy types
    let x = 5;
    let y = x;  // Copy, không move
    println!("x: {}, y: {}", x, y);
    
    // Borrowing
    let s3 = String::from("hello");
    let len = calculate_length(&s3);
    println!("'{}' có độ dài {}", s3, len);
    
    // Mutable reference
    let mut s4 = String::from("hello");
    change(&mut s4);
    println!("s4: {}", s4);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}


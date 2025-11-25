fn main() {
    // Function call
    another_function(5);
    
    // Function với return value
    let x = five();
    println!("x: {}", x);
    
    // Expression vs statement
    let y = {
        let z = 3;
        z + 1  // Expression, không có semicolon
    };
    println!("y: {}", y);
    
    // Function với parameters
    print_labeled_measurement(5, 'h');
    
    // Multiple return values
    let (sum, product) = calculate(4, 5);
    println!("Sum: {}, Product: {}", sum, product);
}

fn another_function(x: i32) {
    println!("Another function với giá trị: {}", x);
}

fn five() -> i32 {
    5  // Expression, không có semicolon
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("Measurement: {}{}", value, unit_label);
}

fn calculate(a: i32, b: i32) -> (i32, i32) {
    (a + b, a * b)
}


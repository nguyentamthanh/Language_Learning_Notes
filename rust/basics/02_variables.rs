fn main() {
    // Immutable variable
    let x = 5;
    println!("x: {}", x);
    
    // Mutable variable
    let mut y = 10;
    y = 20;
    println!("y: {}", y);
    
    // Type annotation
    let z: i32 = 15;
    println!("z: {}", z);
    
    // Shadowing
    let spaces = "   ";
    let spaces = spaces.len();
    println!("spaces: {}", spaces);
    
    // Constants
    const MAX_POINTS: u32 = 100_000;
    println!("MAX_POINTS: {}", MAX_POINTS);
    
    // Tuples
    let tup = (500, 6.4, 1);
    let (a, b, c) = tup;
    println!("Tuple: ({}, {}, {})", a, b, c);
    println!("First element: {}", tup.0);
    
    // Arrays
    let arr = [1, 2, 3, 4, 5];
    println!("Array: {:?}", arr);
    println!("First element: {}", arr[0]);
    
    // Strings
    let s = String::from("Hello");
    println!("String: {}", s);
    
    let s2 = "World";
    println!("&str: {}", s2);
}


// TODO: Viết hàm tính tổng của một slice số nguyên
fn sum_slice(numbers: &[i32]) -> i32 {
    // Your code here
    0
}

// TODO: Viết hàm tìm số lớn nhất trong slice
fn find_max(numbers: &[i32]) -> Option<i32> {
    // Your code here
    None
}

// TODO: Viết hàm đếm số chẵn trong slice
fn count_even(numbers: &[i32]) -> usize {
    // Your code here
    0
}

// TODO: Viết hàm kiểm tra số nguyên tố
fn is_prime(n: u32) -> bool {
    // Your code here
    false
}

fn main() {
    let numbers = vec![1, 2, 3, 4, 5];
    
    // Test sum_slice
    if sum_slice(&numbers) == 15 {
        println!("✓ sum_slice passed");
    } else {
        println!("✗ sum_slice failed");
    }
    
    // Test find_max
    if find_max(&numbers) == Some(5) {
        println!("✓ find_max passed");
    } else {
        println!("✗ find_max failed");
    }
    
    // Test count_even
    if count_even(&numbers) == 2 {
        println!("✓ count_even passed");
    } else {
        println!("✗ count_even failed");
    }
    
    // Test is_prime
    if is_prime(7) && !is_prime(10) {
        println!("✓ is_prime passed");
    } else {
        println!("✗ is_prime failed");
    }
}


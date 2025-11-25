// Giải pháp: Tính tổng của một slice số nguyên
fn sum_slice(numbers: &[i32]) -> i32 {
    numbers.iter().sum()
}

// Giải pháp: Tìm số lớn nhất trong slice
fn find_max(numbers: &[i32]) -> Option<i32> {
    if numbers.is_empty() {
        return None;
    }
    Some(*numbers.iter().max().unwrap())
}

// Giải pháp: Đếm số chẵn trong slice
fn count_even(numbers: &[i32]) -> usize {
    numbers.iter().filter(|&&n| n % 2 == 0).count()
}

// Giải pháp: Kiểm tra số nguyên tố
fn is_prime(n: u32) -> bool {
    if n < 2 {
        return false;
    }
    if n == 2 {
        return true;
    }
    if n % 2 == 0 {
        return false;
    }
    
    let mut i = 3;
    while i * i <= n {
        if n % i == 0 {
            return false;
        }
        i += 2;
    }
    true
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


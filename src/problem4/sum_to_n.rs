// Using While Loop
fn sum_to_n_a(n: i64) -> i64 {
    let mut sum = 0;
    for i in 1..=n {
        sum += i;
    }
    sum
}
// Time complexity of O(n) since there are n loops, 
// performing a constant time operation in each

// Use sum of AP
fn sum_to_n_b(n: i64) -> i64 {
    (n * (n + 1)) / 2
}
// Sum is computed in O(1) time.

// Recursion
fn sum_to_n_c(n: i64) -> i64 {
    if n <= 0 {
        0
    } else {
        n + sum_to_n_c(n - 1)
    }
}
// There will be n function calls, each call performing a constant time
// operation, resulting in  Time Complexity of O(n)

fn main() {
    println!("{}", sum_to_n_a(100));
    println!("{}", sum_to_n_b(100));
    println!("{}", sum_to_n_c(100));
}
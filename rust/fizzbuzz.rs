use std::io::stdio::print;

fn is_divisible_by_three(num: int) -> bool {
  num % 3 == 0
}

fn is_divisible_by_five(num: int) -> bool {
  num % 5 == 0
}

fn main() {
  for num in range(0,100) {
    print(num.to_str());
    print(" ");
    if is_divisible_by_three(num) { print("Fizz"); }
    if is_divisible_by_five(num)  { print("Buzz"); }
    print("\n");
  }
}

#[test]
fn test_is_three_with_one() { assert!(!is_divisible_by_three(1)) }

#[test]
fn test_is_three_with_three() { assert!(is_divisible_by_three(3)) }

#[test]
fn test_is_five_with_one() { assert!(!is_divisible_by_five(1)) }

#[test]
fn test_is_five_with_five() { assert!(is_divisible_by_five(5)) }

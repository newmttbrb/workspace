//Exercise 1.3.  Define a procedure that takes three numbers as arguments and returns the sum of the squares of the two larger numbers. 

fn remove_smallest(x : int, y : int, z : int) -> (int, int) {
  if x <= y
  {
    if x <= z    { (y, z) }
    else         { (x, y) }
  }
  else if y <= z { (x, z) } 
  else           { (x, y) }
}

fn sum_squares_of_larger_two(x : int, y : int, z : int) -> int {
  let (a , b) = remove_smallest(x,y,z);
  a * a + b * b
}

#[test]
fn test_squares_larger1() {
  assert!(sum_squares_of_larger_two(1,2,3) == 2*2+3*3)
}

#[test]
fn test_squares_larger2() {
  assert!(sum_squares_of_larger_two(10,20,30) == 20*20+30*30)
}

#[test]
fn test_remove_smallest_123() {
  println!("123");
  match remove_smallest(1,2,3) {
    (2,3) => (),
    (3,2) => (), 
    _ => fail!("remove smallest failed")
  }
}

#[test]
fn test_remove_smallest_132() {
  match remove_smallest(1,3,2) {
    (2,3) => (),
    (3,2) => (), 
    _ => fail!("remove smallest failed")
  }
}

#[test]
fn test_remove_smallest_213() {
  match remove_smallest(2,1,3) {
    (2,3) => (),
    (3,2) => (), 
    _ => fail!("remove smallest failed")
  }
}

#[test]
fn test_remove_smallest_231() {
  match remove_smallest(2,3,1) {
    (2,3) => (),
    (3,2) => (), 
    _ => fail!("remove smallest failed")
  }
}

#[test]
fn test_remove_smallest_312() {
  match remove_smallest(3,1,2) {
    (2,3) => (),
    (3,2) => (), 
    _ => fail!("remove smallest failed")
  }
}

#[test]
fn test_remove_smallest_321() {
  match remove_smallest(3,2,1) {
    (2,3) => (),
    (3,2) => (), 
    _ => fail!("remove smallest failed")
  }
}

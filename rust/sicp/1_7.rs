/*** Exercise 1.7.  The good-enough? test used in computing square roots will not be very effective for finding the square roots of very small numbers. Also, in real computers, arithmetic operations are almost always performed with limited precision. This makes our test inadequate for very large numbers. Explain these statements, with examples showing how the test fails for small and large numbers. An alternative strategy for implementing good-enough? is to watch how guess changes from one iteration to the next and to stop when the change is a very small fraction of the guess. Design a square-root procedure that uses this kind of end test. Does this work better for small and large numbers? ***/

fn good_enough(guess : f32, last_guess : f32) -> bool {
  std::f32::abs(guess - last_guess) < 0.0000001
}

fn improve(guess : f32, x : f32) -> f32 {
  average(guess, x/guess)
}

fn average(x : f32, y : f32) -> f32 { (x + y) / 2.0 }

fn sqrt_iter(x : f32) -> f32 {
  let mut guess = 1.0;
  let mut last_guess = 0.0;
  while !good_enough(guess,last_guess) { last_guess = guess; guess = improve(guess,x); }
  guess
}

fn main() {
  println!("{}",sqrt_iter(9.0));
  println!("{}",sqrt_iter(137.0));
  let result = sqrt_iter(1000.0);
  println!("{}",result * result);
}

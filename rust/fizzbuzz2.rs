fn fizzp(num : int) -> ~str { if num % 3 == 0 { ~"Fizz" } else { ~"" } }
fn buzzp(num : int) -> ~str { if num % 5 == 0 { ~"Buzz" } else { ~"" } }

fn main() {
  for i in range(0,100) {
    println!("{:d} {:s}{:s}",i,fizzp(i),buzzp(i));
  }
}

#[test]
fn not_div_3() { assert!(fizzp(1) == ~"") }

#[test]
fn div_3() { assert!(fizzp(3) == ~"Fizz") }

#[test]
fn not_div_5() { assert!(buzzp(1) == ~"") }

#[test]
fn div_5() { assert!(buzzp(5) == ~"Buzz") }


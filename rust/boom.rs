fn main() {
  let mut s = ~"BOOM";
  s.shift_char();
  s.unshift_char('Z');
  println!("It go {:s}", s);
}


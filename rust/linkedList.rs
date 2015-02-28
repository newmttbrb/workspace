enum List<T> {
  Cons(T, ~List<T>),
  Nil
}

fn print_list<T: std::fmt::Show>(list: &List<T>) {
  match list {
    &Nil                    =>   println!("[]"),
    &Cons(ref x, ~ref next) => { println!("{}",x); print_list(next); }
  }
}

impl<T: Eq>Eq for List<T> {
  fn eq(&self, ys: &List<T>) -> bool {
    match(self,ys) {
      (&Nil, &Nil) => true,
      (&Cons(ref x, ~ref next_xs), &Cons(ref y, ~ref next_ys)) if x == y => next_xs == next_ys,
      _ => false
    }
  }
}

fn prepend<T>(xs: List<T>, value: T)  -> List<T> {
  Cons(value, ~xs)
}

fn main() {
  //let list = Cons(1, ~Cons(2, ~Cons(3, ~Nil)));
  let mut list = Nil;
  list = prepend(list, 1);
  list = prepend(list, 2); 
  list = prepend(list, 3); 

  let list2 = Cons(3, ~Cons(2, ~Cons(1, ~Nil)));

  assert!( list.eq(&list2));
  assert!(!list.ne(&list2));
  assert!( list ==  list2);
  assert!(!(list !=  list2));
  print_list(&list);
}

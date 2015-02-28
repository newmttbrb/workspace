#[feature(struct_variant)]

struct Point {
  x: f64,
  y: f64
}

enum Shape {
  Circle{center: Point, radius: f64},
  Rectangle{top_left: Point, bottom_right: Point}
}

fn area(sh: Shape) -> f64 {
  match sh {
    Circle(_,size) => std::f64::consts::PI * size * size,
    Rectangle(Point{x,y}, Point{x: x2, y: y2}) => (x2 - x) * (y2 - y)
  }
}

fn main() {
  let point = Point { x: 2.0, y: 2.0 };
  let circle = Circle( point, 4.0 );
  let a = area(circle);
  println!("{}",a);
}

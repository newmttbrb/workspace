void setup() {
  size(480, 120);
}

void draw() {
  if (mousePressed) {
    fill(128);
  } else {
    fill(255);
  }
  ellipse(mouseX, mouseY, 10, 10);
}

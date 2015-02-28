use std::io::timer::sleep;

fn main() {
  let actions = [("Captain America", "bashes", 20),
                 ("Black Widow", "slashes", 25),
                 ("Ironman", "throws cash at", 0),
                 ("Hulk", "SMASHES", 200)];

  let mut outcomes =
    actions.iter()
      .map(|&action| {
             let (hero, attack, damage) = action;
             format!("{:s} {:s} Red Skull for {:d} damage",
               hero, attack, damage)
           });

  for outcome in outcomes {
    spawn(proc() {
      sleep(500);
      println!("{:s}", outcome);
    }); 
  }
}

use std::time::Instant;

fn main() {
    let start = Instant::now();

    println!("Hello, world!");

    println!("{}", start.elapsed().as_nanos());
}

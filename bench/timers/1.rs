use std::env;
use std::time::{Duration, Instant};

fn main() {
    let args: Vec<_> = env::args().collect();
    let n: u32 = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        1000000000
    };

    let duration = Duration::new(0, n);
    let mut counter = 0;
    let start = Instant::now();

    while start.elapsed() < duration {
        counter += 1;
    }

    println!("{}", counter);
}

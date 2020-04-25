use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n: u32 = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        30000000
    };

    let start = Instant::now();

    for _ in 0..n {
        Instant::now();
    }

    println!("{}", start.elapsed().as_nanos());
}

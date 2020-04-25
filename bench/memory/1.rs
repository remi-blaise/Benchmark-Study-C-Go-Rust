use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n: usize = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        1000_000_000
    } / 4;

    let start = Instant::now();

    let mut v = Vec::with_capacity(n);

    for _ in 0..n {
        v.push(42);
    }

    println!("{}", start.elapsed().as_nanos());
}

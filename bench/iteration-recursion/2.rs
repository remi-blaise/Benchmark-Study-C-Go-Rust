use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };

    let start = Instant::now();

    let a = rec(n);

    println!("{}", a);

    println!("{}", start.elapsed().as_nanos());
}

fn rec(n: u64) -> u64 {
    if n <= 0 {
        n
    } else {
        rec(n - 1)
    }
}

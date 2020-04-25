use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut n: u64 = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        40
    };

    let start = Instant::now();

    n = fib(n);

    println!("{}", n);

    println!("{}", start.elapsed().as_nanos());
}

fn fib(n: u64) -> u64 {
    if n <= 1 {
        n
    } else {
        fib(n - 1) + fib(n - 2)
    }
}

use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut n: u64 = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        100
    };

    n = fib(n);

    println!("{}", n);
}

fn fib(n: u64) -> u64 {
    if n <= 1 {
        n
    } else {
        fib(n - 1) + fib(n - 2)
    }
}

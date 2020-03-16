use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let N = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };
    let n = rec(0, N);

    println!("{:?}", n);
}

fn rec(n: u64, N: u64) -> u64 {
    if N <= 0 {
        n
    } else {
        rec(n + 1, N - 1)
    }
}

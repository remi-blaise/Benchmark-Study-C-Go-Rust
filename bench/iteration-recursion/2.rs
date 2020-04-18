use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };
    let a = rec(n);

    println!("{}", a);
}

fn rec(n: u64) -> u64 {
    if n <= 0 {
        n
    } else {
        rec(n - 1)
    }
}

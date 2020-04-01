use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000_000_000
    } / 4;

    println!("{}", n);

    let mut v = Vec::with_capacity::<i32>(n);

    for _ in 0..n {
        v.push(42);
    }
}

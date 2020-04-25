use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n: usize = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        1000_000_000
    } / 4;

    let mut v = Vec::with_capacity(n);

    for _ in 0..n {
        v.push(42);
    }

    let start = Instant::now();

    let mut _tmp: i32 = 0;

    for i in 0..n {
        _tmp = v[i];
    }

    println!("{}", start.elapsed().as_nanos());
}

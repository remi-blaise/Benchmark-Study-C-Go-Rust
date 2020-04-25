use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let s: usize = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        1000000
    } / 4;
    let n: usize = if args.len() > 2 {
        args[2].parse().unwrap()
    } else {
        1000
    };

    let mut _v: Vec<i32>;

    for _ in 0..n {
        _v = Vec::with_capacity(s);
    }
}

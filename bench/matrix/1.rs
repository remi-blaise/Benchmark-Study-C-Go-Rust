use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let m = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000
    };
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000
    };

    let start = Instant::now();

    let mut mat: Vec<Vec<i32>> = Vec::with_capacity(m);

    for _ in 0..m {
        mat.push(Vec::with_capacity(n));
    }

    println!("{}", start.elapsed().as_nanos());
}

use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let s = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        100000000
    };

    let start = Instant::now();

    let mut vec = Vec::with_capacity(s);

    for _ in 0..s {
        vec.push(42);
    }

    println!("{}", start.elapsed().as_nanos());
}

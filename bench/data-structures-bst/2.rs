use std::env;
use std::collections::BTreeSet;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let s = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };

    let mut bt = BTreeSet::new();

    for _ in 0..s {
        bt.insert(rand::random::<i32>());
    }

    let start = Instant::now();

    for _ in 0..n {
        bt.remove(&rand::random::<i32>());
    }

    println!("{}", start.elapsed().as_nanos());
}

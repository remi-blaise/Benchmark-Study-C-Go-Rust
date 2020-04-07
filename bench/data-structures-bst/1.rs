use std::env;
use std::collections::BTreeMap;

fn main() {
    let args: Vec<_> = env::args().collect();
    let s = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };

    let mut bt = BTreeMap::new();

    for _ in 0..s {
        bt.insert(rand::random::<i32>(), 0);
    }
}

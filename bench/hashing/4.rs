use std::env;
use std::collections::HashMap;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        10000000
    };

    let mut map = HashMap::new();

    for _ in 0..n {
        map.insert(rand::random::<i32>(), 0);
    }
}

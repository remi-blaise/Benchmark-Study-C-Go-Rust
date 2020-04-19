use std::env;
use std::collections::HashMap;
use std::hash::BuildHasherDefault;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };

    let mut map = HashMap::with_hasher(BuildHasherDefault::<hashers::fx_hash::FxHasher>::default());

    for _ in 0..n {
        map.insert(rand::random::<i32>(), 0);
    }
}

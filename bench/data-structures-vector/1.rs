use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let s = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };

    let mut vec = Vec::with_capacity(s);

    for _ in 0..s {
        vec.push(42);
    }
}
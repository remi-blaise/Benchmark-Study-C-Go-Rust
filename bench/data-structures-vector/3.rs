use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut s = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };
    let n = if args.len() > 2 {
        args[2].parse::<u64>().unwrap()
    } else {
        1000000
    };

    let mut vec = Vec::with_capacity(s);

    for _ in 0..s {
        vec.push(42);
    }

    for _ in 0..n {
        vec.remove(rand::random::<usize>() % s);
        s -= 1;
    }
}

use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut s = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        100000000
    };
    let n = if args.len() > 2 {
        args[2].parse::<u64>().unwrap()
    } else {
        10000000
    };

    let mut vec = Vec::with_capacity(s);

    for _ in 0..s {
        vec.push(42);
    }

    let mut _sum = 0;

    for _ in 0..n {
        _sum += vec[rand::random::<usize>() % s];

        s -= 1;
    }
}

use std::env;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut n = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };

    let start = Instant::now();

    while n > 0 {
        n -= 1;
    }

    print!("{}", n);

    println!("{}", start.elapsed().as_nanos());
}

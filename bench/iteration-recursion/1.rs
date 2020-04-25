use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut n = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };

    while n > 0 {
        n -= 1;
    }

    print!("{}", n);
}

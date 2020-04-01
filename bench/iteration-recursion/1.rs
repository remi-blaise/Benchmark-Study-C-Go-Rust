use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut N = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };

    while N > 0 {
        N -= 1;
    }

    println!("{}", N);
}

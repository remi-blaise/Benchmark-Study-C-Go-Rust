use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<u64>().unwrap()
    } else {
        1000
    };
    let a = rec(0, n);

    println!("{:?}", a);
}

fn rec(a: u64, b: u64) -> u64 {
    if b <= 0 {
        a
    } else {
        rec(a + 1, b - 1)
    }
}

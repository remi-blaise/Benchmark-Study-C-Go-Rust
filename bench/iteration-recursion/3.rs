use std::env;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n: usize = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        40
    };

    let mut x: i32 = 0;
    let mut y: i32 = 1;
    let mut z: i32;

    for _ in 0..n {
        z = x + y;
        x = y;
        y = z;
    }

    println!("{}", x);
}

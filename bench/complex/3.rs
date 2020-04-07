use std::env;
use num::complex::Complex;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000000
    };

    for _ in 0..n {
        let complex1 = Complex::<f32>::new(rand::random(), rand::random());
        let _ = complex1.norm();
    }
}

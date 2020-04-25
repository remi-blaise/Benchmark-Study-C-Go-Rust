use std::env;
use num::complex::Complex;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        50000000
    };

    let start = Instant::now();

    for _ in 0..n {
        let complex1: Complex<i32> = Complex::new(rand::random(), rand::random());
        let complex2: Complex<i32> = Complex::new(rand::random(), rand::random());
        let _ = complex1 + complex2;
        let _ = complex1 - complex2;
    }

    println!("{}", start.elapsed().as_nanos());
}

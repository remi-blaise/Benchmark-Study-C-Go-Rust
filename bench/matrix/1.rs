use std::env;
use simple_matrix::Matrix;

fn main() {
    let args: Vec<_> = env::args().collect();
    let m = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000
    };
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        1000
    };

    let _mat: Matrix<i32> = Matrix::new(m, n);
}

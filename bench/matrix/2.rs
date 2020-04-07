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

    let mat1: Matrix<i32> = Matrix::new(m, n);
    let mat2: Matrix<i32> = Matrix::new(n, m);

    for i in 0..m {
        for j in 0..n {
            mat1.set(i, j, rand::random::<i32>());
            mat2.set(j, i, rand::random::<i32>());
        }
    };

    mat1 * mat2
}

use std::env;
use std::num::Wrapping;

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

    // Create random matrixes

    let mut a: Vec<Vec<Wrapping<i32>>> = Vec::with_capacity(m);
    let mut b: Vec<Vec<Wrapping<i32>>> = Vec::with_capacity(m);

    for i in 0..m {
        a.push(Vec::with_capacity(n));
        b.push(Vec::with_capacity(n));

        for _ in 0..n {
            a[i].push(Wrapping(rand::random::<i32>()));
            b[i].push(Wrapping(rand::random::<i32>()));
        }
    }

    // Multiply matrixes

    let mut product: Vec<Vec<Wrapping<i32>>> = Vec::with_capacity(m);

    for i in 0..m {
        product.push(Vec::with_capacity(m));

        for j in 0..m {
            product[i].push(Wrapping(0));

            for k in 0..n {
                product[i][j] = product[i][j] + a[i][k] * b[k][j];
            }
        }
    }
}

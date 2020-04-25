use std::env;
use nix::unistd::{fork, ForkResult};
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        10000
    };

    let start = Instant::now();

    for _ in 0..n {
        match fork().unwrap() {
            ForkResult::Parent { .. } => {}
            ForkResult::Child => {
                std::process::exit(0);
            },
        }
    }

    println!("{}", start.elapsed().as_nanos());
}

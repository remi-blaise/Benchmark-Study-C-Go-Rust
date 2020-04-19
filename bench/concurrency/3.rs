use std::env;
use nix::unistd::{fork, ForkResult, Pid};

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        10000
    };

    let mut children: Vec<Pid> = Vec::new();

    for _ in 0..n {
        match fork().unwrap() {
            ForkResult::Parent { child } => {
                children.push(child)
            }
            ForkResult::Child => {
                std::process::exit(1);
            },
        }
    }
}

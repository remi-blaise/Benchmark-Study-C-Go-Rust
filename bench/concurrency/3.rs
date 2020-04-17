use std::env;
use nix::unistd::{fork, ForkResult, Pid};
use nix::sys::wait::{waitpid, WaitStatus};

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        10000
    };

    let mut children: Vec<Pid> = Vec::new();

    for i in 0..n {
        match fork().unwrap() {
            ForkResult::Parent { child } => {
                children.push(child)
            }
            ForkResult::Child => {
                println!("{}", i);
                std::process::exit(1);
            },
        }
    }

    for child in children {
        loop {
            if let WaitStatus::Exited(_pid, _status) = waitpid(child, None).unwrap() {
                break;
            } else {
                unimplemented!();
            }
        }
    }
}

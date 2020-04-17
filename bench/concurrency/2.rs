use std::env;
use std::thread;
use std::sync::{Arc, Mutex};

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        10000
    };

    let mut threads: Vec<thread::JoinHandle<_>> = Vec::new();
    let x = Arc::new(Mutex::new(0));

    for _ in 0..n {
        let x = Arc::clone(&x);
        threads.push(thread::spawn(move || {
            let mut x = x.lock().unwrap();
            *x += 1;
        }));
    }

    for thread in threads {
        thread.join().unwrap();
    }

    println!("{}", x.lock().unwrap());
}

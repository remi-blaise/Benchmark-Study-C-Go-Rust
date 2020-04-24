use std::env;
use std::thread;

static mut X: i32 = 0;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        25000
    };

    let mut threads: Vec<thread::JoinHandle<_>> = Vec::new();

    for _ in 0..n {
        threads.push(thread::spawn(move || {
            unsafe {
                X += 1;
            }
        }));
    }

    for thread in threads {
        thread.join().unwrap();
    }

    unsafe {
        println!("{}", X);
    }
}

use std::env;
use std::net::TcpStream;
use std::io::prelude::*;
use std::io::BufReader;
use std::time::Instant;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        100000
    };

    let stream = TcpStream::connect("127.0.0.1:8080").unwrap();
    let mut buffered_stream = BufReader::new(&stream);
    let mut line;

    let start = Instant::now();

    for _ in 0..n {
        (&stream).write_all(b"Hello from client\n").unwrap();
        line = String::new();
        buffered_stream.read_line(&mut line).unwrap();
        // print!("Reply: {}", line);
    }

    println!("{}", start.elapsed().as_nanos());
}

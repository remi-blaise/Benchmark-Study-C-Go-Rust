use std::env;
use std::net::TcpStream;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    let args: Vec<_> = env::args().collect();
    let n = if args.len() > 1 {
        args[1].parse::<usize>().unwrap()
    } else {
        100
    };

    let mut stream = TcpStream::connect("127.0.0.1:8080").unwrap();

    for _ in 0..n {
        stream.write_all(b"Hello from client\n").unwrap();
        let mut line = String::new();
        BufReader::new(&stream).read_line(&mut line).unwrap();
        print!("Reply: {}", line);
    }
}

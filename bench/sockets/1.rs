use std::net::TcpListener;
use std::io::prelude::*;
use std::io::BufReader;
use std::time::Instant;

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    let stream = listener.accept().unwrap().0;
    let mut buffered_stream = BufReader::new(&stream);
    let mut line = String::new();

    let start = Instant::now();

    while buffered_stream.read_line(&mut line).unwrap() != 0 {
        // print!("Message received: {}", line);
        (&stream).write_all(b"Hello from server\n").unwrap();
        line = String::new();
    }

    println!("{}", start.elapsed().as_nanos());
}

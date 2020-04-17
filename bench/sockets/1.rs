use std::net::TcpListener;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    let mut stream = listener.accept().unwrap().0;
    let mut line = String::new();

    while BufReader::new(&stream).read_line(&mut line).unwrap() != 0 {
        print!("Message received: {}", line);
        stream.write_all(b"Hello from server\n").unwrap();
        line = String::new();
    }
}

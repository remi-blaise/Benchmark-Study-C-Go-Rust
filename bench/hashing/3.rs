use std::fs::File;
use std::io::{BufRead, BufReader};
use crypto::digest::Digest;
use crypto::sha2::Sha512;

const FILENAME: &str = "asset/1000passwords";

fn main() {
    let file = File::open(FILENAME).unwrap();
    let reader = BufReader::new(file);

    for (_index, line) in reader.lines().enumerate() {
        let line = line.unwrap();

        let mut hasher = Sha512::new();
        hasher.input_str(&line);
        let _hash = hasher.result_str();

        #[cfg(debug_assertions)]
        println!("{}", _hash);
    }
}

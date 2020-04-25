use std::fs::File;
use std::io::prelude::*;
use crypto::digest::Digest;
use crypto::sha1::Sha1;
use std::time::Instant;

const FILENAME: &str = "asset/lorem100mb";

fn main() {
    let mut file = File::open(FILENAME).unwrap();
    let mut contents = Vec::new();
    file.read_to_end(&mut contents).unwrap();

    let start = Instant::now();

    let mut hasher = Sha1::new();
    hasher.input(&contents);
    let _hash = hasher.result_str();

    #[cfg(debug_assertions)]
    println!("{}", _hash);

    println!("{}", start.elapsed().as_nanos());
}

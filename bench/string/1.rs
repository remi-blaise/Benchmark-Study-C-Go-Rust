use std::fs::File;
use std::io::prelude::*;
use std::time::Instant;

fn main() -> std::io::Result<()> {
    let mut file = File::open("asset/lorem100mb")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    let start = Instant::now();

    let split = contents.split(" ");

    println!("{}", start.elapsed().as_nanos());

    let _vec: Vec<&str> = split.collect();
    Ok(())
}

use std::fs::File;
use std::io::prelude::*;

fn main() -> std::io::Result<()> {
    let mut file = File::open("foo.txt")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    let split = contents.split(" ");
    let _vec: Vec<&str> = split.collect();
    Ok(())
}

use std::env;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
struct Data {
    a: i64,
    b: f64,
    c: String,
}

fn main() {
    let args: Vec<_> = env::args().collect();
    let n: usize = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        100000
    };

    for _ in 0..n {
        let data = Data { a: 13273828327, b: 382283.537749, c: String::from("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet luctus leo sed imperdiet. Morbi ut dolor eu arcu pretium bibendum. Donec eleifend arcu sit amet sodales ultrices. Nam quis diam vel mi hendrerit egestas quis in velit. Aliquam non vulputate magna. Cras et magna bibendum, facilisis magna et, rhoncus.") };

        let encoded: Vec<u8> = bincode::serialize(&data).unwrap();

        #[cfg(debug_assertions)]
        println!("encoded = {:?}", encoded);

        let _decoded: Data = bincode::deserialize(&encoded[..]).unwrap();

        #[cfg(debug_assertions)]
        println!("decoded = {:?}", _decoded);
    }
}

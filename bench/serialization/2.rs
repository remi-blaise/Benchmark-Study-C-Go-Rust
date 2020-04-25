use std::env;
use serde::{Serialize, Deserialize};
use std::time::Instant;

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

    let start = Instant::now();

    for _ in 0..n {
        let data = Data { a: 13273828327, b: 382283.537749, c: String::from("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet luctus leo sed imperdiet. Morbi ut dolor eu arcu pretium bibendum. Donec eleifend arcu sit amet sodales ultrices. Nam quis diam vel mi hendrerit egestas quis in velit. Aliquam non vulputate magna. Cras et magna bibendum, facilisis magna et, rhoncus.") };

        let serialized = serde_json::to_string(&data).unwrap();

        #[cfg(debug_assertions)]
        println!("serialized = {}", serialized);

        let _deserialized: Data = serde_json::from_str(&serialized).unwrap();

        #[cfg(debug_assertions)]
        println!("deserialized = {:?}", _deserialized);
    }

    println!("{}", start.elapsed().as_nanos());
}

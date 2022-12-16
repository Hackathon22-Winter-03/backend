pub mod markov;

use crate::markov::*;

fn main() {
    let markov = Markov::new(vec![
        Rule::new("woman", "W", false),
        Rule::new("man", "M", false),
        Rule::new("MW", "", false),
        Rule::new("WM", "", false),
    ]);

    println!("{:?}", markov.compute("manmanwomanwomanmanwomanwomanmanwomanmanmanwoman"));
}

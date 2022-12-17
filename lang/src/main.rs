// pub mod turing;
pub mod markov;

// use crate::turing::*;
use crate::markov::*;

fn main() {

    match Markov::new("woman:W\nman:M\nMW:\nWM:\n") {
        Ok(markov) => {
            println!("{:?}", markov.compute("manmanwomanwomanmanwomanwomanmanwomanmanmanwoman"));
        },
        Err(msg) => {},
    }
    // if let Ok(turing) = Turing::new("(S,_):(S,_,R)\n(S,0):(S,_,R)\n(S,1):(one,1,R)\n(one,1):(two,_,L)\n(two,1):(zero,_,R)\n(zero,_):(zero,_,R)\n(zero,1):(one,1,R)\n(zero,A):(_,A,R)\n(one,A):(_,A,R)") {
    //     println!("{:?}", turing.compute("_____________0111010101010101100A________________"));
    // }
    
}
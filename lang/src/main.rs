pub mod markov;

// extern crate libc;
// use std::ffi::{CStr, CString};

use crate::markov::*;

// #[no_mangle]
// pub extern "C" fn rustaceanize(code: *const libc::c_char, input: *const libc::c_char) -> *const libc::c_char {
//     let cstr_code = unsafe { CStr::from_ptr(code) };
//     let cstr_input = unsafe { CStr::from_ptr(input) };
//     let str_code = cstr_code.to_str();
//     let str_input = cstr_input.to_str();
// }

fn main() {
    // let markov = Markov::new(vec![
    //     Rule::new("woman", "W", false),
    //     Rule::new("man", "M", false),
    //     Rule::new("MW", "", false),
    //     Rule::new("WM", "", false),
    // ]);
    match Markov::new("woman:W\nman:M\nMW:\nWM:\n") {
        Ok(markov) => {
            println!("{:?}", markov.compute("manmanwomanwomanmanwomanwomanmanwomanmanmanwoman"));
        },
        Err(msg) => {},
    }
    
}
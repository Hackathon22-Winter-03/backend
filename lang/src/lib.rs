pub mod markov;

extern crate libc;
use std::ffi::{CStr, CString};

use crate::markov::*;

#[no_mangle]
pub extern "C" fn simulate_markov(code: *const libc::c_char, input: *const libc::c_char) -> *const libc::c_char {
    let cstr_code = unsafe { CStr::from_ptr(code) };
    let cstr_input = unsafe { CStr::from_ptr(input) };
    let str_code = cstr_code.to_str().unwrap();
    let str_input = cstr_input.to_str().unwrap();

    let str_output;
    match Markov::new(str_code) {
        Ok(markov) => {
            str_output = markov.compute(str_input);
            CString::new(str_output.into_iter().collect::<String>()).unwrap().into_raw()
        },
        Err(msg) => {
            CString::new(msg.to_string()).unwrap().into_raw()
        },
    }
}

// fn main() {
//     // let markov = Markov::new(vec![
//     //     Rule::new("woman", "W", false),
//     //     Rule::new("man", "M", false),
//     //     Rule::new("MW", "", false),
//     //     Rule::new("WM", "", false),
//     // ]);
//     match Markov::new("woman:W\nman:M\nMW:\nWM:\n") {
//         Ok(markov) => {
//             println!("{:?}", markov.compute("manmanwomanwomanmanwomanwomanmanwomanmanmanwoman"));
//         },
//         Err(msg) => {},
//     }
    
// }
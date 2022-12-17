pub mod turing;
pub mod markov;

extern crate libc;
use std::ffi::{CStr, CString};

use crate::turing::*;
use crate::markov::*;

#[no_mangle]
pub extern "C" fn simulate(code: *const libc::c_char, input: *const libc::c_char) -> *const libc::c_char {
    let cstr_code = unsafe { CStr::from_ptr(code) };
    let cstr_input = unsafe { CStr::from_ptr(input) };
    let str_code = cstr_code.to_str().unwrap();
    let str_input = cstr_input.to_str().unwrap();

    let str_output;
    // if str_model == "markov" {
        match Markov::new(str_code) {
            Ok(mut markov) => {
                markov.set_text(str_input);
                str_output = markov.run();
                return CString::new(str_output.into_iter().collect::<String>()).unwrap().into_raw()
            },
            Err(msg) => {
                return CString::new(msg.to_string()).unwrap().into_raw()
            },
        }
    // } else if str_model == "turing" {
    //     match Turing::new(str_code) {
    //         Ok(mut turing) => {
    //             turing.set_tape(str_input);
    //             str_output = turing.run();
    //             return CString::new(str_output.into_iter().collect::<String>()).unwrap().into_raw()
    //         },
    //         Err(msg) => {
    //             return CString::new(msg.to_string()).unwrap().into_raw()
    //         },
    //     }
    // } else {
    //     CString::new(String::from(""));
    // }
}

#[no_mangle]
pub extern "C" fn step_execute(code: *const libc::c_char, input: *const libc::c_char, model: *const libc::c_char) -> *const libc::c_char {
    let cstr_code = unsafe { CStr::from_ptr(code) };
    let cstr_input = unsafe { CStr::from_ptr(input) };
    let cstr_model = unsafe { CStr::from_ptr(model) };

    let str_code = cstr_code.to_str().unwrap();
    let str_input = cstr_input.to_str().unwrap();
    let str_model = cstr_model.to_str().unwrap();

    let str_output;
    if str_model == "markov" {
        match Markov::new(str_code) {
            Ok(mut markov) => {
                markov.set_text(str_input);
                str_output = markov.step().0;
                return CString::new(str_output.into_iter().collect::<String>()).unwrap().into_raw()
            },
            Err(msg) => {
                return CString::new(msg.to_string()).unwrap().into_raw()
            },
        }
    } else if str_model == "turing" {
        match Turing::new(str_code) {
            Ok(mut turing) => {
                turing.set_tape(str_input);
                str_output = turing.step();
                return CString::new(str_output.into_iter().collect::<String>()).unwrap().into_raw()
            },
            Err(msg) => {
                return CString::new(msg.to_string()).unwrap().into_raw()
            },
        }
    } else {
        return CString::new(String::from("")).unwrap().into_raw();
    }
}

fn main() {
//     if let Ok(mut markov) = Markov::new("woman:W\nman:M\nMW:\nWM:\n") {
//         markov.set_text("manmanwomanwomanmanwomanwomanmanwomanmanmanwoman");
//         println!("{:?}", markov.run());
//     }
    // if let Ok(mut turing) = Turing::new("(S,_):(S,_,R)\n(S,0):(S,_,R)\n(S,1):(one,1,R)\n(one,0):(one,_,R)\n(one,1):(two,_,L)\n(two,1):(zero,_,R)\n(two,_):(two,_,L)\n(zero,_):(zero,_,R)\n(zero,0):(zero,_,R)\n(zero,1):(one,1,R)\n(zero,A):(_,A,R)\n(one,A):(_,A,R)") {
    //     turing.set_tape("_____0111010101010101100A___");
    //     println!("{:?}", turing.run());
    // }
}
pub mod markov;

extern crate libc;
use std::ffi::{CStr, CString};

use crate::markov::*;

#[no_mangle]
pub extern "C" fn simulate(
    code: *const libc::c_char,
    input: *const libc::c_char,
) -> *const libc::c_char {
    let cstr_code = unsafe { CStr::from_ptr(code) };
    let cstr_input = unsafe { CStr::from_ptr(input) };
    let str_code = cstr_code.to_str().unwrap();
    let str_input = cstr_input.to_str().unwrap();

    let str_output;
    match Markov::new(str_code) {
        Ok(mut markov) => {
            markov.set_text(str_input);
            str_output = markov.run();
            CString::new(str_output.into_iter().collect::<String>())
                .unwrap()
                .into_raw()
        }
        Err(msg) => CString::new(msg.to_string()).unwrap().into_raw(),
    }
}

#[no_mangle]
pub extern "C" fn step_execute(
    code: *const libc::c_char,
    input: *const libc::c_char,
    model: *const libc::c_char,
) -> *const libc::c_char {
    let cstr_code = unsafe { CStr::from_ptr(code) };
    let cstr_input = unsafe { CStr::from_ptr(input) };
    let cstr_model = unsafe { CStr::from_ptr(model) };

    let str_code = cstr_code.to_str().unwrap();
    let str_input = cstr_input.to_str().unwrap();
    let str_model = cstr_model.to_str().unwrap();

    println!("{:?} {:?} {:?}", str_code, str_input, str_model);

    match str_model {
        "markov" => match Markov::new(str_code) {
            Ok(mut markov) => {
                markov.set_text(str_input);
                let (str_output, is_terminated) = markov.step();
                println!("{:?} {:?}", str_output, is_terminated);
                if is_terminated {
                    return CString::new(str_output.into_iter().collect::<String>() + "T")
                        .unwrap()
                        .into_raw();
                } else {
                    return CString::new(str_output.into_iter().collect::<String>() + "F")
                        .unwrap()
                        .into_raw();
                }
            }
            Err(msg) => {
                return CString::new(msg.to_string()).unwrap().into_raw();
            }
        },
        _ => {}
    };

    return CString::new(String::from("")).unwrap().into_raw();
}

fn main() {}

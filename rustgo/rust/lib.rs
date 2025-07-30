#![allow(dead_code)]

use std::{mem::MaybeUninit, slice::from_raw_parts_mut, str::from_utf8_unchecked_mut};

use crate::tokenize::{Token, Tokenizer};

mod tokenize;

fn tokenize(s: &str) -> Vec<Token> {
    let mut tokenizer = Tokenizer::new_from_string(s.into());
    let tokens = tokenizer.all();
    return tokens;
}

#[cfg_attr(all(target_arch = "wasm32"), export_name = "tokenize")]
pub unsafe extern "C" fn tokenize_wasm(ptr: u32, len: u32) -> u64 {
    unsafe {
        let json = import_str(ptr, len);
        let tokens = tokenize(&json);
        return export_slice(&tokens);
    }
}

unsafe fn import_str(ptr: u32, len: u32) -> String {
    unsafe {
        let slice = from_raw_parts_mut(ptr as *mut u8, len as usize);
        let utf8 = from_utf8_unchecked_mut(slice);
        return String::from(utf8);
    }
}

unsafe fn export_slice(s: &[Token]) -> u64 {
    return ((s.as_ptr() as u64) << 32) + (s.len() / size_of::<Token>()) as u64;
}

#[cfg_attr(all(target_arch = "wasm32"), export_name = "allocate")]
pub extern "C" fn allocate_wasm(size: u32) -> *mut u8 {
    allocate(size as usize)
}

fn allocate(size: usize) -> *mut u8 {
    let vec: Vec<MaybeUninit<u8>> = vec![MaybeUninit::uninit(); size];
    Box::into_raw(vec.into_boxed_slice()) as *mut u8
}

#[cfg_attr(all(target_arch = "wasm32"), export_name = "deallocate")]
pub unsafe extern "C" fn deallocate_wasm(ptr: u32, size: u32) {
    unsafe {
        deallocate(ptr as *mut u8, size as usize);
    }
}

unsafe fn deallocate(ptr: *mut u8, size: usize) {
    unsafe {
        let _ = Vec::from_raw_parts(ptr, 0, size);
    }
}

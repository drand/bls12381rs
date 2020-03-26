use groupy::{CurveAffine, CurveProjective};
use libc;
use paired::bls12_381::{G1, G2};
use paired::HashToCurve;
use std::slice::from_raw_parts;

#[no_mangle]
pub extern "C" fn maptopointg2(
    msg: *const u8,
    msg_len: libc::size_t,
    dst: *mut u8,
    domain: *const u8,
    domain_len: libc::size_t,
) {
    let message = unsafe { from_raw_parts(msg, msg_len) };
    let csuite = unsafe { from_raw_parts(domain, domain_len) };
    let p = G2::hash_to_curve(&message, &csuite);
    let point_buff = p.into_affine().into_compressed();
    unsafe {
        std::ptr::copy(point_buff.as_ref().as_ptr(), dst, point_buff.as_ref().len());
    }
}

#[no_mangle]
pub extern "C" fn maptopointg1(
    msg: *const u8,
    msg_len: libc::size_t,
    dst: *mut u8,
    domain: *const u8,
    domain_len: libc::size_t,
) {
    let message = unsafe { from_raw_parts(msg, msg_len) };
    let csuite = unsafe { from_raw_parts(domain, domain_len) };
    let p = G1::hash_to_curve(&message, &csuite);
    let point_buff = p.into_affine().into_compressed();
    unsafe {
        std::ptr::copy(point_buff.as_ref().as_ptr(), dst, point_buff.as_ref().len());
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}

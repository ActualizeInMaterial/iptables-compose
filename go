#!/bin/bash
#
unset RUST_BACKTRACE
#workaround issue: 29293 https://github.com/rust-lang/rust/issues/29293
time cargo run -v -- ./example.yaml --reset



build:
	cd map-rs && (RUSTFLAGS="--print native-static-libs"  cargo build --release)
	cp map-rs/target/release/libkyberrs.a map-rs/lib/
	#go build map.go

.PHONY: build

test:
	go test .

run:
	go run .


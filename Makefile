build:
	cd map-rs && cargo build --release
	cp map-rs/target/release/libkyberrs.so map-rs/lib/
	go build -ldflags="-r map-rs/lib" map.go

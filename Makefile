build:
	cd map-rs && cargo build --release
	cp map-rs/target/release/libkyberrs.a map-rs/lib/
	#go build map.go

#test:
	#go get -t .
	#go build -ldflags="-r map-rs/lib" .


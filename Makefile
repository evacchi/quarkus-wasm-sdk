build: hello-world/main.go basic-auth/main.go
	tinygo build -target wasi -o hello-world/hello-world.wasm hello-world/main.go
	tinygo build -target wasi -o fortunes/fortunes.wasm fortunes/main.go
	tinygo build -target wasi -o basic-auth/basic-auth.wasm basic-auth/main.go

test: build
	extism call hello-world/hello-world.wasm request_headers \
		--input '{"headers":{"Content-Type":["application/json"]}}' --wasi
		
	extism call fortunes/fortunes.wasm request_headers \
		--input '{"headers":{"Content-Type":["application/json"]}}' --wasi

	extism call basic-auth/basic-auth.wasm request_headers \
		--input '{"headers":{"Authorization":["Basic YWRtaW46YWRtaW4="]}}' --wasi

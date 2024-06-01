build:
	tinygo build -target wasi -o plugin.wasm main.go

test:
	extism call plugin.wasm request_headers \
		--input '{"headers":{"Content-Type":["application/json"]}}' --wasi

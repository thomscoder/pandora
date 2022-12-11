compile:
	tinygo build -wasm-abi=generic -target=wasi -o main.wasm main.go
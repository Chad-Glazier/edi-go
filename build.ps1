$env:GOARCH = "wasm"
$env:GOOS = "js"
tinygo build -o .\target\edi.wasm .\main.go

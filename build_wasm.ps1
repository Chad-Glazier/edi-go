$goarch = $env:GOARCH
$goos = $env:GOOS
$env:GOARCH = "wasm"
$env:GOOS = "js"
tinygo build -o .\target\edi.wasm -target wasm .\main.go
$env:GOARCH = $goarch
$env:GOOS = $env:GOOS

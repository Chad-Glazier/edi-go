$goarch = $env:GOARCH
$goos = $env:GOOS
$env:GOARCH = "wasm"
$env:GOOS = "js"
tinygo build -o .\target\tiny_edi.wasm -target wasm .\main.go
go build -o .\target\edi.wasm .\main.go
$env:GOARCH = $goarch
$env:GOOS = $goos

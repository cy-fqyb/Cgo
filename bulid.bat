go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
echo now the CGO_ENABLED:
    go env CGO_ENABLED

echo now the GOOS:
    go env GOOS

echo now the GOARCH:
    go env GOARCH

go build main.go
echo "wait building"
go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64

echo now the CGO_ENABLED:
    go env CGO_ENABLED

echo now the GOOS:
    go env GOOS

echo now the GOARCH:
    go env GOARCH

echo "build success"
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

@echo off

rem 设置 PowerShell 脚本文件的路径
set psScriptPath=E:\go\Cgo\push.ps1

rem 使用 powershell.exe 执行脚本
powershell.exe -File "%psScriptPath%"

rem 输出消息
echo push execution completed.

# 设置环境变量
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64

# 输出 CGO_ENABLED、GOOS 和 GOARCH 的当前值
Write-Host "Now the CGO_ENABLED:"
go env CGO_ENABLED

Write-Host "Now the GOOS:"
go env GOOS

Write-Host "Now the GOARCH:"
go env GOARCH

# 构建 main.go
go build main.go
Write-Host "Wait building"

# 恢复环境变量
go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64

# 输出 CGO_ENABLED、GOOS 和 GOARCH 的当前值
Write-Host "Now the CGO_ENABLED:"
go env CGO_ENABLED

Write-Host "Now the GOOS:"
go env GOOS

Write-Host "Now the GOARCH:"
go env GOARCH

Write-Host "Build success"

$server = "159.75.155.236"
$username = "root"
$localFile = "E:\go\Cgo\main"
$remoteDirectory = "/root/cy/Cgo"
$remoteScript = "/root/cy/Cgo/restart.sh"

# 远程服务器上的可执行文件路径
$remoteExecutablePath = "./main"

# 检查远程服务器上是否有运行中的进程
$processToStop = ssh $username@$server "pgrep -f '$remoteExecutablePath'"

if ($processToStop -ne $null) {
    Write-Host "Terminating old process with PID $processToStop"
    
    # 杀死远程服务器上的进程
    ssh $username@$server "kill -TERM $processToStop"
    
    Start-Sleep -Seconds 3  # 根据需要调整等待时间
    
    Write-Host "Stop service success"
} else {
    Write-Host "No running process found"
}

# 使用 SCP 推送文件到服务器
& scp.exe $localFile "${username}@${server}:${remoteDirectory}"
# 输出消息
Write-Host "Push execution completed."
# 启动服务
# Write-Host "Starting service..."
# 在这里添加启动服务的命令，例如 systemctl start your_service

# 远程执行 Shell 脚本
# ssh.exe "${username}@${server}" "bash ${remoteScript}"
# Write-Host "Service updated successfully."



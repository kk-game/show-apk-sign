@echo off

@set GO111MODULE=on
@set GOPROXY=https://goproxy.cn,direct

@set CGO_ENABLED=0
@set GOOS=windows
@set GOARCH=amd64

if exist "build\" (echo build文件夹存在, 继续编译!) else (
    mkdir "build"
    echo 创建build文件夹成功
)

if exist "build\show_apk_sign.exe" (
	del "build\show_apk_sign.exe"
)

go build -o build/show_apk_sign.exe main.go
xcopy "templates\*" "build\templates\*" /Y
xcopy "tools\*" "build\tools\*" /Y
@rem copy "tools\keytool.exe" "build\tools\keytool.exe"

start %cd%\build\
echo 2 秒后退出...
timeout /t 2 /nobreak > nul
@echo off

@set GO111MODULE=on
@set GOPROXY=https://goproxy.cn,direct

@set CGO_ENABLED=0
@set GOOS=windows
@set GOARCH=amd64

if exist "build\" (echo build�ļ��д���, ��������!) else (
    mkdir "build"
    echo ����build�ļ��гɹ�
)

if exist "build\show_apk_sign.exe" (
	del "build\show_apk_sign.exe"
)

go build -o build/show_apk_sign.exe main.go
xcopy "templates\*" "build\templates\*" /Y
xcopy "tools\*" "build\tools\*" /Y
@rem copy "tools\keytool.exe" "build\tools\keytool.exe"

start %cd%\build\
echo 2 ����˳�...
timeout /t 2 /nobreak > nul
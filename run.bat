@echo off
if exist go-tiktok-lite.exe del go-tiktok-lite.exe
set GOPROXY=https://goproxy.io
:: 如果未在环境变量中设置GO的路径，则请在下方GOPATH处设置
set GOPATH=
SET PATH=%PATH%;%GOPATH%
go build
echo tiktok server is running
go-tiktok-lite.exe
pause>nul
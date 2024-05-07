@echo off
copy ..\..\bin\sheeter.exe .\
rmdir /s /q codeCs
rmdir /s /q codeGo
rmdir /s /q json
.\sheeter.exe build --config verify.yaml
dotnet csharpier .
gofmt -w .
del /f sheeter.exe
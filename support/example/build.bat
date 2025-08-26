@echo off
copy ..\..\bin\sheeter.exe .\
rmdir /s /q sheet
.\sheeter.exe build --config config.yaml
csharpier format .
gofmt -w .
del /f sheeter.exe
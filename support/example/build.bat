@echo off
copy ..\..\bin\sheeter.exe .\
rmdir /s /q sheet
.\sheeter.exe build --config config.yaml
@REM 如果已安裝csharpier或gofmt，可以取消下方命令的註解以格式化生成的程式碼
@REM 注意：請確保這些工具已經安裝並且在環境變量PATH中正確設置
@REM dotnet csharpier .
@REM gofmt -w .
del /f sheeter.exe
copy ..\..\..\bin\sheeter.exe .\
rmdir /s /q enum
rmdir /s /q json
rmdir /s /q proto
rmdir /s /q template
.\sheeter.exe build --config verify.yaml
dotnet csharpier .
gofmt -w .
buf format -w .
del /f sheeter.exe
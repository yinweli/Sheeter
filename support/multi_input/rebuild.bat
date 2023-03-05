copy ..\..\bin\sheeter.exe .\
rmdir /s /q json
rmdir /s /q proto
rmdir /s /q template
.\sheeter.exe build --config multi_input.yaml
dotnet csharpier .
gofmt -w .
buf format -w .
del /f sheeter.exe
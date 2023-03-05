copy ..\..\bin\sheeter.exe .\
rmdir /s /q src
rmdir /s /q json
rmdir /s /q proto
rmdir /s /q template
7z e -y -osrc benchmark_count.7z
.\sheeter.exe build --config benchmark.yaml
dotnet csharpier .
gofmt -w .
buf format -w .
del /f sheeter.exe
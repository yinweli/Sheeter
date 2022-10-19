rmdir /s /q json
rmdir /s /q proto
rmdir /s /q template
.\sheeter.exe build --config example.yaml
del /f sheeter.exe
rmdir /s /q json
rmdir /s /q proto
rmdir /s /q template
del /f protoCs.bat
del /f protoCs.sh
del /f protoGo.bat
del /f protoGo.sh
7z e -y benchmark_count.7z
.\sheeter.exe build --config benchmark.yaml
call protoCs.bat
call protoGo.bat
del /f sheeter.exe
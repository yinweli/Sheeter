rmdir /s /q json
rmdir /s /q proto
rmdir /s /q template
del /f protoCs.bat
del /f protoCs.sh
del /f protoGo.bat
del /f protoGo.sh
.\sheeter.exe build --config verify.yaml
call protoCs.bat
call protoGo.bat
del /f sheeter.exe
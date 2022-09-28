rm -rf json
rm -rf proto
rm -rf template
rm -f protoCs.bat
rm -f protoCs.sh
rm -f protoGo.bat
rm -f protoGo.sh
.\sheeter.exe build --config verify.yaml
call protoCs.bat
call protoGo.bat
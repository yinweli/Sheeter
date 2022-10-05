rm -rf json
rm -rf proto
rm -rf template
rm protoCs.bat
rm protoCs.sh
rm protoGo.bat
rm protoGo.sh
sheeter build --config example.yaml
./protoCs.sh
./protoGo.sh
rm sheeter
rm -rf json
rm -rf proto
rm -rf template
rm protoCs.bat
rm protoCs.sh
rm protoGo.bat
rm protoGo.sh
7za x benchmark_count.7z
sheeter build --config benchmark.yaml
./protoCs.sh
./protoGo.sh
rm sheeter
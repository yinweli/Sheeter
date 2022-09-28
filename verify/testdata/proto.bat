mkdir code
protoc --experimental_allow_proto3_optional --proto_path=./target --go_out=./code ./target\realData.proto
protoc --experimental_allow_proto3_optional --proto_path=./target --go_out=./code ./target\a.proto
protoc --experimental_allow_proto3_optional --proto_path=./target --go_out=./code ./target\s.proto
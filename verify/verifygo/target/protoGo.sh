# Code generated by sheeter. DO NOT EDIT.
# Sheeter: https://github.com/yinweli/Sheeter

mkdir proto\codeGo
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --go_out=./proto\codeGo ./proto\schema\item.proto
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --go_out=./proto\codeGo ./proto\schema\reward.proto
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --go_out=./proto\codeGo ./proto\schema\verifyData1.proto
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --go_out=./proto\codeGo ./proto\schema\verifyData2.proto

# Code generated by sheeter. DO NOT EDIT.
# Sheeter: https://github.com/yinweli/Sheeter

mkdir proto\codeCs
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --csharp_out=./proto\codeCs ./proto\schema\exampleData.proto
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --csharp_out=./proto\codeCs ./proto\schema\item.proto
protoc --experimental_allow_proto3_optional --proto_path=./proto\schema --csharp_out=./proto\codeCs ./proto\schema\reward.proto
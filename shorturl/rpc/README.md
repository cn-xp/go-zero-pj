# 生成proto模板
goctl rpc template -o transform.proto

# 生成rpc代码
goctl rpc protoc transform.proto --go_out=. --go-grpc_out=. --zrpc_out=.
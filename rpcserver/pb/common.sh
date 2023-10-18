protoc -I ./ --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./commonV3.proto
 #goctl rpc protoc friend.proto --go_out=.. --go-grpc_out=.. --zrpc_out=../    
 #10440  cp /Users/user/go/src/github.com/difftim/friend/rpcserver/pb/commonV3.proto protobuf/commonV3.proto
# 10441  cp /Users/user/go/src/github.com/difftim/friend/rpcserver/pb/friend.proto protobuf

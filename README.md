# Go-Grpc-practice

Gin + swagger + Grpc

安裝 protoc
sudo apt update
sudo apt install -y protobuf-compiler

安裝 protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"
source ~/.bashrc

使用以下指令來生成 Go 文件
cd ~/go_grpc_practice
protoc --proto_path=/home/tofu/go_grpc_practice/pkg/proto --go_out=. ~/go_grpc_practice/pkg/proto/*.proto
protoc --proto_path=/home/tofu/go_grpc_practice/pkg/proto --go-grpc_out=. ~/go_grpc_practice/pkg/proto/*.proto

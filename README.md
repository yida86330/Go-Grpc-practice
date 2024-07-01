# Go-Grpc-practice

Gin + swagger + Grpc + Gorm + Sqlite/Mysql

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
protoc --proto_path=/home/tofu/go_grpc_practice/internal/proto --go_out=. ~/go_grpc_practice/internal/proto/*.proto
protoc --proto_path=/home/tofu/go_grpc_practice/internal/proto --go-grpc_out=. ~/go_grpc_practice/internal/proto/*.proto

安裝GORM和SQLite
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

安裝mysql
sudo apt install mysql-server mysql-client
sudo apt install libmysqlclient-dev

解決wsl無法連接mysql
https://stackoverflow.com/questions/64883580/wsl-cant-connect-to-local-mysql-server-through-socket-var-run-mysqld-mysqld

chmod 777 -R /var/run/mysqld
chmod 777 -R /var/lib/mysql
chmod 777 -R /var/log/mysql
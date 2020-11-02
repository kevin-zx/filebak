// rpc
protoc -I pkg/service/grpc/ pkg/service/grpc/backup.proto --go_out=plugins=grpc:pkg/service/grpc/ 

// build
go build -o bin/fbak cmd/filebak/main.go

// deploy
scp bin/fbak root@dsdata:/home/kevinzx/filebak/


"# filebak" 

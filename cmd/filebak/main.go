package main

import (
	"filebak/pkg/service"
	backup_rpc "filebak/pkg/service/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := 11858
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	s, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	backup_rpc.RegisterBackupServiceServer(grpcServer, backup_rpc.NewBackupRPCService(s))
	log.Printf("start backup_rpc service at port: %d \n", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server:%v", err)
	}

}

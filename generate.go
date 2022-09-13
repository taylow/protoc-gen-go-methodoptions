package main

//go:generate protoc --go_out=. --go_opt=paths=source_relative example/permissions/permissions.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-methodoptions_out=. --go-methodoptions_opt=paths=source_relative example/service.proto

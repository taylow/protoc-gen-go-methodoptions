// Code generated by protoc-gen-go-methodoptions. DO NOT EDIT.
// source file:   example/service.proto
// proto package: helloworld

package example

import (
	permissions "github.com/taylow/protoc-gen-go-methodoptions/example/permissions"
)

// Greeter

const (
	Greeter_SayBye_Role    = permissions.Role_ROLE_GUEST
	Greeter_SayHello_Role  = permissions.Role_ROLE_OWNER
	Greeter_SayHello_Scope = "aaaaa"
	Greeter_SayHello_Test  = 1234
)

// SayByeRole returns the Role for SayBye.
func (UnimplementedGreeterServer) SayByeRole() permissions.Role {
	return Greeter_SayBye_Role
}

// Role returns the Role for SayBye request.
func (*ByeRequest) Role() permissions.Role {
	return Greeter_SayBye_Role
}

// SayHelloRole returns the Role for SayHello.
func (UnimplementedGreeterServer) SayHelloRole() permissions.Role {
	return Greeter_SayHello_Role
}

// Role returns the Role for SayHello request.
func (*HelloRequest) Role() permissions.Role {
	return Greeter_SayHello_Role
}

// SayHelloScope returns the Scope for SayHello.
func (UnimplementedGreeterServer) SayHelloScope() string {
	return Greeter_SayHello_Scope
}

// Scope returns the Scope for SayHello request.
func (*HelloRequest) Scope() string {
	return Greeter_SayHello_Scope
}

// SayHelloTest returns the Test for SayHello.
func (UnimplementedGreeterServer) SayHelloTest() int32 {
	return Greeter_SayHello_Test
}

// Test returns the Test for SayHello request.
func (*HelloRequest) Test() int32 {
	return Greeter_SayHello_Test
}

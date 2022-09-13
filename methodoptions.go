package main

type Proto struct {
	FileName  string
	GoPackage string
	Package   string
	Imports   map[string]string
	Services  map[string]Service
}

type Service struct {
	Name    string
	Methods map[string]Method
}

type Method struct {
	Name           string
	RequestName    string
	RequestPackage string
	Options        map[string]Option
}

type Option struct {
	Name    string
	Package string
	Type    string
	Value   string
}

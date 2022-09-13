package main

import (
	"bytes"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	fileExtension = ".options.pb.go"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		extTypes := new(protoregistry.Types)
		for _, file := range gen.Files {
			if err := registerAllExtensions(extTypes, file.Desc); err != nil {
				panic(err)
			}
		}

		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}

			buf := bytes.NewBuffer(nil)
			protoData := Proto{
				FileName:  file.Desc.Path(),
				Package:   string(file.Desc.Package()),
				GoPackage: string(file.GoPackageName),
				Imports:   make(map[string]string, 0),
				Services:  make(map[string]Service, 0),
			}

			for _, service := range file.Services {
				serviceData := Service{
					Name:    string(service.Desc.Name()),
					Methods: make(map[string]Method, 0),
				}

				for _, method := range service.Methods {
					methodData := Method{
						Name:        string(method.Desc.Name()),
						RequestName: string(method.Input.GoIdent.GoName),
						Options:     make(map[string]Option, 0),
					}

					// if request is in a different package, add it to the imports
					// if method.Input.GoIdent.GoImportPath != file.GoImportPath {
					// 	importAlias := strings.ReplaceAll(
					// 		string(method.Input.Desc.ParentFile().Package()),
					// 		".",
					// 		"_",
					// 	)
					// 	importPath := strings.Trim(method.Input.GoIdent.GoImportPath.String(), "\"")
					// 	protoData.Imports[importAlias] = importPath
					// 	methodData.RequestPackage = importAlias
					// }

					options, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
					if !ok || options == nil {
						continue
						// panic("not a MethodOptions")
					}

					b, err := proto.Marshal(options)
					if err != nil {
						panic(err)
					}

					options.Reset()
					err = proto.UnmarshalOptions{Resolver: extTypes}.Unmarshal(b, options)
					if err != nil {
						panic(err)
					}

					options.ProtoReflect().
						Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
							if !fd.IsExtension() {
								return true
							}

							importAlias := strings.ReplaceAll(
								string(fd.ParentFile().Package()),
								".",
								"_",
							)
							importPath := strings.Split(fd.ParentFile().Options().(*descriptorpb.FileOptions).GetGoPackage(), ";")[0]
							protoData.Imports[importAlias] = importPath

							optionData := Option{
								Name:    string(fd.TextName()),
								Package: importAlias,
								Type:    fd.Kind().String(),
								Value:   v.String(),
							}

							if fd.Kind() == protoreflect.EnumKind {
								optionData.Value = string(fd.Enum().
									Values().
									ByNumber(v.Enum()).
									Name())

								nameParts := strings.Split(string(fd.Enum().FullName()), ".")
								optionData.Name = nameParts[len(nameParts)-1]
							} else {
								nameParts := strings.Split(string(fd.FullName()), ".")
								optionData.Name = nameParts[len(nameParts)-1]
								optionData.Name = strings.Title(optionData.Name)
							}

							methodData.Options[optionData.Name] = optionData
							return true
						})

					serviceData.Methods[string(method.Desc.Name())] = methodData
				}

				protoData.Services[string(service.Desc.Name())] = serviceData
			}

			err := template.Must(template.New("").Parse(goTemplate)).Execute(buf, protoData)
			if err != nil {
				panic(err)
			}

			filename := file.GeneratedFilenamePrefix + fileExtension
			file := gen.NewGeneratedFile(filename, ".")
			// log.Println("writing to file:", filename)
			file.Write(buf.Bytes())
		}

		return nil
	})
}

// Recursively register all extensions into the provided protoregistry.Types,
// starting with the protoreflect.FileDescriptor and recursing into its MessageDescriptors,
// their nested MessageDescriptors, and so on.
//
// This leverages the fact that both protoreflect.FileDescriptor and protoreflect.MessageDescriptor
// have identical Messages() and Extensions() functions in order to recurse through a single
// function
func registerAllExtensions(extTypes *protoregistry.Types, descs interface {
	Messages() protoreflect.MessageDescriptors
	Extensions() protoreflect.ExtensionDescriptors
}) error {
	mds := descs.Messages()
	for i := 0; i < mds.Len(); i++ {
		registerAllExtensions(extTypes, mds.Get(i))
	}
	xds := descs.Extensions()
	for i := 0; i < xds.Len(); i++ {
		if err := extTypes.RegisterExtension(dynamicpb.NewExtensionType(xds.Get(i))); err != nil {
			return err
		}
	}
	return nil
}

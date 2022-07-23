## protobuf go DEMO
google protobuf有两个版本，分为v1和v2
旧版本，使用MessageV1，已弃用Deprecated  
https://github.com/golang/protobuf  
https://pkg.go.dev/github.com/golang/protobuf

新版，使用MessageV2  
https://github.com/protocolbuffers/protobuf-go  
https://pkg.go.dev/google.golang.org/protobuf

protoc编译器下载 https://github.com/protocolbuffers/protobuf/releases/latest

protoc-gen-go下载https://github.com/protocolbuffers/protobuf-go/releases/latest

### 使用protoc编译生成pb.go文件
将protoc添加到系统路径中
protoc-gen-go修改名称为protoc-gen-goV2,放入$GOPATH/bin中

编写待编译的.proto文件  
添加option go_package=""，其中填入包的路径

使用命令  
protoc --goV2_out=PATH XXX.proto  
例如 protoc --goV2_out=./ person.proto  

PATH为n生成pb.go文件的目录,XXX.proto为需要编译的.proto文件

### proto主要转换工具
google.golang.org/protobuf/proto
中的Marshal与Unmarshal函数进行序列化为[]byte与从[]byte反序列化  

google.golang.org/protobuf/encoding/protojson
中的Marshal与Unmarshal函数转换成json([]byte)或将json([]byte)转换成proto结构体格式

google.golang.org/protobuf/encoding/prototext
Format函数将proto结构转换成string


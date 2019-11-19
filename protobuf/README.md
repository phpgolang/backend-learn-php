# 何为ProtoBuf
官方文档的定义和描述
> Protocol Buffers 是一种语言无关、平台无关、可扩展的序列化结构数据的方法，它可用于（数据）通信协议、数据存储等。  
Protocol Buffers 是一种灵活，高效，自动化机制的结构数据序列化方法-可类比XML，但是比XML更小，更快，更为简单。  
你可以定义数据的结构，然后使用特殊生成的源代码轻松的在各种数据流中使用各种语言进行编写和读取结构数据。你甚至可以更新数据结构，而不破坏由旧数据结构编译的已部署程序

我理解的 ProtoBuf 就是一种序列化的数据结构和XML、JSON一样，通过编译为二进制文件流，更快的传输

[proto3语言指南](./proto3.md)

protoc工具下载
https://github.com/protocolbuffers/protobuf/releases

根据.proto文件生成pd.go文件
```
protoc --go_out=plugins=grpc:. *.proto
```

import "github.com/golang/protobuf/proto"

    info := &pbdata.SuggestInfo{}
bizData, err := proto.Marshal(info)
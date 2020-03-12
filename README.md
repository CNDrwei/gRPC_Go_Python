# demo1.0介绍
[参考文档：亲测可行，感谢李老师](https://www.liwenzhou.com/posts/Go/gRPC/)

1.如果未安装protobuf请自行[安装](https://github.com/protocolbuffers/protobuf/releases),下载完毕解压之后，将bin目录添加到环境变量

2.Go安装proto-gen-geo的插件
```go
go get -u github.com/golang/protobuf/protoc-gen-go
```

3.编写proto代码
gRPC是一种与语言无关的，用于序列化结构化数据的
详情见addressbook.proto
一般用于生成.pb文件指令`proto --go_out=./待转化的文件`
这里需要将将其转化为支持grpc的文件
`模板：protoc --语言_out=plugins=grpc:{输出目录} {proto文件}
示例：E:\aGo\test>protoc --go_out=plugins=grpc:. ./protobuf/addressbook.proto
`
之后就会在protobuf文件夹下生成addressbook.pb.go文件


4.编写Server端Go代码，详情见server/server.go

5.编写client端Go代码，详情见client/client.go

6.TIPS先运行服务端代码，再运行客户端代码，这样才能看到结果

7.跨语言调试结果展示，在Python环境中需要安装gRPC`pip install grpcio    pip install protobuf    pip install grpcio-tools`
`python（语言） -m grpc_tools.protoc（需要在Python中安装的工具） -I ./protobuf --python_out=./client/ --grpc_python_out=./client ./protobuf/addressbook.proto`
8.在test目录下执行`E:\aGo\test>python -m grpc_tools.protoc -I ./protobuf --python_out=./client/ --grpc_python_out=./client ./protobuf/addressbook.proto`

9.在client下你就会看到自动生成两个py文件

10.新建Python文件的client详情见client.py，运行的时候如果没有配置Goland中Python环境，可以CMD运行

11.最大的缺点就是proto文件起名字的时候，乱起了一个，读者可以尝试换个名字，这样就能避免完全照抄，顺便理一下逻辑

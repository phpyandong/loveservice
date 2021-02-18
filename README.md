# Kratos Layout

## Install Kratos
```
go get github.com/go-kratos/kratos/cmd/kratos/v2
go get github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
go get github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

# from source
cd cmd/kratos && go install
cd cmd/protoc-gen-go-http && go install
cd cmd/protoc-gen-go-errors && go install
```
## Create a service
```
# create a template project
kratos new helloworld

cd helloworld
# Add a proto template
kratos proto add api/helloworld/helloworld.proto
# Generate the source code of service by proto file
kratos proto service api/helloworld/helloworld.proto -t internal/service

make proto
make build
make test
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```


── LICENSE
├── Makefile
├── README.md
├── api
│   └── helloworld
│       ├── errors
│       │   ├── helloworld.pb.go
│       │   ├── helloworld.proto
│       │   └── helloworld_errors.pb.go
│       ├── helloworld.proto
│       └── v1
│           ├── greeter.pb.
│           ├── greeter.proto
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd   //负责程序的：启动，关闭，配置初始化
│   └── helloworld
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── configs
│   └── config.yaml
├── go.mod
├── go.sum
└── internal //私有代码目录，go编译器本身支持，用于避免外部项目使用
    ├── biz
    │   ├── README.md
    │   ├── biz.go
    │   └── greeter.go
    ├── conf
    │   ├── conf.pb.go
    │   └── conf.proto
    ├── data
    │   ├── README.md
    │   ├── data.go
    │   └── greeter.go
    ├── server
    │   ├── grpc.go
    │   ├── http.go
    │   └── server.go
    └── service
        ├── README.md
        ├── greeter.go
        ├── helloworld.go
        └── service.go


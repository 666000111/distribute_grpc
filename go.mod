module grpc_server

go 1.16

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.37.0

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/luci/go-render v0.0.0-20160219211803-9a04cc21af0f
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.7

)

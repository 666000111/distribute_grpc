# 分布式第二次作业

protoc -I book/ book/book.proto --go_out=plugins=grpc:proto_gen
这个是proto生成rpc代码用的

我这个demo基本是仿照google官方写的那个route——guide，但是因为不太会配置秘钥所以我这个选择了grpc.WithInsecure的模式


idl放在了book文件夹下，其中query可以满足提的两个要求，可以进行批量查询也支持单独查询
就是效率和单独查询比会第一点，因为这个用的是in，不过对于作业demo问题不大

数据库用的本地的mysql，其中链接数据库用的go里面的gorm库


Client我放在了client文件夹下就
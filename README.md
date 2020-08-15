### 一、基础相关

#### 1.新建一个服务
~~~~
micro new --namespace=user --type=service user-service
~~~~
参数解释:
~~~~
--namespace     服务命令空间
–-type          服务类型（service、web）
最后是目录名
~~~~
注意：目前的模板中很多文件不需要,请按顺序删除 
~~~~
subscriber目录
.gitignore
Dockerfile
generate.go
go.mod
Makefile
plugin.go
README.md
~~~~
#### 2.编译proto文件
~~~~
protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto
~~~~
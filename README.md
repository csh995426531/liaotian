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
cd liaotian
protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto
~~~~

### 二、目录结构
- [.circleci](https://www.baidu.com/) - circleci 持续集成工具的配置
- [argocd-deploy](#argocd-deploy) - argocd 持续部署的配置，里面放服务部署文件。（argocd的安装配置在deploy下）
- [data](#data) - 临时数据
  - [](#)
- [deploy](#deploy) - 部署文件，业务服务其实是用argocd来部署了
  - [apisix](#apisix) - apisix 网关，熔断限流应该可以用这个
  - [argocd](#argocd) - argocd 持续部署工具
  - [friend-service](#friend-service)
  - [friend-web](#friend-web)
  - [k8s](#k8s) - micro使用k8s作为服务注册，需要配置RBAC权限
  - [monitor](#monitor) - 服务监控
  - [traefik](#traefik) - traefik 网关，不用了，用apisix
  - [user-service](#user-service)
  - [user-web](#user-web)
- [friend-service](#friend-service) - 朋友服务
- [friend-web](#friend-web) - 朋友api
- [plugins](#plugins) - 插件
  - [wrapper](#wrapper) - 包装器
- [user-service](#user-service) - 用户服务
- [user-web](#user-web) - 用户api
- project-dirs - 服务名单 (circleci中有匹配文件中的服务名单)
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
- [.circleci](https://github.com/csh995426531/liaotian/tree/master/.circleci) - circleci 持续集成工具的配置
- [argocd-deploy](https://github.com/csh995426531/liaotian/tree/master/argocd-deploy) - argocd 持续交付的配置，里面放服务部署文件。（argocd的安装配置在deploy下）
- [data](#data) - 临时数据
- [deploy](https://github.com/csh995426531/liaotian/tree/master/deploy) - 部署文件，业务服务其实是用argocd来部署了
  - [apisix](https://github.com/csh995426531/liaotian/tree/master/deploy/apisix) - apisix 网关，熔断限流应该可以用这个
  - [argocd](https://github.com/csh995426531/liaotian/tree/master/deploy/argocd) - argocd 持续部署工具
  - [friend-service](https://github.com/csh995426531/liaotian/tree/master/deploy/friend-service)
  - [friend-web](https://github.com/csh995426531/liaotian/tree/master/deploy/friend-web)
  - [k8s](https://github.com/csh995426531/liaotian/tree/master/deploy/k8s) - micro使用k8s作为服务注册，需要配置RBAC权限
  - [monitor](https://github.com/csh995426531/liaotian/tree/master/deploy/monitor) - 服务监控
  - [traefik](https://github.com/csh995426531/liaotian/tree/master/deploy/traefik) - traefik 网关，不用了，用apisix
  - [user-service](https://github.com/csh995426531/liaotian/tree/master/deploy/user-service)
  - [user-web](https://github.com/csh995426531/liaotian/tree/master/deploy/user-web)
- [friend-service](#) - 朋友服务
- [friend-web](#) - 朋友api
- [plugins](#) - 插件
  - [wrapper](https://github.com/csh995426531/liaotian/tree/master/plugins/wrapper) - 包装器
- [user-service](#) - 用户服务
- [user-web](#) - 用户api
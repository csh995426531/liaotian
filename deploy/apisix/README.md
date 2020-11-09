### 介绍
apisix是一个高性能网关，具体使用方式和文档在 [项目地址](https://github.com/apache/apisix)

需要注意几处：
1. manager_conf目录中build.sh内的mysql配置需改为自己的
2. apisix-gw-config-cm.yaml内etcd.host需改为自己的
3. script/db目录中放着mysql初始配置

可以使用Makefile中的命令启用apisix，使用127.0.0.1:180访问dashboard
- 启用：make apply
- 停止：make delete

使用kubectl 命令查看
- kubectl get all -n liaotian
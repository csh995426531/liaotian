###
基于go micro 1.18搞的练习项目，从0开始。
包含k8s、微服务网关、ci、cd、日志等部署到使用。
### 目录结构
- [.circleci](https://github.com/csh995426531/liaotian/tree/master/.circleci) - circleci 持续集成工具的配置
- [.gitlab-ci](https://github.com/csh995426531/liaotian/tree/master/.gitlab-ci) - gitlab-ci使用的ci配置 （**目前都是用的这个，circleci有点慢**）
- [data](#data) - 临时数据
- [app](https://github.com/csh995426531/liaotian/tree/master/app) - 应用服务
  - [im](https://github.com/csh995426531/liaotian/tree/master/app/im) - im 全在这里
- [domain](https://github.com/csh995426531/liaotian/tree/master/domain) - 领域服务
  - [user](https://github.com/csh995426531/liaotian/tree/master/domain/user) - 用户服务
    - [cmd](https://github.com/csh995426531/liaotian/tree/master/domain/user/cmd) - 入口
    - [entity](https://github.com/csh995426531/liaotian/tree/master/domain/user/entity) - 实体+仓储实现
    - [handler](https://github.com/csh995426531/liaotian/tree/master/domain/user/handler) - 服务处理
    - [proto](https://github.com/csh995426531/liaotian/tree/master/domain/user/proto) - 
    - [repository](https://github.com/csh995426531/liaotian/tree/master/domain/user/repository) - 仓库
- [deploy](https://github.com/csh995426531/liaotian/tree/master/deploy) - 相关的部署文件
  - [apisix](https://github.com/csh995426531/liaotian/tree/master/deploy/apisix) - Apisix 网关
  - [ArgoCd](https://github.com/csh995426531/liaotian/tree/master/deploy/argocd) - ArgoCd 持续交付工具
  - [EFK](https://github.com/csh995426531/liaotian/tree/master/deploy/efk) - EFK 日志收集
  - [GitLab](https://github.com/csh995426531/liaotian/tree/master/deploy/gitlab) - GitLab 代码托管
  - [k8s](https://github.com/csh995426531/liaotian/tree/master/deploy/k8s) - micro使用k8s作为服务注册，需要配置RBAC权限
  - [monitor](https://github.com/csh995426531/liaotian/tree/master/deploy/monitor) - 服务监控
  - [SkyWalking](https://github.com/csh995426531/liaotian/tree/master/deploy/skywalking) - SkyWalking 链路追踪
  - [traefik](https://github.com/csh995426531/liaotian/tree/master/deploy/traefik) - traefik 网关，(**不用了**)
  - [user-service](https://github.com/csh995426531/liaotian/tree/master/deploy/user-service) 本地调试偶尔用
  - [user-web](https://github.com/csh995426531/liaotian/tree/master/deploy/user-web) 本地调试偶尔用
- [middlewares](https://github.com/csh995426531/liaotian/tree/master/middlewares) - 插件
  - [common-result](https://github.com/csh995426531/liaotian/tree/master/middlewares/common-result]) - 公共返回
  - [logger](https://github.com/csh995426531/liaotian/tree/master/middlewares/logger]) - 日志
  - [validate](https://github.com/csh995426531/liaotian/tree/master/middlewares/validate]) - 验证
  - [wrapper](https://github.com/csh995426531/liaotian/tree/master/middlewares/wrapper]) - 包装器，（搞的不行，改了原插件，抽时间重新搞）
- [.gitlab-ci.yml](https://github.com/csh995426531/liaotian/tree/master/deploy/middlewares) - gitlab-ci配置
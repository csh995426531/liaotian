### 介绍
argocd是一个gitOps连续交付工具，具体使用方式和文档在 [项目地址](https://github.com/argoproj/argo-cd)

需要注意几处：
1. 默认账号是admin，默认密码使用以下命令查看
~~~
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2
~~~
这里碰到一个坑，密码无效，只能直接重置密码，A123456
~~~
kubectl -n argocd patch secret argocd-secret   -p '{"stringData": {
  "admin.password": "$2a$10$88NHgAw3gSbPmMGvPH8wl.E.wh/JpxF6LpAkN.3YzI8vCKqz92rpi",
  "admin.passwordMtime": "'$(date +%FT%T%Z)'"
}}'
~~~

使用127.0.0.1:7901访问dashboard

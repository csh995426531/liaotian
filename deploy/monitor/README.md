### 
**注意事项**：
1. 创建pv挂载目录
~~~
sudo mkdir /tmp/data/prometheus 
~~~
2. 启动prometheus服务
~~~
make apply
~~~
grafana dashboard 地址 (http://127.0.0.1:32213) - 用户名:admin 密码: prom-operator

prometheus dashboard 地址 (http://127.0.0.1:30090)
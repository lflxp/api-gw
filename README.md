# api-gw
api-gw for micro service and reverse proxy

# 介绍

本项目的作用是进行微服务api-gw网关反向代理的自研，实现动态热更新配置进行7层http代理

学习和开发项目，让小白明白其中的原理，实现自己的价值

# 范围

* http/https
* tcp/udp
* 负载均衡 nginx/lvs

# 使用

目前只是一个DEMO,使用方法如下：

* 修改被代理的链接配置　main.go 

```
h.addrs = []string{"192.168.50.216:8080"}
```

* 运行

> go run main.go

> go build main.go -o main && ./main

# Golang学习

## NET包

[解析ip地址](./code/package/net_ip/net_ip.go)  
[net.DialTCP](./code/package/net_dial_tcp/net_dial_tcp.go)  
[net.ListenTCP](./code/package/net_listen_tcp/net_listen_tcp.go)
[net.DialTimeout](./code/package/net_dial_timeout/net_dial_timeout.go)  
[udp](./code/package/udp/)
[tcp](./code/package/tcp/)  
[kcp](./code/package/kcp/)  
>kcp是一个基于udp实现快速、可靠、向前纠错的的协议，能以比TCP浪费10%-20%的带宽的代价，换取平均延迟降低30%-40%，且最大延迟降低三倍的传输效果。  
>github.com/xtaci/kcp-go 是用go实现了kcp协议的一个库.


## 知识点

[golang中type常用用法](book/golang_type.md)  
[go mod模式下设置代理](https://goproxy.io/)

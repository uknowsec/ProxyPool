## ProxyPool

### 原理

使用[小熊ip](https://www.xiaoxiongip.com/)获取代理ip，配合[Gost](https://github.com/go-gost/gost)的[WebApi](https://gost.run/tutorials/api/overview/)定时更新Gost配置来实现代理池功能。


### 使用方法

#### 下载安装Gost

```
# 安装最新版本 [https://github.com/go-gost/gost/releases](https://github.com/go-gost/gost/releases)
bash <(curl -fsSL https://github.com/go-gost/gost/raw/master/install.sh) --install
```

#### 配置文件
```
Interval: 8
ProxyApi:
  url: https://find.xiaoxiongip.com/find_s5?key=&count=1&type=json&only=0&pw=yes
Gost:
  socks5User: user
  socks5Pass: pass
  socks5Port: 8888
  apiurl: 127.0.0.1:8001
```

- Interval: 获取ip地址的间隔，可以根据小熊ip获取的代理的生效进行设置，比如申请10分钟时效的ip我们就可以设置为8
- url：获取代理ip的api,更换自己的key
- socks5: 配置socks代理的账号密码
- apiurl: gost webapi地址，程序通过这个webapi定时更新上层代理

#### 运行程序

```
./ProxyPool 
2024/06/28 21:35:57 Interval: 8
2024/06/28 21:35:57 Proxy API URL: https://find.xiaoxiongip.com/find_s5?key=&count=1&type=json&only=0&pw=yes
2024/06/28 21:35:57 SOCKS5 Username: user
2024/06/28 21:35:57 SOCKS5 Password: pass
2024/06/28 21:35:57 SOCKS5 Port: 8888
2024/06/28 21:35:57 API URL: 127.0.0.1:8001
2024/06/28 21:35:57 Updating ProxyServer information:
2024/06/28 21:35:57 Server: 117.68.38.140
2024/06/28 21:35:57 Port: 26911
2024/06/28 21:35:57 User: 
2024/06/28 21:35:57 Password: 
2024/06/28 21:35:57 {"name":"chain-0","hops":[{"name":"hop-0","nodes":[{"name":"node-0","addr":"117.68.38.140:26911","connector":{"type":"socks5","auth":{"username":"","password":""}},"dialer":{"type":"tcp","tls":{"serverName":"117.68.38.140"}}}]}]}
2024/06/28 21:35:57 PUT hop successful
```

#### 代理测试
```
curl -x  socks5://127.0.0.1:8888 -U user:pass cip.cc
IP      : 120.242.163.87
地址    : 中国  安徽  
运营商  : 移动

数据二  : 安徽省宿州市 | 移动

数据三  : 中国安徽省宿州市 | 移动

URL     : http://www.cip.cc/120.242.163.87

```

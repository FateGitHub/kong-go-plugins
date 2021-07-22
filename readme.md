# 启动
```
docker build -t kong-demo .
```

```
docker run -ti --rm --name kong-go-plugins \
  -e "KONG_DATABASE=off" \
  -e "KONG_GO_PLUGINS_DIR=/tmp/go-plugins" \
  -e "KONG_DECLARATIVE_CONFIG=/tmp/config.yml" \
  -e "KONG_PLUGINS=key-checker" \
  -e "KONG_PROXY_LISTEN=0.0.0.0:8000" \
  -p 8000:8000 \
  -p 8443:8443 \
  -p 8001:8001 \
  -p 8444:8444 \
  kong-demo
```

# 访问
```
curl localhost:8000/?key=test

curl localhost:8000/?key=mysecretconsumerkey | json_pp
```

# 本地开发
#安装
```
https://www.cnblogs.com/mbpframework/p/12891866.html
```

# 编译.so
```
go get -d -v github.com/Kong/go-pluginserver &&
go build github.com/Kong/go-pluginserver &&
go build -buildmode plugin key-checker.go
生成 go-pluginserver文件，复制到/usr/local/bin目录
将生成的.so文件放到go_plugins_dir定义的目录中
```

# 修改/etc/kong/kong.conf
```
plugins = bundled,key-checker,custom-rate-limiting
go_plugins_dir = /etc/kong/plugins
go_pluginserver_exe = /usr/local/bin/go-pluginserver
```

# 添加服务
```
curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=key-checker-service' \
--data 'url=https://reqres.in/api/users?page=2'

curl -i -X POST \
--url http://localhost:8001/services/key-checker-service/routes \
--data 'name=key-checker' \
--data 'paths[]=/' \

curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=key-checker-service' \
--data 'url=https://reqres.in/api/users?page=2'

curl -i -X POST \
--url http://localhost:8001/services/key-checker-service/routes \
--data 'paths[]=/'

curl -i -X POST \
--url http://localhost:8001/services/key-checker-service/plugins/ \
--data 'name=key-checker'
```
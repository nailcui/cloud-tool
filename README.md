此服务提供一些工具，用以模拟各种场景
- return HTTP 200
- return HTTP 400
- return HTTP 500
- return hostname
- alloc memory 1M、5M、10M、100M、1G
- release memory

```shell script
kubectl apply -f manifests/
```

#### hostname
```
curl localhost/hostname
cloud-tool-sojfioweif
```

#### health
```
# 系统启动之后返回 200状态码，当执行过 `/health/setStatus` 之后状态码会被修改为指定状态
curl http://localhost/health

# 返回200状态码，之后请求 `/health` 接口将返回404
curl http://localhost/health/setStatus?status=404

# return HTTP 200
curl http://localhost/health/status?status=200

# return HTTP 404
curl http://localhost/health/status?status=404

# return HTTP 500
curl http://localhost/health/status?status=500
```

#### memory
```
# 申请内存，单位 M
curl http://localhost/memory/alloc?size=1

# 手动释放内存，并还给操作系统
curl http://localhost/memory/release

```

#### other
```
# 进程退出，code=1
curl http://localhost/exit
```

#### developing

```shell
# 打镜像

docker build . -t naildocker/cloud-tool:0.1.1

# 推送远程
# 在自己电脑上第一次可能需要 docker login
docker push naildocker/cloud-tool:0.1.1
```

#### roadmap

| version | desc |
| --- | --- |
| 0.1.2 | add os.signal listener |

# minids

it's a mini distributed system.
it provides tcp client, http client

### release
- go mod And gen main cmd:
```
docker run --rm -v ~/gopath:/gopath -v ~/workspace/znddzxx112/minids:/workspace centos7_golang:1.12.1 /go/bin/go build -o /workspace/main /workspace/minids.go 
```
- build app's image And run:
```
docker build -t minids:latest .
docker run -d --net=host --name minids_con minids:latest
```

### dev or debug:
- run server:
```
docker run -it --net=host --name minids_debug --rm -v ~/gopath:/gopath -v ~/workspace/znddzxx112/minids:/workspace centos7_golang:1.12.1 /go/bin/go run /workspace/minidsd.go
```

- build tcp client:
```
docker run -it --net=host --name minids_debug --rm -v ~/gopath:/gopath -v ~/workspace/znddzxx112/minids:/workspace centos7_golang:1.12.1 /go/bin/go build /workspace/minids.go
```


# minids

it's a mini distributed system.
it includes three parts which is apigateway,zookeeper,services.

## apigateway

it generates a reverse-proxy server which send the request to services

## zookeeper

it provides Service registration and discovery

## httppingservice

it's a http service named "ping"

## quick start
```
# cd $GOPATH/src
# git clone http://www.github.com/znddzxx112/minids
# go get github.com/gin-gonic/gin
```

### run zookeeper
```
# docker run --name minids-zookeeper --privileged --restart always -d registry.docker-cn.com/library/zookeeper

offical images of zookeeper[https://hub.docker.com/_/zookeeper/]

use zookeeper client to create "/httpping" node:

# docker run -it --rm --link minids-zookeeper:zookeeper registry.docker-cn.com/library/zookeeper zkCli.sh -server zookeeper
[zk: zookeeper(CONNECTED) 1] create /httpping httppintservice
```

### run httppingservice
```
# docker run -d --rm --name minids-httpping-service1 --link minids-zookeeper:zookeeper -v $GOPATH:/go --privileged registry.docker-cn.com/library/golang:latest go run /go/src/minids/httppingservice/main.go
# docker run -d --rm --name minids-httpping-service2 --link minids-zookeeper:zookeeper -v $GOPATH:/go --privileged registry.docker-cn.com/library/golang:latest go run /go/src/minids/httppingservice/main.go
# docker run -d --rm --name minids-httpping-service3 --link minids-zookeeper:zookeeper -v $GOPATH:/go --privileged registry.docker-cn.com/library/golang:latest go run /go/src/minids/httppingservice/main.go
```

### run apigateway
```
# docker run -d --rm --name minids-apigateway --link minids-zookeeper:zookeeper -v $GOPATH:/go --privileged registry.docker-cn.com/library/golang:latest go run src/minids/apigateway/main.go
```

### test service
```
# docker run -it --rm --name minids-test --link minids-zookeeper:zookeeper --link minids-apigateway:apigateway -v $GOPATH:/go --privileged registry.docker-cn.com/library/golang:latest bash
root@b433cdf05f1c:/go# cd src/minids/tests/
root@b433cdf05f1c:/go/src/minids/tests# go test
```

## to do
```
thrift service
apigateway translate restful api
tcp server
```


FROM golang:1.13.5 AS build

WORKDIR /minids

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go version && go env && mkdir -p dist \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/minids ./cmd/client \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/minidsd ./cmd/server \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/minids-web ./cmd/web \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/minids-ws ./cmd/websocket \
    && mkdir -p /app \
    && cp -r ./conf /app \
    && mv dist/* /app/

FROM alpine:latest
# 切换软件源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    #&& apk upgrade \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
    #&& apk --no-cache add ca-certificates \
    #&& apk add --no-cache bash git openssh gcc musl-dev \
    #&& apk --no-cache add bash
WORKDIR /app
COPY --from=build /app .
# 8115 minidsd
EXPOSE 8115
CMD ["./minidsd", "--config=./config/minidsd.yaml"]
#CMD ["./minids", "--config=./config/minids.yaml"]
#CMD ["./minids-web", "--config=./config/web.yaml"]
#CMD ["./minids-ws", "--config=./config/websocket.yaml"]

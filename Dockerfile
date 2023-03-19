FROM golang:1.19-alpine AS builder

# 设置环境变量
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 导入项目
COPY . /data/service/silgo

# 设置工作目录
WORKDIR /data/service/silgo

# 构建
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata && \
    CGO_ENABLED=0 go build -ldflags "-s -w" -o silgo

# 导入新的环境
FROM alpine as runner

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /data/service/silgo/silgo /data/service/
COPY --from=builder /data/service/silgo/configs /data/service/configs

WORKDIR /data/service

CMD ["./silgo", "silgo-api", "--config", "configs/dev.yaml"]

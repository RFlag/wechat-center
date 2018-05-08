FROM registry.cn-hangzhou.aliyuncs.com/ranbb/golang:alpine-git AS build
WORKDIR /go/src/wechat
COPY . /go/src/wechat
RUN set -x; \
    mkdir /app && \
    git rev-list --max-count=1 HEAD | tee /app/VERSION && \
    GOOS=linux GOARCH=amd64 go build -o /app/wechat main.go

FROM registry.cn-hangzhou.aliyuncs.com/ranbb/alpine:ca-ld-tz
HEALTHCHECK --interval=32s --timeout=16s --start-period=6s --retries=2 \
    CMD wget -nv -O - http://localhost:80/health || exit 1
EXPOSE 80
CMD ["/app/wechat"]
COPY --from=build /app /app

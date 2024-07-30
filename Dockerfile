FROM golang:1.22-bullseye AS builder

ENV GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    GOPROXY=goproxy.cn,direct

WORKDIR /src

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .
RUN make gen
RUN make build

FROM alpine:3.20

RUN apk add --no-cache tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
COPY --from=builder /src/_output/platforms/linux/amd64/godfrey-apiserver .
CMD ["./godfrey-apiserver"]
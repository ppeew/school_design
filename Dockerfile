#FROM alpine
#
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
#
#RUN apk update --no-cache
#RUN apk add --update gcc g++ libc6-compat
#RUN apk add --no-cache ca-certificates
#RUN apk add --no-cache tzdata
#ENV TZ Asia/Shanghai
#
#COPY ./main /main
#COPY ./config/settings.demo.yml /config/settings.yml
#COPY ./go-admin-db.db /go-admin-db.db
#EXPOSE 8000
#RUN  chmod +x /main
#CMD ["/main","server","-c", "/config/settings.yml"]

FROM golang:1.21-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/my_app /app/main.go

FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/my_app /app/my_app
COPY --from=builder /app/config/settings.demo.yml /config/settings.yml
COPY --from=builder /app/go-admin-db.db /go-admin-db.db

CMD ["./my_app","server","-c", "/config/settings.yml"]
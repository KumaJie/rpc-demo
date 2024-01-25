FROM golang:1.20

ENV GOPROXY https://goproxy.cn/,direct

WORKDIR /build

COPY ./src ./src
COPY ./script ./script
COPY go.mod .
COPY go.sum .

#RUN apt update \
#    && apt -y install ffmpeg
# 设置中文编码
ENV LANG C.UTF-8

RUN go mod download \
    && bash ./script/build_all.sh
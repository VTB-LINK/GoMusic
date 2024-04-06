FROM alpine:3

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ Asia/Shanghai

COPY ./GoMusic /GoMusic

EXPOSE 8081

CMD ["./GoMusic"]

FROM golang:alpine3.15 AS build
LABEL maintainer="{{author}}"
WORKDIR /build
COPY . .
RUN apk add git
RUN go build -o app .

FROM golang:alpine3.15 AS run
LABEL maintainer="{{author}}"

RUN addgroup -g 1000 golang && adduser -u 1000 -S -G golang -D golang
USER golang

WORKDIR /app
COPY --from=build /build/app .
ENTRYPOINT ["./app"]

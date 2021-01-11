#@version 3.7.5
FROM golang:1.14.6-alpine3.12 as base

WORKDIR /go/src/
COPY ./go/src ./is_nude_api
RUN rm -R is_nude_api/github.com


RUN apk add git
WORKDIR /go/src/is_nude_api

RUN go get -v ./...  \
    && go build -ldflags="-s -w"
RUN ls -la

CMD ["is_nude_api"]

#FROM scratch as prod
FROM alpine as prod

COPY --from=base /go/src/is_nude_api/is_nude_api .
COPY ./www ./public
CMD ["/is_nude_api"]

FROM golang:latest as server
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /go/src/github.com/o-sk/iad/server
COPY server .
RUN go mod download
RUN go build main.go

FROM node:10-alpine AS front
WORKDIR /work
COPY front .
RUN yarn
RUN yarn build

FROM alpine
WORKDIR /app
COPY --from=server /go/src/github.com/o-sk/iad/server /app
COPY --from=front /work/dist /app/dist
RUN apk add ca-certificates

CMD /app/main $PORT
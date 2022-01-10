FROM golang:1.17 AS build
WORKDIR /pastebean-create
ADD . .
RUN cd create; CGO_ENABLED=0 GOOS=linux go build -o /create main.go

FROM busybox

FROM scratch
COPY --from=build /create /
ENV GIN_MODE=release
CMD [ "/create" ]

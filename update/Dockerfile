FROM golang:1.17 AS build
WORKDIR /pastebean-update
ADD . .
RUN cd update; CGO_ENABLED=0 GOOS=linux go build -o /update main.go

FROM busybox

FROM scratch
COPY --from=build /update /
ENV GIN_MODE=release
CMD [ "/update" ]

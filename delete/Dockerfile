FROM golang:1.17 AS build
WORKDIR /pastebean-delete
ADD . .
RUN cd delete; CGO_ENABLED=0 GOOS=linux go build -o /delete main.go

FROM scratch
COPY --from=build /delete /
ENV GIN_MODE=release
CMD [ "/delete" ]

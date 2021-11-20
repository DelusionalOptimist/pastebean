FROM golang:1.17 AS build
WORKDIR /pastebean-read
ADD . .
RUN cd read; CGO_ENABLED=0 GOOS=linux go build -o /read main.go

FROM scratch
COPY --from=build /read /
CMD [ "/read" ]

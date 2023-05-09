FROM golang:alpine as build_container
WORKDIR /app
COPY ./Euprava-saobracajna/go.mod ./Euprava-saobracajna/go.sum ./
RUN go mod download
COPY ./Euprava-saobracajna/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
WORKDIR /root/
COPY --from=build_container /app/main .
COPY --from=build_container /app/fonts ./fonts

EXPOSE 8001
ENTRYPOINT ["./main"]
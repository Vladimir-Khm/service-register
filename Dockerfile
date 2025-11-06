FROM golang:alpine AS build

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build -o bin -ldflags "-s -w" -tags=go_json,viper_bind_struct ./app/main.go 


FROM alpine
COPY --from=build /app/bin /app/bin


EXPOSE 8080

WORKDIR /app

CMD ["./bin"]

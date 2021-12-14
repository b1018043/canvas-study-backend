FROM golang:alpine as BUILD

WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM scratch
COPY --from=BUILD /app/main .
CMD ["./main"]
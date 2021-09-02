# for cloud run
FROM golang:alpine AS builder
COPY . /app
WORKDIR /app
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -o main
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/main .
# COPY .env .
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./main"]
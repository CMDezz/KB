#Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl

#Run
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
# COPY app.env . 
COPY start.sh .
# COPY wait-for .

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "sh" ,"/app.start.sh"]
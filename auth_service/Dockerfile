FROM golang:1.21-rc-alpine

RUN apk add --no-cache make 
WORKDIR /app
COPY . .
COPY config.toml .

RUN make service

RUN ["chmod", "+x", "./build/dist/authService"]

CMD ["./build/dist/authService", "-settings", "config.toml"]
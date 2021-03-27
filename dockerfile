FROM golang:1.13-alpine3.11 AS gobuilder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o app.linux .


FROM alpine:3.11

EXPOSE 8080
WORKDIR /home/user/app/
COPY --from=gobuilder /build/app.linux .

CMD ["./app.linux"]
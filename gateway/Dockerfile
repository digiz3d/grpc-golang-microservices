FROM golang:latest as build
RUN apt update && apt install -y protobuf-compiler
WORKDIR /app
COPY . .
WORKDIR /app/comments
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN protoc --go_out=. --go-grpc_out=. ./pb/*.proto
WORKDIR /app/gateway
RUN protoc --go_out=. --go-grpc_out=. ./pb/*.proto
RUN go generate
RUN go build

FROM debian:latest as release
WORKDIR /bin
COPY --from=build /app/gateway/gateway .
CMD ["gateway"]
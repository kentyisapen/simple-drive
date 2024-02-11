FROM golang:1.21.5
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# protocのバージョンを指定
ARG PROTOC_VER=3.19.1

# protocコンパイラをダウンロードして解凍し、バイナリを/usr/local/binに移動
RUN apt-get update && apt-get install -y unzip && \
  curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip && \
  unzip protoc-${PROTOC_VER}-linux-x86_64.zip -d /usr/local && \
  rm protoc-${PROTOC_VER}-linux-x86_64.zip

CMD ["air"]

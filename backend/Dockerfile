FROM golang:1.22
WORKDIR /app
RUN go install github.com/cosmtrek/air@v1.51.0 && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

# protocのバージョンを指定
ARG PROTOC_VER=26.1

# protocコンパイラをダウンロードして解凍し、バイナリを/usr/local/binに移動
RUN apt-get update && apt-get install -y unzip imagemagick && \
  curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip && \
  unzip protoc-${PROTOC_VER}-linux-x86_64.zip -d /usr/local && \
  rm protoc-${PROTOC_VER}-linux-x86_64.zip

COPY policy.xml /etc/ImageMagick-6/policy.xml

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air"]

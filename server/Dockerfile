FROM appditto/imagemagick_go:latest as base

FROM base as builder

WORKDIR /root

# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and buil
# Install dependencies, imagemagick, go, and cleanup
RUN apt-get update && apt-get install -y \
    gcc pkg-config \
    && go get \
    && go build -o monkey-server \
    && rm -rf /var/lib/apt/lists/*

ENV ENVIRONMENT "production"

# run main.go
CMD ["./monkey-server", "-host=0.0.0.0", "-port=5555", "-logtostderr"]

FROM appditto/libvips:latest as builder

WORKDIR /root

# add source code
ADD . .
# add assets
ADD assets assets
# Build
RUN  go build -o /root/monkey

FROM ubuntu:22.04

ADD assets assets
COPY --from=builder /root/monkey /usr/local/bin/monkey

ENV MALLOC_ARENA_MAX=2
ENV GIN_MODE="release"

# Install libvips
RUN apt-get update && apt-get install -y \
    libvips-dev \
    && rm -rf /var/lib/apt/lists/*

# run main.go
CMD ["monkey", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
FROM appditto/libvips:latest

WORKDIR /root

# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and buil
# Install dependencies, libvips, go, and cleanup
RUN  go get \
    && go build -o monkey-server

ENV ENVIRONMENT "production"

# run main.go
CMD ["./monkey-server", "-host=0.0.0.0", "-port=5555", "-logtostderr"]

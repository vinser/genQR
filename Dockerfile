FROM golang as builder

# build genqr
COPY . /genqr
WORKDIR /genqr
RUN go build .

FROM alpine
COPY --from=builder /genqr/genqr /genqr/genqr

# run genqr with musl instead of glibc
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# configure directories
RUN mkdir -p /genqr/certs
COPY certs/server.* /genqr/certs

# expose ports
EXPOSE 8080

# volumes
VOLUME /certs

# # probes
# HEALTHCHECK --interval=30s --timeout=3s \
#   CMD wget --no-verbose --tries=1 --spider http://localhost:8080 || exit 1

# run command
ENTRYPOINT ["/genqr/genqr -port=8080 -cert=/certs/server.crt -key=/certs/server.key"]

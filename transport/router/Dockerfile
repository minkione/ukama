FROM alpine:3.14

COPY build/srvc_router /usr/bin/router

COPY build/lib /lib

COPY build/lib64 /lib64

ENTRYPOINT ["/usr/bin/router", "--host", "localhost", "--p", "8080", "--l", "debug"]


FROM golang:1.17-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata 

COPY ./bin/users /users
COPY ./cmd/users/sql /sql

FROM scratch AS final

# Import the time zone files
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# Import the CA certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled go executable

COPY --from=builder /users /users
COPY --from=builder /sql /sql

WORKDIR /

ENTRYPOINT ["/users"]

EXPOSE 8080
EXPOSE 8081
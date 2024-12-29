# syntax=docker/dockerfile:1
FROM golang:1.23 AS build
RUN apt install ca-certificates
ARG GEO_LITE_URL="https://s3.sspnet.tech/api/v1/buckets/public/objects/download?prefix=GeoLite2-City.mmdb"
ARG GEO_LITE_FILE_PATH="geolite2/GeoLite2-City.mmdb"
ADD build /exchange
WORKDIR /exchange/exchange
RUN curl -o GeoLite2-City.mmdb "${GEO_LITE_URL}"
RUN sed -i -e 's|@{build}|@{build_number}|g' -e 's|@{version}|@{version_number}|g' -e 's|@{commit}|@{commit_hash}|g' internal/main.go
RUN go generate ./...
RUN go build -o ../main internal/main.go

FROM ubuntu:latest
WORKDIR /exchange@{test_suffix}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /exchange/main .
COPY --from=build /exchange/GeoLite2-City.mmdb ./${GEO_LITE_FILE_PATH}
EXPOSE 8080
USER root:root
CMD ["./main"]

FROM golang:alpine3.13 as builder

WORKDIR /tmp/go
COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-w -s' -o location-history

FROM scratch
COPY --from=builder /tmp/go/location-history /location-history
VOLUME /data
EXPOSE 8080
ENTRYPOINT [ "/location-history" ]
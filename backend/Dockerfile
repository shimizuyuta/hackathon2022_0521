FROM golang:1.17 as builder

WORKDIR /go/src/backend/
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/backend


FROM scratch

COPY --from=builder /bin/backend /bin/backend
EXPOSE 8080

ENTRYPOINT [ "/bin/backend" ]
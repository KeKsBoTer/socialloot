
FROM golang:1.11 as builder
WORKDIR /server/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -ldflags="-s -w" -a -o socialloot .


FROM gcr.io/distroless/base
WORKDIR /root/
COPY conf/app.conf conf/app.conf
COPY static static
COPY views views
COPY --from=builder /server/socialloot .
ENTRYPOINT [ "./socialloot"]
EXPOSE 80
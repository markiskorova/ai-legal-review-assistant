FROM golang:1.22 as build
WORKDIR /src
COPY . /src
RUN go mod tidy && CGO_ENABLED=0 go build -o /out/worker ./cmd/worker

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /out/worker /worker
ENTRYPOINT ["/worker"]

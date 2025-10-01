FROM golang:1.22 as build
WORKDIR /src
COPY . /src
RUN go mod tidy && CGO_ENABLED=0 go build -o /out/api ./cmd/api

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /out/api /api
ENTRYPOINT ["/api"]

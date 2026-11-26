FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod main.go main_test.go ./
RUN go test ./... && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /api .

FROM scratch
COPY --from=build /api /api
USER 10001:10001
EXPOSE 8080
ENTRYPOINT ["/api"]

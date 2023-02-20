FROM golang:1.20-alpine as build

WORKDIR /opt/app

# cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/main ./lambdafunc/main.go

# copy artifacts to a clean image
FROM alpine

WORKDIR /

COPY --from=build /opt/app/bin/main /bin/main

ENTRYPOINT [ "/bin/main" ]  
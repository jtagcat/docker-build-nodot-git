# https://github.com/jtagcat/dotfiles/blob/main/scripts/template/gobuild.Dockerfile
# non-working workarounds: https://gist.github.com/jtagcat/189b2fd239687ab700f54faa46907df4

FROM golang:1.23 AS builder
WORKDIR /wd

COPY go.mod go.sum ./
RUN go mod download

# https://github.com/docker/docker.github.io/issues/14609
COPY *.go ./
#COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker-build-nodot-git

# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
LABEL org.opencontainers.image.source="https://github.com/jtagcat/docker-build-nodot-git"
WORKDIR /wd

COPY --from=builder /wd/docker-build-nodot-git ./
CMD ["./docker-build-nodot-git"]
EXPOSE 8080

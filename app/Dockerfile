FROM golang:1-buster

WORKDIR /go/src/app

# temporarily commented out until external library is added
# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . ./
RUN go install -v ./...

CMD ["app"]

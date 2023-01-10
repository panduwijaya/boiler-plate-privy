FROM golang:1.19

WORKDIR /go/src/cake-store/cake-store
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...


COPY go.mod .
COPY go.sum .
COPY . .


RUN go build -o cake-store

EXPOSE 8081

EXPOSE 3306 33060

ENTRYPOINT ./cake-store http

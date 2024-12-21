FROM golang:alpine

WORKDIR /project/go-docker/

COPY go.* ./

COPY . .

RUN go build -o /project/go-docker/build/goshort .

EXPOSE 8080
ENTRYPOINT [ "/project/go-docker/build/goshort" ]
FROM golang:1.13

WORKDIR /home/proxy/

COPY . /home/proxy/

RUN go build

CMD ./filterproxy
FROM golang:1.25.1

WORKDIR /go/src/github.com/ulvinamazow/studentApi
COPY . .
RUN go get -v 
RUN go build -o main .
CMD [ "./main" ]
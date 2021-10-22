FROM golang:alpine

MAINTAINER IDerr <ibrahim@derraz.fr>

WORKDIR /go

COPY . .

RUN go build -v -o bin/app src/app.go

EXPOSE 9144

CMD ["./src/app"]
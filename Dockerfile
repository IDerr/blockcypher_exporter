FROM golang:alpine

MAINTAINER IDerr <ibrahim@derraz.fr>

WORKDIR /app

COPY . .

RUN go build -v -o bin/app app.go

EXPOSE 9144

CMD ["./src/app"]
FROM golang:latest

WORKDIR /app

RUN go mod init rest_mysql

COPY . .

RUN go build -o main .

EXPOSE 8181

CMD [ "./main" ]


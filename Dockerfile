
FROM golang:1.17
 
WORKDIR /usr/local/go/bin/irctc-telegram-bot

ADD . /usr/local/go/bin/irctc-telegram-bot


RUN go build ./main.go

CMD ["./main"] 


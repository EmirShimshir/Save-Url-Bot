FROM golang:1.18

ENV TG_TOKEN=your-token

RUN mkdir -p /usr/src/app/
workdir /usr/src/app/

COPY . /usr/src/app/

RUN go build -o bot.exe main.go

CMD ["./bot.exe"]


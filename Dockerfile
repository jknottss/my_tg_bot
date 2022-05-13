FROM golang:latest
WORKDIR /my_tg_bot
COPY . .
RUN go mod download
RUN go build -o my_tg_bot ./cmd...
CMD ./my_tg_bot
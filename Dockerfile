FROM golang:latest

ENV GOPATH=/

COPY . .

RUN go mod download
RUN chmod +x start_servers.sh

CMD ./start_servers.sh

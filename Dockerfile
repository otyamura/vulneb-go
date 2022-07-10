FROM golang:1.18

ENV ROOT=/usr/src
WORKDIR ${ROOT}
RUN apt-get update && apt-get install -y \
  tldr
RUN mkdir -p /root/.local/share/tldr
RUN tldr -u
COPY go.mod go.sum ./
RUN go mod download

WORKDIR ${ROOT}/app

COPY . .
EXPOSE 8080
RUN chmod +x wait-for-it.sh

CMD ["./wait-for-it.sh", "db:3306", "--", "go", "run", "./cmd/vulneb-go/main.go"]

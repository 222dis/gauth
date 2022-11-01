FROM golang:1.19.2
# ENV GO111MODULE=on
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
# COPY go.sum .
RUN go mod download

# COPY . .

# RUN go build 

# CMD ["air", "-c", ".air.toml"]

# CMD air
# ENTRYPOINT ["air"]

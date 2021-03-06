FROM golang:1.14 as builder

#WORKDIR /app
RUN mkdir -p /go/src/gitlab.com/oneplanet/corona-backend/api
WORKDIR /go/src/gitlab.com/oneplanet/corona-backend/api

ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on

COPY api/go.mod api/go.sum ./
RUN go mod download

ADD api/ ./

ARG GIN_MODE
RUN GIN_MODE=${GIN_MODE} go build -o server -v main.go

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
FROM debian:stretch
EXPOSE 8080

RUN apt-get update && apt-get install -y ca-certificates --no-install-recommends && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /go/src/gitlab.com/oneplanet/corona-backend/api .

ARG GIN_MODE
ARG PORT
CMD GIN_MODE=${GIN_MODE} PORT=${PORT} ./server
#ENTRYPOINT ["./server"]
#CMD GIN_MODE=${GIN_MODE} go run -v main.go

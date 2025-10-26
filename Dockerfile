FROM golang:1.23-alpine AS build
LABEL authors="ilia"

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app ./server/server.go


FROM alpine:latest AS run

COPY --from=build /app /app

#WORKDIR /
EXPOSE 6969
CMD ["/app"]

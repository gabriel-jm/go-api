FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app-exe cmd/main.go

###

FROM scratch

WORKDIR /

COPY --from=build /app/app-exe /app-exe

EXPOSE 8000

ENTRYPOINT ["/app-exe"]

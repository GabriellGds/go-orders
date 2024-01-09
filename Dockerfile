###################
# Builder
###################
FROM golang:1.21.5-alpine3.17 as builder 

RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Create the final image, running the API and exposing port 5000
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
ARG PORT
ENV PORT=$PORT
EXPOSE $PORT

CMD ["./main"]

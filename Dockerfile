# ---- Build stage ----
FROM golang:1.21.4-alpine AS build

WORKDIR /src

# Cache dependencies first.
COPY go.mod ./
# COPY go.sum ./   # add back once external deps introduce a go.sum
RUN go mod download

COPY . .

# Static, stripped binary. Assets are embedded via go:embed.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" \
    -o /out/portfolio ./cmd/portfolio

# ---- Runtime stage ----
FROM alpine:3.19

RUN apk add --no-cache ca-certificates wget \
    && addgroup -S app && adduser -S app -G app

WORKDIR /app
COPY --from=build /out/portfolio /app/portfolio

USER app

ENV PORT=8080
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget -qO- http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["/app/portfolio"]

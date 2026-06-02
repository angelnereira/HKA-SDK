# Build the HKA gateway as a small static image.
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/hka-gateway ./cmd/hka-gateway

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/hka-gateway /hka-gateway
EXPOSE 8080
ENV ADDR=:8080
# HKA_ENDPOINT defaults to the demo endpoint when unset.
ENTRYPOINT ["/hka-gateway"]

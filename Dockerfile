# Stage 1 - build
FROM golang:1.26.7-bookworm AS build
WORKDIR /build
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/exe ./cmd/app/main.go

# Stage 2 - final
FROM alpine:3.20
WORKDIR /Smarttrain
COPY --from=build /build/exe .
CMD ["/Smarttrain/exe"]
# syntax=docker/dockerfile:1

# If you need more help, visit the Dockerfile reference guide at
# https://docs.docker.com/go/dockerfile-reference/

################################################################################

FROM golang:1.25.1 AS build
WORKDIR /main
#WORKDIR /app
#cahe go dep
COPY go.mod go.sum ./
RUN go mod download 




COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o server .



FROM gcr.io/distroless/base-debian12
#WORKDIR /main


USER nonroot:nonroot
COPY --from=build /main/server /main/server
#COPY --from=build /app/server /app/server
# Expose the port that the application listens on.
EXPOSE 80

CMD [ "/main/server" ]

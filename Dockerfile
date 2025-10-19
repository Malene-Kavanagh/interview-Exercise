# syntax=docker/dockerfile:1

################################################################################

FROM golang:1.25.1 AS build
WORKDIR /main
#WORKDIR /app
#cahe go dep
COPY go.mod go.sum ./
RUN go mod download 
#copy rest of the files

COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o server .
#output binary file named server

FROM gcr.io/distroless/base-debian12
#debian12 base image 

USER nonroot:nonroot
COPY --from=build /main/server /main/server
#copy the binary file from the build stage

# Expose the port that the application listens on.
EXPOSE 80

CMD [ "/main/server" ]
# Run the binary.
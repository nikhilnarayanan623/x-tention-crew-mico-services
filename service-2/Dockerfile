FROM golang:1.21.1-alpine3.18 AS build-stage

WORKDIR /app

# copy go.mod and go.sum
COPY go.mod go.sum ./

#dwoload dependency
RUN go mod download

# copy the source code
COPY cmd cmd/
COPY pkg pkg/

# build the exec file for service
RUN go build -v -o ./build/bin/api ./cmd/api

# final stage for packe the build file into a build executable image
FROM gcr.io/distroless/static-debian11

# copy build file from build stage to fnal stage
COPY --from=build-stage /app/build/bin/api api

CMD ["/api"]
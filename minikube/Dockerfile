
# http://www.matthiassommer.it/programming/compile-and-run-golang-executable-with-docker/


# FROM golang:alpine AS builder
# # Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git
# WORKDIR $GOPATH/src/pfistera/minikube
# COPY . .
# # Fetch dependencies.
# # Using go get.
# RUN go get -d -v
# # Build the binary.
# #RUN go build -o /go/bin/hello
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello
# ############################
# # STEP 2 build a small image
# ############################
# FROM scratch
# # Copy our static executable.
# COPY --from=builder /go/bin/hello /go/bin/hello
# # Run the hello binary.
# ENTRYPOINT ["/go/bin/hello"]



FROM golang:1.15.3-alpine3.12 AS builder

LABEL maintainer="rosso0815"

WORKDIR /go/src/github.com/rosso0815/minikube

RUN apk add --update git
# RUN go get gopkg.in/alecthomas/kingpin.v2

# copy the code and data
#COPY data/example_input.csv ./data/
COPY main.go .

# build the app
RUN GOOS=linux GOARCH=amd64 go build -o /go/app 

# Second stage
FROM alpine

#COPY --from=builder /go/src/github.com/rosso0815/minikube/ .
COPY --from=builder /go/app /go/app

# ENV INPUT ""
# ENV OUTPUT ""

# CMD ./app $INPUT $OUTPUT
#and the docker build and docker run command.
#docker build --rm -t csv2gpx .
#docker run -e INPUT=data/example_input.csv -e OUTPUT=data/example_input.gpx csv2gpx




# FROM golang:alpine AS builder

# WORKDIR $GOPATH/src/mylekkepackage/mylekkeapp/

# COPY run.go .

# RUN go build -o /go/app

# FROM scratch
# COPY --from=builder /go/app /go/app
EXPOSE 3000
CMD ["/go/app"]
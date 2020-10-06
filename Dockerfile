FROM golang:alpine as builder

WORKDIR '/build'
COPY . .
RUN apk update && apk add bash && apk add git && apk add build-base
RUN /bin/bash -c "echo GIT_TERMINAL_PROMPT=1 >> /etc/environment"
RUN /bin/bash -c "GIT_TERMINAL_PROMPT=1 GOSUMDB=off go get ./..."
RUN /bin/bash -c "go test -cover ./..."
RUN /bin/bash -c "cd cmd/app && go build"

FROM alpine:latest
RUN mkdir /server
RUN apk update && apk add bash
COPY --from=builder /build/cmd/app/app ./
CMD ./app && /bin/bash
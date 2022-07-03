# build stage
FROM golang:1.18-alpine AS dev

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ADD ../ /src
RUN cd /src && go build -o myretail cmd/main.go

# final stage
FROM alpine as prod
WORKDIR /app
COPY --from=dev /src/myretail /app/
ENTRYPOINT ./myretail

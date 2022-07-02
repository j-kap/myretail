# build stage
FROM golang:1.18-alpine AS build
ADD ../ /src
RUN cd /src && go build -o myretail cmd/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build /src/myretail /app/
ENTRYPOINT ./myretail

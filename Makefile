SHELL := /bin/bash

build:
	@ go build -o myretail cmd/main.go

run:
	@ source .env && export PRODUCTS_URL && go run cmd/main.go

docker/build:
	@ docker build . -f docker/Dockerfile -t myretail

docker/run:
	@ docker run -p 8000:8080 -v "${GOOGLE_APPLICATION_CREDENTIALS}":/gcp/creds.json:ro --env GOOGLE_APPLICATION_CREDENTIALS=/gcp/creds.json --env-file .env myretail

docker/build_run: docker/build docker/run

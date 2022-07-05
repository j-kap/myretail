SHELL := /bin/bash

build:
	@ go build -o myretail cmd/main.go

run:
	@ source .env && export PRODUCTS_URL && go run cmd/main.go

docker/build:
	@ docker build . -t myretail

docker/run:
	@ docker run -p 8000:8000 -v "${GOOGLE_APPLICATION_CREDENTIALS}":/gcp/creds.json:ro --env GOOGLE_APPLICATION_CREDENTIALS=/gcp/creds.json --env-file .env myretail

docker/build_run: docker/build docker/run

docker/compose:
	@ docker compose up --build

docker/tests:
	@ docker compose --profile tests up --build --abort-on-container-exit --exit-code-from=app-tests --attach app-tests

build:
	@ go build -o myretail cmd/main.go

run:
	@ . .env; \
	  export PRODUCTS_URL PROJECT_ID PORT; \
	  go run cmd/main.go

docker/build:
	@ docker build . -t myretail

docker/run:
	@ docker run -p 8000:8000 -v "${GOOGLE_APPLICATION_CREDENTIALS}":/gcp/creds.json:ro \
	    --env GOOGLE_APPLICATION_CREDENTIALS=/gcp/creds.json --env-file .env myretail

docker/build_run: docker/build docker/run

docker/compose:
	@ docker compose up --build

docker/tests:
	@ docker compose --profile tests up --build --abort-on-container-exit --exit-code-from=app-tests --attach app-tests

gcloud/tag:
	@ . .env; \
	  export PROJECT_ID; \
	  docker tag myretail gcr.io/$${PROJECT_ID}/myretail:latest

gcloud/push:
	@ . .env; \
	  export PROJECT_ID; \
	  docker push gcr.io/$${PROJECT_ID}/myretail:latest

gcloud/deploy:
	@ . .env; \
	  export PRODUCTS_URL PROJECT_ID; \
	  gcloud run deploy myretail --image gcr.io/$${PROJECT_ID}/myretail --region us-central1 --allow-unauthenticated \
	    --set-env-vars PRODUCTS_URL=$${PRODUCTS_URL} --set-env-vars PROJECT_ID=$${PROJECT_ID}

gcloud/build_deploy: docker/build gcloud/tag gcloud/push gcloud/deploy

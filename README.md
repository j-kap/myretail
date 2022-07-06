# myRetail RESTful service

myRetail is a microservice written in Go utilizing Google Cloud's Firestore NoSQL database. It retrieves product details from an external HTTP API and reads the pricing information, if available, from Firestore. It can also add/update the price information in Firestore. 


## How to Use

myRetail supports one endpoint, `/products/<PRODUCT_ID>`, with the verbs `GET` and `PUT`.

`GET` returns the product name and current price, if available. 

`PUT` can be used to set/update the current price data.

### GET Product
```
$ http localhost:8000/products/13860428
HTTP/1.1 200 OK
Content-Length: 115
Content-Type: application/json; charset=utf-8
Date: Wed, 06 Jul 2022 15:33:07 GMT

{
    "current_price": {
        "currency_code": "USD", 
        "value": "13.99"
    }, 
    "id": "13860428", 
    "name": "The Big Lebowski (Blu-ray)"
}
```

### PUT Product
```
$ http PUT localhost:8000/products/13860428 < price.json 

HTTP/1.1 200 OK
Content-Length: 72
Date: Wed, 06 Jul 2022 14:51:17 GMT
content-type: application/json; charset=utf-8

{
    "current_price": {
        "currency_code": "USD", 
        "value": "1.99"
    }, 
    "id": "13860428"
}
```

price.json contents:
```
{
    "id": "13860428",
    "current_price": {
        "currency_code": "USD", 
        "value": "1.99"
    }
}
```

## How to Run

This repo contains a Makefile that can be used to build and run the app, execute integration tests, and deploy to Google Cloud Run. The targets are detailed below. Run a given target with the command `make <target_name>`. Example: `make terraform/build_deploy`

Most of the make targets expect you to have a file named `.env` in the root of the project directory with some required environment variables defined. Just copy the `env-example` to `.env` and modify its contents accordingly.


### Makefile targets


#### `build`
Simply builds the myRetail service and outputs the binary named `myretail`.

#### `run`
Runs the myRetail service locally using a real Google Firestore instance. Requires your GCP Project ID to be set in `.env` and your GCP project to be initialized (see make target `gcloud/init`).

#### `docker/build`
Builds the myRetail service docker image tagged `myretail`.

#### `docker/run`
Runs the myRetail docker image created in the previous target. Requires the environment variable `GOOGLE_APPLICATION_CREDENTIALS` to be set with the path to a GCP service account json file.

#### `docker/build_run`
Executes both make targets `docker/build` and `docker/run`.

#### `docker/compose`
Runs the myRetail service in docker compose utilizing a local Firestore emulator docker container. This does not require a real GCP Project to be configured.

#### `docker/tests`
Runs the myRetail service integration tests in docker compose, also utilizling the local Firestore emulator. Note: this command only attaches to the docker logs for the container running the integration tests. To view application logs for debugging purposes you can run `docker logs myretail-app-1` in a separate terminal window.

#### `gcloud/init`
Initializes the GCP project specified in your `.env` file. This is only intended to be run once on a new GCP Project that already has billing enabled. It enables the various GCP APIs needed to deploy the application and creates the Firestore database. This setup is required on any project intended to be used for any of the `gcloud/` or `terraform/` targets that follow. Note: this setup was not handled in terraform due to GCP limitations that break idempotency, specifically around deletes.

#### `gcloud/push`
Tags and pushes the previously created `myretail` docker image to Container Registry in your GCP project.

#### `gcloud/deploy`
Deploys the previously pushed docker image to Google Cloud Run.

#### `gcloud/build_push`
Runs `docker/build` and `gcloud/push`.

#### `gcloud/build_deploy`
Runs `docker/build`, `gcloud/push`, and `gcloud/deploy`.

#### `terraform/init`
Initializes the terraform directory with the required providers. This must be run once before running the other `terraform/` targets.

#### `terraform/plan`
Displays the resources terraform will create in GCP and creates a plan file to be used in `terraform/apply`.

#### `terraform/apply`
Applies the terraform plan created with `terraform/plan`. This will deploy the myRetail service and output its unique service URL. This requires the myRetail docker image to already exist in GCP (see `terraform/build_deploy`).

#### `terraform/destroy`
Destroys the previously applied resources.

#### `terraform/deploy`
Runs `terraform/plan` and `terraform/apply`.

#### `terraform/build_deploy`
Runs `gcloud/build_push` and `terraform/apply`.

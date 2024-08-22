# {{bootstrap_template}}

#### Run Project

- Starting API in development mode

```sh
go run main.go
```

- Run tests local

```sh
go test ./... -coverprofile cover.out
```

- Run coverage report

```sh
go tool cover -html=cover.out
```

- Create mocks with mockery

```sh
mockery --all --keeptree 

```

### Local Env

````
APP_NAME={{APP_NAME}}
APP_PORT={{APP_PORT}}
GOOGLE_APPLICATION_CREDENTIALS=
FIRESTORE_COLLECTION=
FIRESTORE_PATH=
PROJECT_ID=

POSTGRES_HOST=
POSTGRES_DATABASE=
POSTGRES_PORT=
POSTGRES_DRIVER=
POSTGRES_USER=
POSTGRES_PASSWORD=

INBOUND_SUBSCRIPTION=
OUTBOUND_TOPIC=



````



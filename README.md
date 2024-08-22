# Golang template scaffolding

### Sumary
This project simple scaffolding for create templates of golang project

### Supported GCP Services
- Pub/Sub
- Firestore

### Supported Technologies
- Redis
- Gorilla Mux
- Logrus
- Postgres
- Kustomize K8S

### Features
- Uses clean architecture 
- Bootstrap initializer
- You can add or remove functionalities, depending your requirements
- Generate a copy of template in other folder


### Project Structure
````
├── README.md
├── go.mod
├── go.sum
├── main.go
└── template
    ├── Dockerfile
    ├── README.md
    ├── application
    │   └── adapter
    │       └── api
    │           └── health
    │               ├── health_check.go
    │               └── health_check_test.go
    ├── kustomization
    │   ├── base
    │   │   ├── deployment.yaml
    │   │   ├── kustomization.yaml
    │   │   └── service.yaml
    │   ├── development
    │   │   ├── app-container-config.yaml
    │   │   ├── autoscaling.yaml
    │   │   ├── env
    │   │   ├── ingress.yaml
    │   │   ├── kustomization.yaml
    │   │   └── workload-identity.yaml
    │   ├── production
    │   │   ├── app-container-config.yaml
    │   │   ├── autoscaling.yaml
    │   │   ├── env
    │   │   ├── ingress.yaml
    │   │   ├── kustomization.yaml
    │   │   └── workload-identity.yaml
    │   ├── test
    │   │   ├── app-container-config.yaml
    │   │   ├── autoscaling.yaml
    │   │   ├── env
    │   │   ├── ingress.yaml
    │   │   ├── kustomization.yaml
    │   │   └── workload-identity.yaml
    │   └── uat
    │       ├── app-container-config.yaml
    │       ├── autoscaling.yaml
    │       ├── env
    │       ├── ingress.yaml
    │       ├── kustomization.yaml
    │       └── workload-identity.yaml
    ├── main.go
    └── pkg
        ├── bootstrap.go
        ├── cache
        │   ├── factory_redis.go
        │   ├── redis_client.go
        │   └── redisrepo
        │       └── repository.go
        ├── database
        │   ├── config.go
        │   ├── factory_nosql.go
        │   ├── factory_sql.go
        │   ├── firestore_handler.go
        │   ├── nosql
        │   │   └── nosql.go
        │   └── postgres_handler.go
        ├── log
        │   ├── factory.go
        │   ├── logger
        │   │   └── logger.go
        │   └── logrus.go
        ├── publisher
        │   ├── config.go
        │   └── publisher_handler.go
        ├── rest
        │   ├── config.go
        │   ├── factory.go
        │   └── rest-client.go
        ├── router
        │   ├── factory.go
        │   ├── gorilla_mux.go
        │   └── sanitize_middleware.go
        ├── subscriber
        │   ├── config.go
        │   ├── event_handler.go
        │   └── factory.go
        └── utils
            └── array_utils.go


````



### Usage

Please, build the project
````
go build .
````

Options or flags
````
Flags:
  -a, --author string    Name of team or author(required)
  -h, --help             help for template-cli
  -m, --modname string   Go module name to replace (required)
  -d, --target string    Target directory path (required)

````
Example usage:
````
 ./template-go-cli -a mrch -m template -d  ./template1 
````
Now you can use this project as possible in another folder.


Enjoy :)

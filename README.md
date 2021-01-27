# Monitoring Configuration Service

Run the service inside Docker container:

```bash
$ make dockerize
app_1  | ts=2021-01-23T17:16:59.115450377Z caller=main.go:49 transport=http address=:8080 msg=listening
```

Get the configuration:

```bash
$ curl http://localhost:8080/configurations
{"configuration":[{"domains":[{"site_id":1,"name":"domain1.tld","plugins":[{"id":1,"namespace": "http","name":"http_status200"}]}]}]}
```

## How to seed a database

```bash
MONITORING_CONFIGURATION_SERVICE_URL=http://localhost:8080 OUTPUT_DIR=./public MONITORING_CONFIGURATION_DB_SEED=1 go run main.go
```

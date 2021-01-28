# Monitoring Configuration Service

Start database:

```bash
docker-compose up db
```

Start web-server:

```bash
$ make run
```

Get the configuration:

```bash
$ curl http://localhost:8080/configurations | jq
```

```json
{
  "configuration": {
    "domains": [
      {
        "site_id": 1,
        "name": "domain1.tld",
        "plugins": [
          {
            "id": 1,
            "namespace": "content",
            "name": "contains_string",
            "settings": {
              "value": "content1",
              "resource": "/resource1"
            }
          },
          {
            "id": 2,
            "namespace": "content",
            "name": "does_not_contain_string",
            "settings": {
              "value": "content2",
              "resource": "/resource2"
            }
          },
          {
            "id": 3,
            "namespace": "content",
            "name": "valid_json",
            "settings": {}
          },
          {
            "id": 10,
            "namespace": "files",
            "name": "robots_txt",
            "settings": {}
          },
          {
            "id": 11,
            "namespace": "files",
            "name": "sitemap_xml",
            "settings": {}
          },
          {
            "id": 12,
            "namespace": "http",
            "name": "http_status200",
            "settings": {}
          },
          {
            "id": 13,
            "namespace": "http",
            "name": "http_to_https_redirect",
            "settings": {}
          },
          {
            "id": 14,
            "namespace": "http",
            "name": "non_existent_url_returns404",
            "settings": {}
          },
          {
            "id": 15,
            "namespace": "http",
            "name": "valid_http_status_code",
            "settings": {
              "value": 301,
              "resource": "/resource"
            }
          },
          {
            "id": 16,
            "namespace": "http",
            "name": "www_to_non_www_redirect",
            "settings": {}
          },
          {
            "id": 17,
            "namespace": "other",
            "name": "database_connection_issue",
            "settings": {}
          },
          {
            "id": 18,
            "namespace": "ownership",
            "name": "dns_txt_record_verification",
            "settings": {}
          },
          {
            "id": 19,
            "namespace": "ownership",
            "name": "file_verification",
            "settings": {}
          }
        ]
      },
      {
        "site_id": 2,
        "name": "domain2.tld",
        "plugins": [
          {
            "id": 4,
            "namespace": "domain",
            "name": "dns_a_records",
            "settings": {}
          },
          {
            "id": 5,
            "namespace": "domain",
            "name": "domain_expiration",
            "settings": {}
          },
          {
            "id": 6,
            "namespace": "domain",
            "name": "ssl_certificate_expiration",
            "settings": {}
          }
        ]
      }
    ]
  }
}
```

## How to seed a database

```bash
MONITORING_CONFIGURATION_DB_SEED=1 go run main.go
```

## How to clean a database

```bash
rm -rf .data
```

## Tests

```bash
make test
```

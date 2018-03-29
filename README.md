# prom-dynamodb-storage
External storage gateway for prometheus - NOT READY FOR PROD

[![Build Status](https://travis-ci.org/Labbs/prom-dynamodb-storage.svg?branch=master)](https://travis-ci.org/Labbs/prom-dynamodb-storage) [![Go Report Card](https://goreportcard.com/badge/github.com/Labbs/prom-dynamodb-storage)](https://goreportcard.com/report/github.com/Labbs/prom-dynamodb-storage)

## Options
```
--listen_port value, --lp value    (default: 1234) [$LISTEN_PORT]
--dynamo_local, --dl                [$DYNAMO_LOCAL]
--dynamo_table value, --dt value   (default: "prometheus") [$DYNAMO_TABLE]
--dynamo_region value, --dr value  (default: "eu-west-1") [$DYNAMO_REGION]
```

## Information
AWS session use the current credentials if set.

## Features
* Bulk write to dynamodb
* Read path
* Metrics path
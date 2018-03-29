# prom-dynamodb-storage
External storage gateway for prometheus

## Options
```
--listen_port value, --lp value    (default: 1234) [$LISTEN_PORT]
--dynamo_local, --dl                [$DYNAMO_LOCAL]
--dynamo_table value, --dt value   (default: "prometheus") [$DYNAMO_TABLE]
--dynamo_region value, --dr value  (default: "eu-west-1") [$DYNAMO_REGION]
```

## Information
AWS session use the current credentials if set.
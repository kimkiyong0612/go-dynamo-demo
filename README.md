# go-dynamo-demo

### create table
```
$ aws dynamodb create-table --endpoint-url http://localhost:8000 \
    --table-name MyFirstTable \
    --attribute-definitions \
        AttributeName=MyHashKey,AttributeType=S \
        AttributeName=MyRangeKey,AttributeType=N \
    --key-schema \
        AttributeName=MyHashKey,KeyType=HASH \
        AttributeName=MyRangeKey,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=1,WriteCapacityUnits=1
```

### delete table
```
$ aws dynamodb delete-table --endpoint-url http://localhost:8000 \
    --table-name MyFirstTable
```
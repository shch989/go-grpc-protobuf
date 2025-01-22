```
protoc --go_out=. ./proto/basic/*.proto
```
```
protoc --go_opt=module=my-protobuf --go_out=. ./proto/basic/*.proto
```
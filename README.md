```
protoc --go_out=. ./proto/basic/*.proto
```
```
protoc --go_opt=module=my-protobuf --go_out=. ./proto/basic/*.proto
```
```
go mod tidy
```
```
go mod vendor
```

```
protoc --go_opt=module=my-protobuf --go_out=. ./proto/basic/*.proto ./proto/dummy/*.proto ./proto/jobsearch/*.proto
```

```
protoc --validate_out="lang=go:./generated" --go_opt=module=my-protobuf --go_out=. ./proto/car/*.proto
```
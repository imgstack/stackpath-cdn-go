# StackPath CDN Go Client

Go client for the StackPath CDN OpenAPI spec at https://developer.stackpath.com/specs/cdn.swagger.json.

## How to regenerate the bindings
1. Install https://github.com/go-swagger/go-swagger.

2. Generate Go files
```
swagger generate client -f cdn.swagger.json -A stackpath-cdn-go
```

3. Run `go get -u -f ./...`
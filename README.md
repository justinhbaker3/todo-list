# todo-list

Uses Gin (https://gin-gonic.com) web framework

### Build and run
```
make build
./todo-list
```

### Healthcheck
```
curl -X GET localhost:8080/ping
```

### Example Request
```
curl -X POST localhost:8080/list -d '{"Title": "First_List"}'
```

### Dependencies
```
go get <dependency>
go mod vendor
```

### TODO
1. Implement delete item functionality
2. Implement delete list functionality
3. Implement timed lists where a list will log something when it's
time is up
4. Unit tests

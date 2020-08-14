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
4. Implement delete functionality in the mysql store
5. Unit tests

### Git
```
git clone https://github.com/justinhbaker3/todo-list.git
git checkout -b "add_to_readme"
git status
git add .
git commit -m "updating readme"
git push
```

### Docker
Create image:
`make image` or `docker build -t justinhbaker3/todo-list .`

Run container:
`make docker-run` or `docker run -p 8080:8080 justinhbaker3/todo-list`

### MYSQL
Build image:
`make build-mysql`

Start container:
`make run-mysql`

Start a bash shell in the mysql container:
`docker exec -it todo-list-mysql bash`

Alternatively, install mysql and run
`mysql --host=127.0.0.1 -u root --password=pass --protocol=TCP` to connect from outside the container

From within the mysql container (enter `pass` when prompted):
`mysql --user root --password=pass`

Test commands to run:

`use main;`
`show tables;`
`describe lists;`

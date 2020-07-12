NAMESPACE = justinhbaker3
APPNAME = todo-list

.PHONY: build
build:
	go build -mod vendor github.com/$(NAMESPACE)/$(APPNAME)/cmd/todo-list

.PHONY: clean
clean:
	-rm todo-list

.PHONY: image
image:
	docker build -t $(NAMESPACE)/$(APPNAME) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(NAMESPACE)/$(APPNAME)

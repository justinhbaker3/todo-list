NAMESPACE = justinhbaker3
APPNAME = todo-list

.PHONY: build
build:
	go build -mod vendor github.com/$(NAMESPACE)/$(APPNAME)/cmd/todo-list

.PHONY: clean
clean:
	rm todo-list

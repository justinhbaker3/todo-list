FROM golang:1.14
WORKDIR /code
COPY . .
RUN make clean && make build

EXPOSE 8080
ENTRYPOINT ["./todo-list"]
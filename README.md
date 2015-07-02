#### Golang use protobuf

##### Linux or Mac install protobuf compiler [download](https://github.com/google/protobuf)

When you installed, Like this

```
~/code/golang/src/go-protobuf-demo on ⭠ master ⌚ 22:12:20
$ protoc --version
libprotoc 2.6.1

```

##### Download the go-protobuf client [url](https://github.com/golang/protobuf)

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

##### Compiler your `*.proto` file to `.go` 

```
protoc --go_out=./pb-go ./pb-proto
```
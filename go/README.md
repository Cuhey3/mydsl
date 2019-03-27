## mydsl

mydsl is yaml-based DSL library for JavaScript/Node.js/Go.

mydsl can replace your code with a YAML file.


### Library Testing
Testing is also written in YAML DSL and you can use it as a DSL reference.

#### execute test for go

```
go test github.com/cuhey3/mydsl/go -coverprofile=$GOPATH/src/github.com/cuhey3/mydsl/examples/public/cover.out
```

or

```
go test $GOPATH/src/mydsl/go -coverprofile=$GOPATH/src/mydsl/examples/public/cover.out
```

#### convert .out to .html
```
go tool cover -html=$GOPATH/src/github.com/cuhey3/mydsl/examples/public/cover.out -o $GOPATH/src/github.com/cuhey3/mydsl/examples/public/cover.html
```

#### start server to see coverage html
```
go run examples/server.go
```

#### access coverage html page
```
http(s)://your-host/public/cover.html
```

# list direct deps
```shell
$ go list -m -f '{{if not .Indirect}}{{.Path}}{{end}}' all 
```

# compile
```shell
$ cd src
$ ./all.bash
```
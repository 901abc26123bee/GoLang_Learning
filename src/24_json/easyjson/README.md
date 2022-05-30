## 更快的 JSON 解析
* EasyJSON 采用代码生成而非反射
[easyjson](https://github.com/mailru/easyjson)
```sh
# install, for Go < 1.17
$ go get -u github.com/mailru/easyjson/...

# for Go >= 1.17
go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest

# use
$ easyjson -all <file>.go

$ cd src/24_json/easyjson
$ easyjson -all struct_def.go
# or
$ ~/go/bin/easyjson  -all struct_def.go
# The above will generate <file>_easyjson.go containing the appropriate marshaler and unmarshaler funcs for all structs contained in <file>.go.
```
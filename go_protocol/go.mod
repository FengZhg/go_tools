module github.com/FengZhg/go_tools/go_protocol

go 1.18

require (
	github.com/FengZhg/go_tools/errs v0.0.0-20220328073327-4f127cbd2444
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/golang/protobuf v1.5.2
)

require google.golang.org/protobuf v1.28.0 // indirect

replace github.com/FengZhg/go_tools/errs => ../errs

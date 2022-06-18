module example.com/hello

go 1.18

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

module github.com/plainboring/config_client

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/pingcap/errors v0.11.4
	github.com/pingcap/kvproto v0.0.0-20191023083849-24285ce8b84a
	github.com/pingcap/log v0.0.0-20190715063458-479153f07ebd
	github.com/pingcap/pd v1.1.0-beta.0.20191023071330-85b46b6623a3
	github.com/spf13/cobra v0.0.5
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64 // indirect
	google.golang.org/grpc v1.23.0
)

// replace github.com/pingcap/kvproto => github.com/plainboring/kvproto master
replace github.com/pingcap/kvproto => github.com/plainboring/kvproto v0.0.0-20191026132945-c4450272525d

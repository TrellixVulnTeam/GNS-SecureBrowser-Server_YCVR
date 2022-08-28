module gnsdeviceserver

go 1.17

require go.bug.st/serial v1.3.4

require (
	github.com/kardianos/service v1.2.1
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/creack/goselect v0.1.2 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/mitchellh/panicwrap v1.0.0 // indirect
	github.com/pkg/profile v1.6.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

replace go.bug.st/serial => ../go-serial

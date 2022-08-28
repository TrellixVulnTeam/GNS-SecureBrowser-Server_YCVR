package main

import (
	"context"
	"flag"
	"fmt"
	"gnsdeviceserver/proto/gnsrpc"
	"log"
	"net"
	"time"

	"github.com/kardianos/service"
	"google.golang.org/grpc"
)

var (
	Port = flag.Int("port", 50051, "The server port")
)

// mock server to implement GNSBadgeData Service
type mockServer struct {
	gnsrpc.UnimplementedGNSBadgeDataServer
}

func (s *mockServer) ReadUUID(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.UUID, error) {
	return &gnsrpc.UUID{Uuid: "12343243443243"}, nil
}

func (s *mockServer) FormatCard(ctx context.Context, in *gnsrpc.UUID) (*gnsrpc.GNSBadgeDataParam, error) {
	return &gnsrpc.GNSBadgeDataParam{}, nil
}

func (s *mockServer) GetFreeSites(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.FreeSites, error) {
	//free_sites := []uint32{1, 2, 3, 4, 5}
	return &gnsrpc.FreeSites{Idx: []uint32{1, 2, 3, 4, 5}}, nil
}

func (s *mockServer) GetFreeWinCreds(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.FreeWinCreds, error) {
	return &gnsrpc.FreeWinCreds{Idx: []uint32{1, 2, 3, 4, 5}}, nil
}
func (s *mockServer) ReadSiteCreds(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.Sites, error) {
	sites := []*gnsrpc.SiteCred{{Idx: 1, Username: "Peter", Password: "pass1", Code: "1AF"},
		{Idx: 2, Username: "Peter2", Password: "pass2", Code: "2B"},
		{Idx: 2, Username: "Peter3", Password: "pass3", Code: "31"}}
	return &gnsrpc.Sites{Sites: sites}, nil
}
func (s *mockServer) ReadWinCreds(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.WinCreds, error) {
	wincreds := []*gnsrpc.WinCred{{Idx: 1, Username: "Peter", Password: "pass1", Domain: "1AF"},
		{Idx: 2, Username: "Peter2", Password: "pass2", Domain: "2B"},
		{Idx: 2, Username: "Peter3", Password: "pass3", Domain: "31"}}

	return &gnsrpc.WinCreds{Wincreds: wincreds}, nil
}
func (s *mockServer) DeleteSiteCred(ctx context.Context, in *gnsrpc.SiteCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Printf("Receive Delete SiteCred username: %s password: %s\n", in.Username, in.Password)
	return &gnsrpc.GNSBadgeDataParam{}, nil
}
func (s *mockServer) DeleteWinCred(ctx context.Context, in *gnsrpc.WinCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Printf("Receive Delete Wincred username: %s password: %s\n", in.Username, in.Password)
	return &gnsrpc.GNSBadgeDataParam{}, nil
}
func (s *mockServer) WriteSiteCred(ctx context.Context, in *gnsrpc.SiteCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Printf("Receive Write SiteCred username: %s password: %s\n", in.Username, in.Password)
	return &gnsrpc.GNSBadgeDataParam{}, nil
}
func (s *mockServer) WriteWinCred(ctx context.Context, in *gnsrpc.WinCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Printf("Receive Write Wincred username: %s password: %s\n", in.Username, in.Password)
	return &gnsrpc.GNSBadgeDataParam{}, nil
}

func (s *mockServer) Execute(ctx context.Context, in *gnsrpc.Text) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Printf("Received Execute command %s\n", in.Text)
	return &gnsrpc.GNSBadgeDataParam{}, nil
}

func (s *mockServer) StreamCardStatus(param *gnsrpc.GNSBadgeDataParam, srv gnsrpc.GNSBadgeData_StreamCardStatusServer) error {
	ctx := srv.Context()
	var startTS time.Time
	startTS = time.Now()
	log.Println("We got a new client")
	//var toggle int
	//toggle = 0
	for {
		select {
		case <-ctx.Done():
			log.Println("stream context done")
			return ctx.Err()
		default:
		}
		if time.Since(startTS) > 2*time.Second {
			startTS = time.Now()
			result := gnsrpc.CardStatus{}

			result.Type = gnsrpc.CardStatus_USB

			/*if toggle == 1 {
				result.Status = gnsrpc.CardStatus_Connected
				toggle = 0
			} else {
				result.Status = gnsrpc.CardStatus_Disconnected
				toggle = 1
			}*/
			result.Status = gnsrpc.CardStatus_Authenticated

			if err := srv.Send(&result); err != nil {
				log.Printf("send error %v", err)
			}
		}
	}
}

type program struct{}

var web_host *grpc.Server

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	fmt.Println("Setting up GNS GRPC server")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	web_host = grpc.NewServer()

	gnsrpc.RegisterGNSBadgeDataServer(web_host, &mockServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := web_host.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	web_host.Stop()
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GNSBadgeService",
		DisplayName: "GNSBadge Service",
		Description: "GNS Badge Data server",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

}

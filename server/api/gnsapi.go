package gnsapi

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gnsdeviceserver/proto/gnsrpc"
	gnsserial "gnsdeviceserver/serial"
	"log"
	"strings"
	"time"
)

var (
	Port = flag.Int("port", 50051, "The server port")
)

// server is used to implement GNS Badge Service
type Server struct {
	gnsrpc.UnimplementedGNSBadgeDataServer
}

func (s *Server) ReadUUID(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.UUID, error) {
	log.Println("Received ReadUUID request")
	uuid, err := gnsserial.ReadUUID()
	if err != nil {
		log.Println(err.Error())
	}
	return &gnsrpc.UUID{Uuid: uuid}, err

}

// return UUID from zone 2
func (s *Server) ReadUUIDZone2(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.UUID, error) {
	uuid, err := gnsserial.ReadUUIDZone2()
	if err != nil {
		return nil, err
	}
	var result gnsrpc.UUID
	result.Uuid = uuid

	return &result, nil
}

// Read HW UUID then store to Zone 2
func (s *Server) StoreUUID(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.GNSBadgeDataParam, error) {
	err := gnsserial.StoreUUID()
	return &gnsrpc.GNSBadgeDataParam{}, err
}

func (s *Server) GetFreeSites(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.FreeSites, error) {
	log.Println("Received GetFreeSites request")
	free_sites, err := gnsserial.GetFreeSites()
	if err != nil {
		log.Println(err.Error())
	}
	return free_sites, err
}

func (s *Server) FormatCard(ctx context.Context, in *gnsrpc.UUID) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Println("Received FormatCard request")
	err := gnsserial.FormatCard(int(in.Mode), in.Uuid)
	if err != nil {
		log.Println(err.Error())
	}
	return &gnsrpc.GNSBadgeDataParam{}, err
}

func (s *Server) WriteSiteCred(ctx context.Context, in *gnsrpc.SiteCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Printf("Received WriteSiteCred request for index: %d\n", in.Idx)
	if in.Idx > 32 {
		log.Println("site credential index out of range > 32")
		return &gnsrpc.GNSBadgeDataParam{}, errors.New("site credential index out of range > 32")
	}

	err := gnsserial.WriteSiteCred(in)
	if err != nil {
		log.Println(err.Error())
	}
	return &gnsrpc.GNSBadgeDataParam{}, err
}

func (s *Server) DeleteSiteCred(ctx context.Context, in *gnsrpc.SiteCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Println("Received DeleteSiteCred request")
	err := gnsserial.DeleteSiteCred(in)
	if err != nil {
		log.Println(err.Error())
	}
	return &gnsrpc.GNSBadgeDataParam{}, err
}

func (s *Server) ReadSiteCreds(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.Sites, error) {
	log.Println("Received ReadSiteCreds request")
	sites, err := gnsserial.ReadSiteCreds()
	if err != nil {
		log.Println(err.Error())
	}
	return sites, err
}

func (s *Server) ReadSiteCred(ctx context.Context, in *gnsrpc.SiteCred) (*gnsrpc.SiteCred, error) {
	log.Println("Recceived ReadSiteCred request")
	site, err := gnsserial.ReadSiteCred(in)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return site, err
}

func (s *Server) GetFreeWinCreds(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.FreeWinCreds, error) {
	log.Println("Received GetFreeWinCreds request")
	free, err := gnsserial.GetFreeWinCreds()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return free, nil
}

func (s *Server) ReadWinCreds(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.WinCreds, error) {
	log.Println("Received ReadWinCreds request")
	creds, err := gnsserial.ReadWinCreds()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return creds, nil
}

func (s *Server) ReadWinCred(ctx context.Context, in *gnsrpc.WinCred) (*gnsrpc.WinCred, error) {
	log.Println("Received ReadWinCred request")
	cred, err := gnsserial.ReadWinCred(in)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return cred, nil
}

func (s *Server) WriteWinCred(ctx context.Context, in *gnsrpc.WinCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Println("Received WriteWinCred request")
	err := gnsserial.WriteWinCred(in)
	return &gnsrpc.GNSBadgeDataParam{}, err
}

func (s *Server) DeleteWinCred(ctx context.Context, in *gnsrpc.WinCred) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Println("Received DeleteWinCred request")
	err := gnsserial.DeleteWinCred(in)
	return &gnsrpc.GNSBadgeDataParam{}, err
}

func (s *Server) Execute(ctx context.Context, in *gnsrpc.Text) (*gnsrpc.GNSBadgeDataParam, error) {
	log.Println("Received Execute request: ", in.Text)
	if strings.Contains(in.Text, "UNLOCKED ON") {
		gnsserial.SetUnlockMode(true)
		return &gnsrpc.GNSBadgeDataParam{}, nil
	} else if strings.Contains(in.Text, "UNLOCKED OFF") {
		gnsserial.SetUnlockMode(false)
		return &gnsrpc.GNSBadgeDataParam{}, nil
	} else {
		return &gnsrpc.GNSBadgeDataParam{}, errors.New("invalid command")
	}
}

func (s *Server) UnlockCard(ctx context.Context, in *gnsrpc.GNSBadgeDataParam) (*gnsrpc.Text, error) {
	var text gnsrpc.Text
	uuidstr, err := gnsserial.UnlockCard()
	text.Text = uuidstr

	return &text, err
}

func (s *Server) StreamCardStatus(param *gnsrpc.GNSBadgeDataParam, srv gnsrpc.GNSBadgeData_StreamCardStatusServer) error {
	ctx := srv.Context()
	var startTS time.Time
	startTS = time.Now()
	log.Println("We got a new client")
	fmt.Println("We got a new streaming client")
	for {
		select {
		case <-ctx.Done():
			log.Println("stream context done")
			return ctx.Err()
		default:
		}
		if time.Since(startTS) > 1000*time.Millisecond {
			startTS = time.Now()
			result := gnsrpc.CardStatus{}

			result.Type = gnsrpc.CardStatus_USB

			if !gnsserial.UnlockedMode {
				if gnsserial.IsAuthenticated {
					result.Status = gnsrpc.CardStatus_Authenticated
				} else if gnsserial.IsConnected {
					result.Status = gnsrpc.CardStatus_Connected
				} else {
					result.Status = gnsrpc.CardStatus_Disconnected
				}
			} else {
				if gnsserial.IsConnected {
					result.Status = gnsrpc.CardStatus_UnlockedModeReady
				} else {
					result.Status = gnsrpc.CardStatus_UnlockedMode
				}
			}

			if err := srv.Send(&result); err != nil {
				log.Printf("send error %v", err)
			}
		}
	}
}

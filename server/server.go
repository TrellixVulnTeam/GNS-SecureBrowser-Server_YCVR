package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	gnsapi "gnsdeviceserver/api"
	"gnsdeviceserver/proto/gnsrpc"
	gnsserial "gnsdeviceserver/serial"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"runtime/pprof"

	"github.com/kardianos/service"
	"github.com/mitchellh/panicwrap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type program struct{}

var web_host *grpc.Server
var log_location string
var panicFile *os.File
var bundle = []byte(`
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,917a262f5c469000

KrJPZozMSBO/7WaGQjTMY4D8Fq1YjYRjogyLkrNpTfmX3x5lOoRdFyVtmXlsNK5l
6Xkpp2yWynErb9OihH542nw8arBXjV34MP9IYFQ/W0QTAtOOa9P/7BUaQP+33Mnd
gJDRbowVnuG8e1euh5v7036iaxHL/5x3441GkMEOKm7msQjKKB5SD0QP/ctAjepD
UotD1zFuvdQAKrdPW4djWQuTKcgAJWG4cD8tQsLgxN8ySUf8zrkE4DP1o2t41EUu
iokZuUgiv+Iy7XmxuAdYu2yaZ/NGZJJo4ERkXlTqi3A6Ae3/4+43WplZur5lc+uw
62Akf2Eku/V8+My6exLCP0yHce2ArWoXJuNXq9nDsBzD7UvULjXVP3q1b0Hu68qG
4KttrsR9JXSLv4W21temV02D82kmE9v1QK39t7sbiYFTfL6zCaZ7s33Ll0xGFeGM
s0OqZME4iqh6K+Ap5rR3D/oi5S3CbyB0+Ati0+LhMfRNpk2AGbNQ3SaAom4Tr50K
1JIfhlPD/myUuU2uOitry9QiPoNQfwX+rQxf4MZ7MP4DFNHiPru5Wx1eA9iStplS
pKZS13Kl6DGyKkS0i1FbQHGvekRj5KNk9tq0WjUvi37yKFNBjRIpTEQLdAf5r8R3
2sw+kcKT23yfRVneyrO1+iu9+3YaTyXTzZUorgEUfAIE1j0THHZmNFXet6ooGL9w
ozLN2B1bWrwdVnWTDYqO3MenwsW/NJ5r2TEYBkP+jiJvYO4MDvoiTdMQsj7d1sjy
0S42WP94iil5V50fdqHloJgeNc/r1ZoMsMQJO2DpQJz34Xsp2FIOUWsY6SYwDrm1
QLdicv8eRfRu3Q/DHWNTBrkoT5QYuDZz4FTAmvg7I7lxULpjYy6Qb8OQF2s/E3nA
a1/DuH9iLpvMTkYthgiyx1lVIYoOrUab1Xp6V8HyAyiLgPB/jU2f5z1zEsnx2XiE
pA/XWSU4MT/8Oh79mXB2GCo/5mSQ6iFKZqOCzHv85BYU5A6jsexFu+sUymrGrdxz
4XC2Z4s42Ku+JV4t9+/HbFRwjLVY6ZZxCN3z8gF3O9okpFew9qydKN+uQL4b37V/
BEoZblVnUJD9OrxD/LTFJI7ZcC6P5Xs/nvNNRlhHgJoVk89FEjmCwCemJi56qmzS
AEWO16T2pEFBymoe4roqoGn2OklIZV8PtS5MMvmPubaAxRvd1Kcv7uXhqkgDSap7
eyBkqIFw8Tlvs1qo/tx2oQ1cwSkR0G0MfaErUtE6hq3aoyOvcmDARohhRGKzxcNp
wAytCqgCnYy/KJnEASN/rRP1A8urdpgDAlUmuXmscNkTd8SVm3i02LNnFx1OzbOy
nXx1QQCjtXQB/TO81Hqs5ttll804DRjlZ+Dyrhb8yUbnf9KCVhUT2eg92uuLKxDu
t1+IAmnOSoqwQt6SxbXzCtcCPBjjuI9wI9m69kd6dr8Lo5Os2cv/vCdNiCq9E0IE
5kONAQlf2BEq0LVZGJHpDluSw1zxuDFXx6IPdhjIsKk8h/zneIzmaovLpdIyH+6A
NmJrBH2zf1l5QGgaCqCqrAm9taeHDnMMoC5ltrMecFIScPnqwC2BeP/i/3f34mA3
-----END RSA PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIESjCCAjKgAwIBAgIRAIHuaOiLi4v+DW//1bTNAGQwDQYJKoZIhvcNAQELBQAw
IzEhMB8GA1UEAxMYZ2xvYmFsLW5ldC1zb2x1dGlvbnMuY29tMB4XDTIyMDMyMDA1
MjkyN1oXDTIzMDkyMDA1MzMzNVowGDEWMBQGA1UEAxMNR05TR1JQQ1NlcnZlcjCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL+bqSVF2J7lEcqQt3J7hS2r
Ohjn/SX0A5v1rYKoGykkp/D5Tz5AO/qzAH96oCPoZxy81oul5fyEiDrQkvLzYxBs
Gn6I/N3hAuHcUj3B94CTMl3TABMKWYJESXs9G5u72LLuMNzi0rR5DmS6hb/Z7V/z
kWiQse64hAEdq23/dz6EXU//RtMOWXFfuNCZytjwTnvuDc7XZCIvG1PPXi+dGGfJ
f5sR/oDD8B4EJoKjukatJVteTNafY9h8b8wMRGyOENdAXy7KZ9oKFnDYaeJFEHWn
rixGtk1v2V1t6CSgoQsBPoIERR2UYDpiEk7xoouOZlzJ8WkRLFiZ3jblThkl5sMC
AwEAAaOBgzCBgDAOBgNVHQ8BAf8EBAMCA7gwHQYDVR0lBBYwFAYIKwYBBQUHAwEG
CCsGAQUFBwMCMB0GA1UdDgQWBBQbSIURsYtJsMHZGwDVk2nMWOTHoDAfBgNVHSME
GDAWgBQQyG2yEAGccJZD+G1lkgxrfFcqSTAPBgNVHREECDAGhwR/AAABMA0GCSqG
SIb3DQEBCwUAA4ICAQBPxp+rTVXPoWl9cEWeOIZhp78bg4ynBeYQAjtc1DyVrjF2
/JvQA0zYP5aWg5U66+v9urlbfmeH1a8EUN9NwbDs4TzNodNhkSWaVUiqqMveuDuB
8EFwU26hEZbOwcjHYvbdFrrNnb9ZGrNP+2LVAbpUOL0IoxAHcuQK/VR2YZ59Phty
vUb05yLl5FyQPxh+3hXw5rVA9jnlRNGUwh7FvSzlR1yAbEuPS+31kIVDpazT5Xfr
nXuPKGanO5u7UpF4JfgEO9aEYBjW68KJfjkb9BTNEkO/NAEaezzZZFDfJfuKc+OQ
9hULLorgcobb8R3kxPDqKpKyMUsIAc8AQZJbDJqorUQPhNUC8JteyKKYiv86TzdR
V3L/BLr16N8+1yfFMyIO/VuACz5HHG1CqI11tX9hNsWVIbYCbj9h8uO9FYK5IDA5
qUlaShYRkWVv6QYetOkpm4V6BJql4LeGrww/zIcl/Xt/pd7QwcVGmHWeWGlfSclA
kmnXkxoFNR/Erf5plCl+c3kbmpMpJhuP1jXiZrmo5xOInCSNXX+qRpXtxytzxcxD
J22lQ+RsMtwAe89YTcH0IM2TLbTwH3gQ9dtxoNbccOLbbBPh7KYBst4eqSec9tEV
NDqrZIrYrcqNSSqeYG17spQZdg8cmcVMTckSWGYRpw/ax9WSgE7ajnSo9t2JpA==
-----END CERTIFICATE-----
`)
var clientBundle = []byte(`
-----BEGIN CERTIFICATE-----
MIIFBjCCAu6gAwIBAgIBATANBgkqhkiG9w0BAQsFADAjMSEwHwYDVQQDExhnbG9i
YWwtbmV0LXNvbHV0aW9ucy5jb20wHhcNMjIwMzIwMDUyMzQyWhcNMjMwOTIwMDUz
MzM2WjAjMSEwHwYDVQQDExhnbG9iYWwtbmV0LXNvbHV0aW9ucy5jb20wggIiMA0G
CSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQC2fc6u3/E58VBmEkbWRK2ZRvuWhdKO
SLCDbBxlsWvXiws06EplsBe8lTKd1yYmo2Ugv3cRFVTALoVa22xciEijZXnC6Qva
r/gdrjdFkrdzuM7PWRf8IDNCgK9V0u7SRkHZ0PUg2i8Ec7zqA/OkK/SsM6OsDZYc
OGsTwfHEKol7/doPj+nLdbJaWbCvNFODXmr8ziC/dL+1DwWkeMMFMck/hWYdnnNx
Ojj26nkQIzRZWT4ByRcgyGP4yBPB1Oy4+mCzhuMZRLR3TQy3yCO3jjF19D1lr2MN
Ky3+/eW+0CIVXjT1CEGkCQygLECyuhl1T8KvNswZuvsqi9mC1tJ5W6g9eMZGChq2
wv9rIxw1rY9fInlCbT5kKjRcWBNDPd4l1fZ4E55Fsv0uqnCQDyPbzm8HPFPII+4C
mgsE/ro8q3MImeZwb1F21JMJDpCSZlnxcBBp3Entc84xbEhbbVFqibF6qa15vuQX
BFYP1KiavTpS6MoCgJMYpHX4kJdgat7TeX3PsRkNsUAM9tAT6KyAq68v87LaTOPO
nejqEzHl/RBTVnl7Xq6Nf6DyPCY+0SP586pcY69hYXMagaXD/g4l1jis2BbeeozL
UH033iVK47x3ht/mz5eKDJIgxWZ9QwecW5EGS19HjjmKxdm92KilfEPtt8B+v7Na
cdzvCnHttGYeJwIDAQABo0UwQzAOBgNVHQ8BAf8EBAMCAQYwEgYDVR0TAQH/BAgw
BgEB/wIBADAdBgNVHQ4EFgQUEMhtshABnHCWQ/htZZIMa3xXKkkwDQYJKoZIhvcN
AQELBQADggIBAFHEntIGaHROaq95dxRUXghRakGKur8BLkPg87hqqGzVj0qQzEzo
DPZpfcjhdMiufX9w2uHzNMegdlbdUsNBOqBFYKGuDlnyUY38pKgs5mTOeUL1E0y8
jGDmmSnLXcZfvARbyfV95j0Ggh+cAHbs3Mbp0pP/0ZZBkDeaZ/Ks+IVi5PqV2hC/
04ak+o/cDJZTdQqvoFMxaibw3fxka0rcfC/5An9L8aqBE4Xi2GVu2gaN0ftAfnCA
gRrTa6aMz6gRVS4pYBRZiR6/uujR6CkTkSMGzI0wykF+kQO/WO/S7lW0+0iiZR6q
5sQ+ZM8LwxQCog1eLngEEullAnbeDCjkvAY6V8CnYq3TWJzdAyOzyYsoffGWeNT6
4i6Hh63Da+cz0+8QL5w/ju5/0fY4x6hlRD3xvWBS8RRU/xuzi65nbbwMkVcQsRvO
5yF6ITlWWhHxzVdc31+OMQeHrNfaKY9XykJbBhFfOYi+KqQjYjt5HFAkP5ZDU9kr
qhRHHTXBJ73+5Xm9FCcoI2q3+4iB8D2n6kgDKJBqxilERlJCorQNgxnjD+n2Glso
y+5kA3Oh86rYRZGcgMnPocYae+U2bBbHNFG62FYwwATxowRrAu3unqq7mw9xPrdY
ZWlew2p+LbA9W0nv/H2185SNJ1LmRG1goSLB39z6P0a9qT193Tom08ZT
-----END CERTIFICATE-----
`)

func GenerateCredential() (credentials.TransportCredentials, error) {
	keyBlock, certsPEM := pem.Decode(bundle)
	//fmt.Println(certsPEM)
	// Decrypt key
	keyDER, err := x509.DecryptPEMBlock(keyBlock, []byte("mikeserver"))
	if err != nil {
		return nil, err
	}
	//fmt.Println("Decrypted PEM successfully")
	keyBlock.Bytes = keyDER
	keyBlock.Headers = nil
	keyPEM := pem.EncodeToMemory(keyBlock)
	cert, err := tls.X509KeyPair(certsPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(clientBundle) {
		log.Fatal("failed to add client CA's certificate")
	}
	//ClientAuth: tls.RequireAndVerifyClientCert
	//ClientAuth: tls.NoClientCert
	MyTLS := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(MyTLS), nil

}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	go gnsserial.Init(false)
	log.Println("Setting up GNS GRPC server")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *gnsapi.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gns_cred, err := GenerateCredential()
	if err != nil {
		log.Fatal(err)
	}
	//web_host = grpc.NewServer(grpc.Creds(gns_cred))
	fmt.Println(gns_cred)
	//web_host = grpc.NewServer()
	web_host = grpc.NewServer(grpc.Creds(gns_cred))
	gnsrpc.RegisterGNSBadgeDataServer(web_host, &gnsapi.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := web_host.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Println("Stopping serial")
	//web_host.Stop()
	gnsserial.Stop()
	return nil
}

func CPUprofiling() {
	fd, err := os.Create(".cpu.prof")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer fd.Close()
	pprof.StartCPUProfile(fd)
	defer pprof.StopCPUProfile()
}

func main() {

	/*
		cpu, err := os.Create(".cpu.prof")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer cpu.Close()
		pprof.StartCPUProfile(cpu)
		defer pprof.StopCPUProfile()
	*/

	svcConfig := &service.Config{
		Name:        "GNSBadgeService",
		DisplayName: "GNSBadge Service",
		Description: "GNS Badge Data server",
	}

	//Get file path from where the exe is launched
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log_location = dir + "/logs"

	if _, err := os.Stat(log_location); os.IsNotExist(err) {
		os.MkdirAll(log_location, 0777)
	}

	f, err := os.OpenFile(log_location+"/gnsserver_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	panicFile, err = os.OpenFile(log_location+"/gnsserver_panic_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening pnaic file: %v", err)
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	exitStatus, err := panicwrap.BasicWrap(panicHandler)
	if err != nil {
		// Something went wrong setting up the panic wrapper. Unlikely,
		// but possible.
		panic(err)
	}

	// If exitStatus >= 0, then we're the parent process and the panicwrap
	// re-executed ourselves and completed. Just exit with the proper status.
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func panicHandler(output string) {
	// output contains the full output (including stack traces) of the
	// panic. Put it in a file or something.
	defer panicFile.Close()
	dt := time.Now()
	panicFile.WriteString("\n\n")
	panicFile.WriteString("Panic at: " + dt.String())
	panicFile.WriteString(output)
	//fmt.Printf("The child panicked:\n\n%s\n", output)
	os.Exit(1)
}

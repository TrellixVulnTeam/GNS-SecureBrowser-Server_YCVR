package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"gnsdeviceserver/proto/gnsrpc"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr = flag.String("addr", "127.0.0.1:50051", "the address to connect to")
)

var bundle = []byte(`
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,0438a089477a9f93

bl1XE43gNl0lj1USoT5UjSgmjh9vo2wOemnpflo5Xa84ApDJiHEISuSzHcMLUOmm
QV9ekpubtj4LC4jN4XsORkonupjO2t2cwO8DtWVtoE+pblJWu+HNK8tfwhY0A0j9
LnTtXX46LV2c1mlh0xZajKjtVeoOtNlbbLeB7txEyhd7on+dLlGucM0+aucGfO5s
f+M3U3Xi8VAZtGlJVVFwelkmcgZO4dTDqRB9AiJLyeXlFxoGyHdYBeiP9p0l7Nyk
npNdqJzJQaw08ubFvTuIcN9NqetuldJjUm9vt10tC7GizrXY69jaEaDkGFlq9hwk
PSWYqleg5J/fJT6TlzeVfA23szHtwOORGU7UDSqkQl1ZRboL7bKdCewpstNDhmvk
bEB8PpTh0dc3ZqTsm8w6WPitwNKHeRKj7HRpbjSTVNJv9C6+QexRVTl+QhcLhIkp
TdeefaJs25WKMc8Hbbc34cZ9fa59mEKeub7EpXUTwB7maUWrVAx+2COtsb6gfh3W
TCssqXJ0vzcEFsvdNaFR5RURHqQRlDcbF9zHAXeJJBAg5YN97qcWdx73hoccyiJY
t4xvQZr0D9bA+65fcfYxDfVuwKdFmF9m7q7B9nj4u9RTZkIeRedCjWLa+mn7Plwq
ta/j5XzXWOCoLEn3ZV2T+uIC0OrK5OxfPIDn70cJuLGeJw/NRs/fVC9xvYWfv7z7
KmuHOiBjThAoeP17xASGJ5CzLP7lrmfen0w6ONr1E+TBYTkdSFr1elXQ3wEC0JB4
R3hBq4n7gkIemGLgjEeCP4q9u0iC9yQwDP4nnOfbGG0zMY90MF8Ih6T7n7nd2dHS
6TBgNcN+e45sXnLa1jieImy2kPUq5ZdOXxcC97A7JH0DgG0ZHgulWY5srm/YQ/3/
lsI4y3H3WQT5KKJlRxdn/Edio+NU9S2NKNzgixlD+9m7qvjMtbgbBT2e8SJ0Gu/2
celrCb4EZbJgyC3WNFtLNVzybnXg9Uod7zV9AeF1x61mE2BYsti2B/yygrYQ1rEr
rxHqmqOhDEEYIEOuy85F0Z4RSMAVUw6dBBj0KffDSXKwaBhTocfDPTvulJ14KEv6
dPOebAUrLKN+DI5cNTZ+aP1fL0lJO3zcQzEeWw+8lnDs9KNEXFxB/K2XxWEmoGol
xeYUrvNmGRqSbcFtEFcKjQsm2E39PV28gqeEB2kIY0d0y7WXkmdKFDjhFUmHa9K/
7IU/kfcZOZptyvzWRj5S6vsd32YW4mONNZ1ATkINRnMLO9spsv/ZwUkvEp2uHnxc
DsPb/tArjU0OTwW35HRRE5B8UvCgdvjuERdmvSGEt2omkzRA+q3JMGoYBdJbB7Q+
7QEEY+NFXOvq75PALosMRd+BmkHF7v+EHYGPAXUkd5lD3NrdWDbSZasS8A3rIUUa
24nfsORqZb172ocIKYrzEzNOvYUD1ZP69JdsQb6KR4pJrXH8sIuZLZ60ZaK9Jv4o
KUQcAmOV+M0ejCopYTeaBVcyCAHuANVaRIpn7xAmMBAoErjzNDGQZs8/99mnaDaT
X9/G4zGYf/IFiX1sSLbe0ppdSMIoZsThVAgRxnI6dcbyCdmNFRyCM186BBJwNkNU
-----END RSA PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIELzCCAhegAwIBAgIQeD13MKKSMca6d6OdDhp0yDANBgkqhkiG9w0BAQsFADAj
MSEwHwYDVQQDExhnbG9iYWwtbmV0LXNvbHV0aW9ucy5jb20wHhcNMjIwMzIwMDUz
MTE4WhcNMjMwOTIwMDUzMzM1WjARMQ8wDQYDVQQDEwZDbGllbnQwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDIeKqYGiKvbppfOz8+vcbyTTU7f+pKQf1d
+cOrYXOUt1cjL/3sYQSHdMWzPdw/SI6lDpoWb/oYryhPKsy53X2h5Xmr58lpbAh8
Qe8w3WouNjk62JMgjAPPhQU/w7cbMLGc/KBJMAKaffZTXPKRHbOhhCOTcvcY/o13
pVcCPM6DWVOMcDKV2EpnptAgjrLPYmEau+aClf/QNFsEFt+MYddN0AdvFuXpaMDx
XIgaXH648jkgmWrkuP+3Wzu/4492T/XhRTWOdFuepnygDbHltfmRfisu8E0TWJxo
14vxprjHdb5s0eI1by5OfxKrPQeSJBrBrTg98zYtFxw/1JSPcMXtAgMBAAGjcTBv
MA4GA1UdDwEB/wQEAwIDuDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIw
HQYDVR0OBBYEFKtAWnSSJ8UnZwOu9B0//wyyKFc9MB8GA1UdIwQYMBaAFBDIbbIQ
AZxwlkP4bWWSDGt8VypJMA0GCSqGSIb3DQEBCwUAA4ICAQAFjgZwRuL5+z3gj+w/
5bdvErqkKQTgWR/uPMQKAs4CclQ/Axr6v1nDISLCtmM1FjxwsxupDUuPxsUXpWqQ
t/WeKnzyqWNndrhJ8Hh/OqBFdelACD5XLpWu7tKV9Tt3ef+unyMxRvoDsWtR+IOY
88JEXrJVeTcNQJuw8UNbvCi2R72VehvsR7smtUzexfRHZA5z7ylqvVvrKu1nptTp
xl7G+Mnc7zYxz54u1HKg9BvwC9M8ilSObDtScN1zXDjR8gSDZkXA3u2xyUTSysDH
GPjuR0SWIPXpEWRFsNKg7/Famny2tGU19LLoeUvzjq5O3pBNYlIFq5jknhmaxlU1
XfOb8yUY+ZzFLvCVxcla6rxhRT1dZptUEM9AcVGZW7Zv2ivRPCqh2qNNBgYUWORX
sOAN9b3tRONrK5+EyT/uLFtoRc7B9rSGxH7VWEGrc3Lon3jXvhKgouojJJpsTxyz
V6TViUSk/UXJbugTctYXDoYL8A1xH3A3edChHdzYRob3RTeHfgdQjqT4qptIvEYP
HrfraofBiEpUekPVe48U71+5i0IwOxq3/A2/hq6hqJJI+8uQAYvhRW4efV/fIndb
S7OdsXCtcRptssDbRy2O3UMqefSR0Q9qyIOOocBCs6wre6NOe8KRHGKjuCXtlcxQ
pzseGmSo8Lq2Lo+D3neBCIwoIg==
-----END CERTIFICATE-----
`)

var CABundle = []byte(`
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
	keyDER, err := x509.DecryptPEMBlock(keyBlock, []byte("mikeclient"))
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

	_, caPEM := pem.Decode(CABundle)
	fmt.Println(caPEM)
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(CABundle) {
		log.Fatal("failed to add client CA's certificate")
	}

	MyTLS := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(MyTLS), nil

}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	tlsCredentials, err := GenerateCredential()
	if err != nil {
		log.Fatal(err.Error())
	}
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gnsrpc.NewGNSBadgeDataClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ruuid, err := c.ReadUUID(ctx, &gnsrpc.GNSBadgeDataParam{})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	if ruuid != nil {
		fmt.Println("HW UUID: " + ruuid.Uuid)
	}

	/*
		free_sites, err := c.GetFreeSites(ctx, &gnsrpc.GNSBadgeDataParam{})
		if err != nil {
			log.Printf("could not greet: %v", err)
		}
		fmt.Println("Free sites: ")
		if free_sites != nil {
			for i := 0; i < len(free_sites.Idx); i++ {
				fmt.Println(free_sites.Idx[i])
			}
		}

		var site1 gnsrpc.SiteCred
		site1.Code = "CH1"
		site1.Idx = 1
		site1.Username = "peter1"
		site1.Password = "peter1password"

		var site2 gnsrpc.SiteCred
		site2.Code = "CH2"
		site2.Idx = 21
		site2.Username = "peter2"
		site2.Password = "peter2password"

		fmt.Println("Writing site1 with idx: ", site1.Idx, " and site2 with idx: ", site2.Idx)
		_, err = c.WriteSiteCred(ctx, &site1)
		_, err = c.WriteSiteCred(ctx, &site2)

		fmt.Println("Reading all sites")

		sites, err := c.ReadSiteCreds(ctx, &gnsrpc.GNSBadgeDataParam{})
		if err != nil {
			log.Printf("could not greet: %v", err)
		} else {
			if sites != nil {
				for i := 0; i < len(sites.Sites); i++ {
					fmt.Printf("Idx: %d, Code: %s, Username: %s, Password: %s\n", sites.Sites[i].Idx, sites.Sites[i].Code,
						sites.Sites[i].Username, sites.Sites[i].Username)
				}
			}

		}*/

	fmt.Println("Calling store UUID")

	_, err = c.StoreUUID(ctx, &gnsrpc.GNSBadgeDataParam{})

	fmt.Println("Calling readzone2 UUID")
	uuid, err := c.ReadUUIDZone2(ctx, &gnsrpc.GNSBadgeDataParam{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("zone2 uuid: " + uuid.Uuid)

	/*
			var read_site1 gnsrpc.SiteCred
			read_site1.Idx = 21

			read2, err := c.ReadSiteCred(ctx, &read_site1)
			fmt.Println("individiual read site2: code,idx,usename,password : ", read2.Code, read2.Idx, read2.Username, read2.Password)

			fmt.Println("Deleting site 2 with idx: ", site2.Idx)

			_, err = c.DeleteSiteCred(ctx, &site2)

			fmt.Println("Reading sites again")
			sites, err = c.ReadSiteCreds(ctx, &gnsrpc.GNSBadgeDataParam{})
			if err != nil {
				log.Printf("could not greet: %v", err)
			} else {
				if sites != nil {
					for i := 0; i < len(sites.Sites); i++ {
						fmt.Printf("Idx: %d, Code: %s, Username: %s, Password: %s\n", sites.Sites[i].Idx, sites.Sites[i].Code,
							sites.Sites[i].Username, sites.Sites[i].Username)
					}
				}

			}

			fmt.Println("Get free wincreds")
			wincredsIdx, err := c.GetFreeWinCreds(ctx, &gnsrpc.GNSBadgeDataParam{})
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(wincredsIdx)

			var windows1 gnsrpc.WinCred
			windows1.Domain = "MY1"
			windows1.Idx = 2
			windows1.Username = "peter"
			windows1.Password = "pham"

			fmt.Println("Writing 1 wincred")
			fmt.Println(windows1)
			_, err = c.WriteWinCred(ctx, &windows1)

			fmt.Println("Get free wincreds")
			wincredsIdx, err = c.GetFreeWinCreds(ctx, &gnsrpc.GNSBadgeDataParam{})
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(wincredsIdx)

			fmt.Println("Reading wincreds")

			creds, err := c.ReadWinCreds(ctx, &gnsrpc.GNSBadgeDataParam{})
			if err != nil {
				fmt.Println(err.Error())
			}
			if creds != nil {
				for i := 0; i < len(creds.Wincreds); i++ {
					fmt.Printf("Idx: %d, Domain: %s, Username: %s, Password: %s\n", creds.Wincreds[i].Idx, creds.Wincreds[i].Domain,
						creds.Wincreds[i].Username, creds.Wincreds[i].Username)
				}
			}
			var windows2 gnsrpc.WinCred
			windows2.Idx = 2
			fmt.Println(windows2)
			fmt.Println("Reading 1 wincred at idx ", windows2.Idx)
			cred, err := c.ReadWinCred(ctx, &windows2)
			if err != nil {
				fmt.Println(err.Error())
			}
			if cred != nil {

				fmt.Printf("Idx: %d, Domain: %s, Username: %s, Password: %s\n", cred.Idx, cred.Domain,
					cred.Username, cred.Username)

			}

			fmt.Println("Deleting wincred at idx ", windows2.Idx)
			_, err = c.DeleteWinCred(ctx, &windows2)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Get free wincreds")
			wincredsIdx, err = c.GetFreeWinCreds(ctx, &gnsrpc.GNSBadgeDataParam{})
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(wincredsIdx)


		var cmd gnsrpc.Text
		cmd.Text = "UNLOCKED OFF"
		_, err = c.Execute(ctx, &cmd)
	*/

	stream, err := c.StreamCardStatus(context.Background(), &gnsrpc.GNSBadgeDataParam{})
	if err != nil {
		log.Printf("openn stream error %v", err)
	}
	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
			}
			if err != nil {
				log.Printf("can not receive %v", err)
				break
			}

			fmt.Println(resp.Status)
		}
	}()

	<-done

}

package gnsserial

import (
	"encoding/hex"
	"errors"
	"fmt"
	gnscrypto "gnsdeviceserver/crypto"
	"gnsdeviceserver/proto/gnsrpc"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

var myport serial.Port
var IsAuthenticated bool
var IsConnected bool
var UnlockedMode bool
var portError bool
var portName string
var debug bool
var shared_key []byte

var portInUsed sync.Mutex

var pemLoc = ""

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func FindPort() (string, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return "", err
	}
	if len(ports) == 0 {
		return "", nil
	}
	system_os := runtime.GOOS
	//fmt.Println(len(ports))
	//fmt.Println(system_os)

	for _, port := range ports {
		if port.IsUSB {
			//fmt.Println(port.Name)
			if strings.Contains(port.Name, "usbmodem") || strings.Contains(port.Name, "COM") ||
				strings.Contains(port.Name, "tty") {
				//fmt.Println("PID: " + port.PID)
				//fmt.Println("Product: " + port.Product)
				//fmt.Println("VID: " + port.VID)
				//fmt.Println("Serial: " + port.SerialNumber)

				switch system_os {
				case "darwin":
					var lastChar = port.Name[len(port.Name)-1:]
					if lastChar == "3" {
						return port.Name, nil
					}
				case "linux":
					var lastChar = port.Name[len(port.Name)-1:]
					if lastChar == "1" {
						return port.Name, nil
					}
				default:
					//fmt.Println(port.SerialNumber)
					if strings.Contains(port.SerialNumber, "MI_02") {
						return port.Name, nil
					}
				}

			}
		}
	}

	// we did not find a port
	return "", nil

}

var StopInitiated bool

func monitorPort() {
	//defer wg.Done()
	var err error
	var success bool
	for {
		if StopInitiated {
			if IsConnected {
				myport.Close()
			}
			log.Println("Stopping monitoring port")
			return
		}
		if (!IsAuthenticated && !UnlockedMode) || (!IsConnected && UnlockedMode) {
			//Try to connect and pingpong If we are not authenticated and not in unlocked moe
			//or we are not connected and in enabledunlockedmoe
			time.Sleep(1000 * time.Millisecond)
			portName, err = FindPort()
			if err != nil {
				log.Println(err.Error())
				IsConnected = false
				continue
			}
			if len(portName) > 1 {
				if !IsConnected {
					mode := &serial.Mode{
						BaudRate: 115200,
						Parity:   serial.NoParity,
						DataBits: 8,
						StopBits: serial.OneStopBit,
					}
					//fmt.Println("Connecting to port: " + portName)
					myport, err = serial.Open(portName, mode)

					if err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						IsConnected = true
						myport.SetReadTimeout(120 * time.Millisecond)
					}
				}
				if UnlockedMode == false {
					success, err = pingPong()
					if success {
						IsAuthenticated = true
						if debug {
							fmt.Println("Connected to: " + portName)
						}
					} else {
						if debug {
							fmt.Println("Failed pingpong test")
						}
						IsConnected = false
						IsAuthenticated = false
						myport.Close()
					}
				} else {
					fmt.Println("We are in unlocked mode no need to try pingpong")
					IsConnected = true
				}

			} else {
				//we did not find port with valid name
				IsAuthenticated = false
				IsConnected = false
				if debug {
					fmt.Println("No port found yet")
					fmt.Println(IsConnected)
				}
			}
		} else {
			//if already connected  let's try to see if connection is alive
			time.Sleep(1000 * time.Millisecond)
			_, error := myport.GetModemStatusBits()
			if error != nil {
				if debug {
					fmt.Println("Lost connection or error on Port. Will try to re-connect")
				}
				IsAuthenticated = false
				IsConnected = false
			}
		}
	}
}

func Stop() {
	StopInitiated = true
}

func SetUnlockMode(in bool) {
	portInUsed.Lock()
	defer portInUsed.Unlock()
	UnlockedMode = in
	fmt.Println("setting unlockmode to: ", UnlockedMode)
}

func Init(isDebug bool) {
	StopInitiated = false
	/*
		system_os := runtime.GOOS
		switch system_os {
		case "darwin":
			pemLoc = "/Library/GNS/GNSPrivate.pem"
			log.Println("Using pem from: " + pemLoc)
		default:
			pemLoc = "GNSPrivate.pem"
		}*/
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	pemLoc = dir + "/GNSPrivate.pem"
	log.Println("Using pem from: " + pemLoc)
	debug = isDebug
	IsConnected = false
	IsAuthenticated = false
	portError = false
	UnlockedMode = false
	portName = ""
	if debug {
		fmt.Println("Starting to monitor serial ports")
	}
	monitorPort()
}

func readResponseDecode(port serial.Port, count int, command string, delay int, needFlush bool) (string, error) {

	encrypted := gnscrypto.EncodeData(command, shared_key, pemLoc)
	response, err := readResponse(port, count, encrypted, delay, needFlush)
	if err != nil || len(response) < 1 {
		return "", err
	}
	//log.Println(response)
	result := gnscrypto.DecodeData(response, shared_key, pemLoc)
	return result, nil
}

func writeResponseDecode(port serial.Port, count int, command string, delay int) error {

	encrypted := gnscrypto.EncodeData(command, shared_key, pemLoc)
	err := writeResponse(port, count, encrypted, delay)
	if err != nil {
		return err
	}
	return nil
}

func readResponse(port serial.Port, count int, command string, delay int, needFlush bool) ([]byte, error) {

	var n int
	var err error
	//var data []byte
	sent := 0
	for sent < len([]rune(command)) {
		increment := min(MaxWriteBytes, len([]rune(command))-sent)
		//fmt.Printf("Sending %d to %d\n", sent, sent+increment)
		n, err = port.Write([]byte(command[sent : sent+increment]))
		if err != nil {
			return nil, err
		}
		sent += increment
		//fmt.Printf("Sent %v bytes\n", n)
		//fmt.Println("Remaining: " + command)
	}
	_, err = port.Write([]byte("\n"))
	if err != nil {
		return nil, err
	}

	if needFlush {
		time.Sleep(3 * time.Millisecond)
		_, err = flushResponse(port)
		if err != nil {
			return nil, err
		}
		//fmt.Println("Finished flushing response" + string(data))
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
	// Read and print the response
	buff := make([]byte, count)
	var result []byte
	for {
		// Reads up to 100 bytes
		n, err = port.Read(buff)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			//fmt.Println("No more data in response")
			//result = nil
			break
		}
		//fmt.Printf("%s", string(buff[:n]))
		//fmt.Println()
		//fmt.Println(buff[:n])

		// If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			//fmt.Println("Found new line in response")
			result = append(result, buff[:n]...)
			break
		}
		result = append(result, buff[:n]...)

	}
	return result, nil
}

func flushResponse(port serial.Port) ([]byte, error) {

	var result []byte
	buff := make([]byte, 400)
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			result = nil
			break
		}
		//fmt.Printf("%s", string(buff[:n]))
		// If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			result = buff[:n]
			break
		}

	}
	return result, nil
}

func writeResponse(port serial.Port, count int, command string, delay int) error {

	var err error
	sent := 0
	for sent < len([]rune(command)) {
		increment := min(MaxWriteBytes, len([]rune(command))-sent)
		_, err = port.Write([]byte(command[sent : sent+increment]))
		if err != nil {
			return err
		}
		sent += increment
		time.Sleep(time.Duration(delay) * time.Millisecond)
		flushResponse(myport)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	_, err = port.Write([]byte("\n"))
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	flushResponse(myport)
	return nil
}

func pingPong() (bool, error) {
	var err error
	var result []byte
	t_tot := time.Now()
	t_now := time.Now()
	//result,err = writeResponse(myport, "LOAD_PUBLIC "+public_key, 100, false)

	/*fmt.Println("Sending GET_ID command:")
	result, err = readResponse(myport, 100, "GET_ID", 150, false) //true to flush out ST-Safe Driver Initialized
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println("UUID response: " + string(result))
	fmt.Println(result)*/

	//fmt.Println("Sending Auth command:")
	result, err = readResponse(myport, 390, "AUTH", 100, false) //true to flush out ST-Safe Driver Initialized
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	//fmt.Println("Auth string cert:")
	//fmt.Println(string(result))
	//fmt.Println("AUTH cert length:")
	//fmt.Println(len(result))
	if len(result) < 300 {
		return false, errors.New("AUTH response has invalid cert length")
	}
	cert := result
	//fmt.Println(hex.EncodeToString(cert))

	// Create raw challenge
	raw_challenge, digest := gnscrypto.GenerateChallenge()
	//fmt.Println("Raw challenge: ")
	//fmt.Println(raw_challenge)
	//fmt.Println("Digest: ")
	//fmt.Println(digest)
	//fmt.Println(len(digest))
	//fmt.Println("Sending challenge")
	result, err = readResponse(myport, 200, "CHALLENGE "+digest, 100, false)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	//fmt.Println("Challenge result: ")
	//fmt.Println(string(result))
	//fmt.Println(len(result))
	if len(result) < 64 {
		//fmt.Println("Got invalid challenge length")
		return false, errors.New("CHALLENGE response has invalid length")
	}
	s_challenge := result[0:64]
	//fmt.Println(s_challenge)
	//fmt.Println(len(s_challenge))

	verified := gnscrypto.VerifySignature(s_challenge, raw_challenge, cert)
	//fmt.Println(verified)
	if verified > 0 {
		//fmt.Println("We got verified signature")
	} else {
		//fmt.Println("Signature verify failed")
		return false, errors.New("signature verify failed")
	}

	//fmt.Println("Sending REM_AUTH")
	result, err = readResponse(myport, 200, "REM_AUTH", 100, false)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	//fmt.Println("REM_AUTH RESP: " + string(result))

	if len(result) < 16 {
		return false, errors.New("invalid rem_auth response")
	}

	card_s_byte := result[0:16]
	//fmt.Println(len(card_s_byte))

	// Sign challenge
	//fmt.Println("Sending REM_SIGN")
	signed_challenge := gnscrypto.SignChallenge(card_s_byte, pemLoc)
	//fmt.Println((signed_challenge))
	result, err = readResponse(myport, 200, "REM_SIGN "+signed_challenge, 100, false)
	if err != nil {
		//fmt.Println(err.Error())
		return false, err
	}
	//fmt.Println("Response: " + string(result))

	fmt.Println("elapsed after REM_SIGN: ")
	fmt.Println(time.Since(t_now))

	//time.Sleep(200 * time.Millisecond)

	fmt.Println("Sending KEY_ESTABLISH")
	sharedKey, err := readResponse(myport, 200, "KEY_ESTABLISH", 100, false)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println("KEY_ESTABLISH response: " + string(sharedKey))
	fmt.Println(len(sharedKey))
	if len(sharedKey) < 65 {
		fmt.Println("Invalid shared key length")
		return false, nil
	}
	if len(sharedKey) > 65 {
		shared_key = sharedKey[0:65]
	} else {
		shared_key = sharedKey[0:65]
		flushResponse(myport)
	}

	fmt.Println("Trimmed Share key: " + string(shared_key))
	fmt.Println(len((shared_key)))
	fmt.Println("elapsed after KEY_ESTABLISH: ")

	time.Sleep(100 * time.Millisecond)
	flushResponse(myport)
	fmt.Println(time.Since(t_now))
	fmt.Println("Sending PING")
	t_now = time.Now()
	encrypted := gnscrypto.EncodeData("PING", shared_key, pemLoc)
	fmt.Println("elapsed after Encode: ")
	fmt.Println(time.Since(t_now))

	fmt.Println("Encoded PING: " + encrypted)
	t_now = time.Now()
	result, err = readResponse(myport, 80, encrypted, 100, false)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println("elapsed after sending PING: ")
	fmt.Println(time.Since(t_now))
	fmt.Println("receive hex: " + hex.EncodeToString(result))
	fmt.Println(string(result))
	fmt.Println(len(result))
	if len(result) <= 1 {
		fmt.Println("Invalid response length")
		return false, nil
	}
	//convert result byte to hex string
	// fmt.Println("Hex result: " + hex_result)
	// fmt.Println(len(hex_result))
	t_now = time.Now()
	decrypted := gnscrypto.DecodeData(result, shared_key, pemLoc)
	fmt.Println("elapsed after Decode: ")
	fmt.Println(time.Since(t_now))
	fmt.Println("Response: " + decrypted)
	if decrypted != "PONG" {
		return false, nil
	}

	fmt.Println("elapsed: ")
	fmt.Println(time.Since(t_now))

	fmt.Print("Total pingpong duration: ")
	fmt.Println(time.Since(t_tot))
	return true, nil
}

func ReadUUID() (string, error) {

	if !IsAuthenticated {
		return "", errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()
	cmd := "QUERY_ID"
	encrypted_cmd := gnscrypto.EncodeData(cmd, shared_key, pemLoc)
	result, err := readResponse(myport, 100, encrypted_cmd, 100, false)
	//result, err := readResponse(myport, 100, cmd, 100, false)
	if err != nil {
		log.Println(err.Error())
		return "", err
	} else {
		log.Println("QUERY_ID response: " + string(result))
		log.Printf("QUERY_ID response length: %d\n", len(result))
		if len(result) >= 9 {
			result = result[0:9]
			uuidStr := hex.EncodeToString(result)
			return uuidStr, nil
		} else {
			return "", errors.New("length of uuid is invalid")
		}
	}

}

func ReadUUIDZone2() (string, error) {
	var zone, offset, length int
	var cmd, result string
	var err error
	if !IsAuthenticated {
		return "", errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()
	zone = 2 //change back to zone 2 later
	offset = 0
	length = 64
	cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
	result, err = readResponseDecode(myport, 256, cmd, 80, false)
	if err != nil {
		return "", errors.New("error reading UUID on badge")
	}
	if len(result) < 1 {
		return "", errors.New("bad UUID")
	}
	result = strings.Trim(result, "\x00")
	return result, nil
}

func WriteUUID(uuid string) error {
	var zone, offset int
	var cmd string
	if !IsAuthenticated {
		errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()
	filler := strings.Repeat(SiteNull, 64-len([]rune(uuid)))
	uuid = uuid + filler
	offset = 0
	zone = 2
	fmt.Println("Writing data to zone 2: " + uuid)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(uuid), uuid)
	err := writeResponseDecode(myport, 200, cmd, 80)
	return err
}

// Read HW UUID then store to Zone3
func StoreUUID() error {
	fmt.Println("Caling store UUID")
	if !IsAuthenticated {
		errors.New("badge is not ready")
	}
	uuid, err := ReadUUID()
	if err != nil {
		return err
	}
	fmt.Println("Found hw UUID: " + uuid)
	return WriteUUID(uuid)
}

var IndexerBlockLen int = 5
var MaxZones int = 32
var IndexerZone int = 5
var IndexerZoneMaxLen int = IndexerZone * MaxZones
var SiteBlockLen int = 50
var SiteZoneStart int = 7
var SiteZone1Len int = 800
var SiteZone2Len int = 900
var SiteNull string = "\x00"
var WinCredZone int = 6
var WinCredBlockLen int = 80
var WinCredZoneMaxLen int = 700
var WinCredNull string = "\x00"
var MaxReadBytes int = 500
var MaxWriteBytes int = 100

func FormatCard(mode int, uuid string) error {
	//mode 0: everything
	//mode 1: Sites;
	//mode 2: WinCred;
	//Mode 3: Cert; //TODO
	var zone, offset, length int
	var data, cmd string
	var err error

	if !IsAuthenticated {
		return errors.New("badge is not ready")
	}

	portInUsed.Lock()
	defer portInUsed.Unlock()

	myport.SetReadTimeout(80 * time.Millisecond)
	defer myport.SetReadTimeout(80 * time.Millisecond)

	/*
		if len([]rune(uuid)) > 1 {
			data = uuid //"2eff852cd75f54ccfdb9e1214e042f60"
		} else {
			data = "2eff852cd75f54ccfdb9e1214e042f60"
		}

		filler := strings.Repeat(SiteNull, 64-len([]rune(uuid)))
		data = data + filler
		offset = 0
		zone = 2
		cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
		err = writeResponseDecode(myport, 100, cmd, 80)
		fmt.Println("Write response: " + string(temp))
		if err != nil {
			return err
		}*/

	if mode == 0 || mode == 1 {
		zone = IndexerZone
		offset = 0

		for i := 0; i < MaxZones; i++ {
			offset = i * IndexerBlockLen
			data = strings.Repeat("0", IndexerBlockLen)
			cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
			err = writeResponseDecode(myport, 100, cmd, 80)
			if err != nil {
				fmt.Println("Got write error on indexer format")
				return err
			}
			fmt.Printf("Offset :%d\n", offset)
		}

		// CLEAR SITE DATA
		zone = SiteZoneStart
		offset = 0
		write := 0
		for write < SiteZone1Len {
			if write+MaxWriteBytes < SiteZone1Len {
				length = MaxWriteBytes
			} else {
				length = SiteZone1Len - write
			}
			data = strings.Repeat(SiteNull, length)
			cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, length, data)
			err = writeResponseDecode(myport, 100, cmd, 80)
			if err != nil {
				return err
			}
			write += length
			offset += length
		}

		zone = SiteZoneStart + 1
		offset = 0
		write = 0
		for write < SiteZone2Len {
			if write+MaxWriteBytes < SiteZone2Len {
				length = MaxWriteBytes
			} else {
				length = SiteZone2Len - write
			}
			data = strings.Repeat(SiteNull, length)
			cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, length, data)
			err = writeResponseDecode(myport, 100, cmd, 80)
			if err != nil {
				return err
			}
			write += length
			offset += length
		}

	}

	if mode == 0 || mode == 2 {
		// CLEAR WINCRED DATA
		zone = WinCredZone
		offset = 0
		write := 0
		for write < WinCredZoneMaxLen {
			if write+MaxWriteBytes < WinCredZoneMaxLen {
				length = MaxWriteBytes
			} else {
				length = WinCredZoneMaxLen - write
			}
			data = strings.Repeat(WinCredNull, length)
			cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, length, data)
			err = writeResponseDecode(myport, 100, cmd, 80)
			if err != nil {
				return err
			}
			write += length
			offset += length
		}

	}
	log.Println("Done formatting")
	return nil
}

func ReadIndexer() (*gnsrpc.Sites, error) {

	var zone, offset, length, i, j int
	var cmd, result string
	var err error
	sites := []*gnsrpc.SiteCred{}

	zone = IndexerZone //change back to zone 2 later
	offset = 0
	length = IndexerZoneMaxLen
	cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
	result, err = readResponseDecode(myport, 256, cmd, 80, false)

	if err != nil {
		return nil, err
	}

	if len(result) < IndexerZoneMaxLen {
		return &gnsrpc.Sites{Sites: sites}, errors.New("invalid read on indexer")
	}
	for i = 0; i < length; i += IndexerBlockLen {
		site := gnsrpc.SiteCred{}
		j = i
		sitecode := result[j : j+3]
		if sitecode == "000" || sitecode == "\x00\x00\x00" {
			continue
		} else {
			site.Code = sitecode
		}
		j += 3

		offsetStr := result[j : j+2]
		intVar, err := strconv.ParseInt(offsetStr, 16, 32)
		if err != nil {
			continue
		}
		site.Idx = uint32(intVar)
		//site.Idx = uint32(k)

		sites = append(sites, &site)
	}
	return &gnsrpc.Sites{Sites: sites}, nil
}

func findAndRemove(input []uint32, val uint32) []uint32 {
	for i := 0; i < len(input); i++ {
		if input[i] == val {
			return append(input[:i], input[i+1:]...)
		}
	}
	return input
}

func GetFreeSites() (*gnsrpc.FreeSites, error) {

	if !IsAuthenticated {
		return nil, errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()

	free_sites := []uint32{}
	used_site, err := ReadIndexer()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for i := 0; i < 32; i++ {
		free_sites = append(free_sites, uint32(i))
	}

	for _, site := range used_site.Sites {
		fmt.Printf("site code: %s, siteIdx: %d\n", site.Code, site.Idx)
		free_sites = findAndRemove(free_sites, site.Idx)
	}
	return &gnsrpc.FreeSites{Idx: free_sites}, nil
}

func getZoneOffset(offset uint32) (int, int) {
	//convert absolute offset to zone relative offset

	if int(offset)*SiteBlockLen < SiteZone1Len {
		return SiteZoneStart, int(offset) * SiteBlockLen
	} else if int(offset)*SiteBlockLen >= SiteZone1Len {
		return (SiteZoneStart + 1), int(offset)*SiteBlockLen - SiteZone1Len
	} else if int(offset)*SiteBlockLen > (SiteZone1Len + SiteZone2Len) {
		return SiteZoneStart, -1
	}
	//we got invalid offset here
	return SiteZoneStart, -1
}

func serializeSiteCred(site *gnsrpc.SiteCred) string {
	if len(site.Username)+len(site.Password) > SiteBlockLen-1 {
		return ""
	} else {
		return site.Username + "\x0a" + site.Password
	}
}

func WriteSiteCred(site *gnsrpc.SiteCred) error {
	var zone, offset int
	var cmd, data string
	var err error
	if !IsAuthenticated {
		return errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()

	zone = IndexerZone
	offset = IndexerBlockLen * int((site.Idx)&0xff)

	//make sure site code isn't more than 3 character
	var fixed_site_code string
	if len(site.Code) > 3 {
		fixed_site_code = site.Code
		fixed_site_code = fixed_site_code[0:3]
	} else {
		fixed_site_code = site.Code
	}
	//need to convert to ascii hex representation for site index
	data = fixed_site_code + fmt.Sprintf("%02x", site.Idx) //fmt.Sprintf("%02d", site.Idx)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
	err = writeResponseDecode(myport, 100, cmd, 80)
	if err != nil {
		return err
	}
	data = serializeSiteCred(site)
	if len(data) > SiteBlockLen {
		return errors.New("site data length is longer than max siteblocklen")
	}
	if len(data) < SiteBlockLen {
		filler := strings.Repeat(SiteNull, SiteBlockLen-len(data))
		data = data + filler
	}
	zone, offset = getZoneOffset(site.Idx)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
	err = writeResponseDecode(myport, 100, cmd, 80)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSiteCred(site *gnsrpc.SiteCred) error {
	var zone, offset int
	var cmd, data string
	var err error

	if !IsAuthenticated {
		return errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()
	//**Clear indexer
	zone = IndexerZone
	offset = IndexerBlockLen * int(site.Idx)
	//bit-wise AND site.idx with 0xff to make sure we dont have any index that takes over 1 byte
	idx_str := fmt.Sprintf("%02d", site.Idx&0xff)
	data = "000" + idx_str
	// log.Println("Writing site code to indexer: " + data)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
	err = writeResponseDecode(myport, 100, cmd, 80)
	if err != nil {
		return err
	}

	//**Clear username/password
	zone, offset = getZoneOffset(site.Idx)
	data = strings.Repeat(SiteNull, SiteBlockLen)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
	err = writeResponseDecode(myport, 100, cmd, 80)
	if err != nil {
		return err
	}

	return nil
}

func setSiteCred(input string) gnsrpc.SiteCred {
	result := gnsrpc.SiteCred{}
	s := strings.Split(input, "\x0a")
	if len(s) >= 2 {
		result.Username = strings.Trim(s[0], SiteNull)
		result.Password = strings.Trim(s[1], SiteNull)
	} else if len(s) == 1 {
		result.Username = strings.Trim(s[0], SiteNull)
		result.Password = ""
	} else {
		result.Username = ""
		result.Password = ""
	}
	return result
}

func populateSite(sites *gnsrpc.Sites) error {
	var zone, offset, length int
	var cmd, result string
	var err error
	length = SiteBlockLen
	for i, site := range sites.Sites {
		//fmt.Printf("Populating site idx: %d\n", site.Idx)
		zone, offset = getZoneOffset(site.Idx)
		cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
		//fmt.Println(cmd)
		result, err = readResponseDecode(myport, 128, cmd, 80, false)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if len(result) == 0 {
			return errors.New("invalid site read")
		}
		// log.Println("Read site data: ")
		// log.Println(result)
		temp := setSiteCred(result)
		sites.Sites[i].Username = temp.Username
		sites.Sites[i].Password = temp.Password
	}
	return nil
}

func ReadSiteCreds() (*gnsrpc.Sites, error) {

	if !IsAuthenticated {
		return nil, errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()
	sites, err := ReadIndexer()
	if err != nil {
		return nil, err
	}
	err = populateSite(sites)
	if err != nil {
		return nil, err
	}
	return sites, nil
}

func ReadSiteCred(site *gnsrpc.SiteCred) (*gnsrpc.SiteCred, error) {
	//read a single site specified by site index
	var zone, offset, length int
	var cmd, result, sitecode string
	length = SiteBlockLen
	var err error
	var output gnsrpc.SiteCred

	if !IsAuthenticated {
		return nil, errors.New("badge is not ready")
	}
	portInUsed.Lock()
	defer portInUsed.Unlock()

	//Get site code
	zone = IndexerZone //change back to zone 2 later
	offset = 0
	length = IndexerZoneMaxLen
	cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
	result, err = readResponseDecode(myport, 256, cmd, 80, false)

	if site.Idx < 0 || site.Idx > 31 {
		return nil, errors.New("site index out of range. It has to be between 0 and 31")
	}
	//each site is spaced
	j := int(site.Idx) * IndexerBlockLen
	sitecode = result[j : j+3]

	//Get site data
	zone, offset = getZoneOffset(site.Idx)
	cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
	fmt.Println(cmd)
	result, err = readResponseDecode(myport, 128, cmd, 80, false)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("invalid site read")
	}
	// log.Println("Read site data: ")
	// log.Println(result)
	temp := setSiteCred(result)
	output.Username = temp.Username
	output.Password = temp.Password
	output.Idx = site.Idx
	output.Code = sitecode

	return &output, nil

}

func GetFreeWinCreds() (*gnsrpc.FreeWinCreds, error) {
	var result []uint32
	zone := WinCredZone
	offset := 0
	var length, i int
	read := 0
	var cmd string
	var readData string
	var data string
	var err error

	for read < WinCredZoneMaxLen {
		if read+MaxReadBytes < WinCredZoneMaxLen {
			length = MaxReadBytes
		} else {
			length = WinCredZoneMaxLen - read
		}
		cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
		readData, err = readResponseDecode(myport, MaxReadBytes, cmd, 100, false)
		//fmt.Println("Got read bytes: ", len(readData))
		//fmt.Println(readData)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		data += readData
		read += length
		offset += length
	}
	if len(data) < 8*WinCredBlockLen {
		return nil, errors.New("missing data when reading windcred zone, read bytes: " + fmt.Sprint(len(data)))
	}
	j := 0
	for i < 8*WinCredBlockLen {
		_, err := strconv.ParseUint(data[i:i+1], 10, 16)
		if err != nil {
			//it's a free wincred if there's a readable
			result = append(result, uint32(j))
		}
		i += WinCredBlockLen
		j++
	}

	return &gnsrpc.FreeWinCreds{Idx: result}, nil
}

func SetWinCred(in string) (*gnsrpc.WinCred, error) {
	var output gnsrpc.WinCred
	arr := strings.Split(in, "\t")
	if len(arr) >= 4 {
		index, err := strconv.ParseUint(strings.Trim(arr[0], WinCredNull), 10, 32)
		if err != nil {
			return nil, errors.New("error convert index")
		}
		output.Idx = uint32(index)
		output.Username = strings.Trim(arr[1], WinCredNull)
		output.Password = strings.Trim(arr[2], WinCredNull)
		output.Domain = strings.Trim(arr[3], WinCredNull)
	} else if len(arr) == 3 {
		index, err := strconv.ParseUint(strings.Trim(arr[0], WinCredNull), 10, 32)
		if err != nil {
			return nil, errors.New("error convert index")
		}
		output.Idx = uint32(index)
		output.Username = strings.Trim(arr[1], WinCredNull)
		output.Password = strings.Trim(arr[2], WinCredNull)

	} else if len(arr) == 2 {
		index, err := strconv.ParseUint(strings.Trim(arr[0], WinCredNull), 10, 32)
		if err != nil {
			return nil, errors.New("error convert index")
		}
		output.Idx = uint32(index)
		output.Username = strings.Trim(arr[1], WinCredNull)
	} else {
		return nil, errors.New("unable to decode wincred")
	}

	return &output, nil
}

func ReadWinCred(in *gnsrpc.WinCred) (*gnsrpc.WinCred, error) {
	zone := WinCredZone //change back to zone 2 later
	offset := 0
	var cmd string
	var readData string
	var err error
	if !IsAuthenticated {
		return nil, errors.New("badge is not ready")
	}
	offset = int(in.Idx) * WinCredBlockLen
	cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, WinCredBlockLen)
	readData, err = readResponseDecode(myport, MaxReadBytes, cmd, 100, false)
	//fmt.Println(readData)
	//fmt.Println(len(readData))
	//LogMsg(string.Format("Read Free wincred data size: {0}", data.Length));
	if len(readData) < WinCredBlockLen {
		return nil, errors.New("error reading wincred data")
	}

	wincred, err := SetWinCred(readData)
	if err != nil {
		return nil, errors.New("error decoding read wincred data")
	}

	return wincred, nil

}
func ReadWinCreds() (*gnsrpc.WinCreds, error) {
	wincreds := []*gnsrpc.WinCred{}
	//List<WinCred> result = new List<WinCred>();
	zone := WinCredZone //change back to zone 2 later
	offset := 0
	var length int
	read := 0
	var cmd string
	var readData string
	var data string
	var err error
	if !IsAuthenticated {
		return nil, errors.New("badge is not ready")
	}
	for read < WinCredZoneMaxLen {
		if read+MaxReadBytes < WinCredZoneMaxLen {
			length = MaxReadBytes
		} else {
			length = WinCredZoneMaxLen - read
		}
		cmd = fmt.Sprintf("READ_E %d %d %d", zone, offset, length)
		readData, err = readResponseDecode(myport, MaxReadBytes*2, cmd, 100, false)
		if err != nil {
			return nil, err
		}

		data += readData
		read += length
		offset += length
	}
	log.Println("Read Wincred Data size: ", len(data))
	//LogMsg(string.Format("Read Free wincred data size: {0}", data.Length));
	if len(data) < 8*WinCredBlockLen {
		return &gnsrpc.WinCreds{}, errors.New("error reading wincreds")
	}
	for i := 0; i < 8; i++ {
		startIdx := i * WinCredBlockLen
		substr := data[startIdx : startIdx+WinCredBlockLen]
		//fmt.Println(substr)
		wincred, err := SetWinCred(substr)
		if err != nil {
			continue
		}
		wincreds = append(wincreds, wincred)
	}
	return &gnsrpc.WinCreds{Wincreds: wincreds}, nil
}

func DeleteWinCred(in *gnsrpc.WinCred) error {
	zone := WinCredZone //change back to zone 2 later
	offset := 0
	var cmd string
	var data string
	var err error
	if !IsAuthenticated {
		return errors.New("badge is not ready")
	}
	offset = WinCredBlockLen * int(in.Idx)
	data = strings.Repeat(WinCredNull, WinCredBlockLen)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
	err = writeResponseDecode(myport, 100, cmd, 80)
	if err != nil {
		return err
	}
	return nil
}

func SerializeWinCred(in *gnsrpc.WinCred) string {
	result := fmt.Sprintf("%d", in.Idx) + "\t" + in.Username + "\t" + in.Password + "\t" + in.Domain + "\t"
	return result
}
func WriteWinCred(in *gnsrpc.WinCred) error {
	zone := WinCredZone //change back to zone 2 later
	offset := 0
	var cmd string
	var data string
	var err error
	if !IsAuthenticated {
		return errors.New("badge is not ready")
	}
	// Cannot write more than 8 slots for wincred
	if in.Idx < 0 || in.Idx > 7 {
		return errors.New("write index out of range. index has to be between 0 and 7")
	}

	offset = WinCredBlockLen * int(in.Idx)
	data = SerializeWinCred(in)
	//fmt.Println("Serialized wincred data: " + data)
	cmd = fmt.Sprintf("WRITE_E %d %d %d %s", zone, offset, len(data), data)
	err = writeResponseDecode(myport, 100, cmd, 80)
	if err != nil {
		return err
	}
	return nil
}

func UnlockCard() (string, error) {
	var result []byte
	var err error
	var unlocked bool
	var paired bool
	var loadedpub bool

	portInUsed.Lock()
	defer portInUsed.Unlock()

	unlocked = false
	loadedpub = false

	//** Start UNLOCK
	stSafePassword := "474e5354525553545345435552495459"
	result, err = readResponse(myport, 200, "UNLOCK "+stSafePassword, 200, false)

	unlockResp := string(result)
	fmt.Println("UNLOCK Response: " + unlockResp)
	if strings.Contains(unlockResp, "ST-Safe is locked - unlocking") || strings.Contains(unlockResp, "Unlock not needed") {
		unlocked = true
	} else {
		return "", errors.New("did not get expected unlock response")
	}

	log.Println("UNLOCK Response: " + string(result))
	log.Println("hex")
	log.Println(hex.EncodeToString(result))
	//might want to check for expected unlock response

	time.Sleep(300 * time.Millisecond)

	//** Start PAIR
	result, err = readResponse(myport, 200, "PAIR", 200, false)
	pairResp := string(result)
	fmt.Println("Pair Response: " + pairResp)
	if err != nil {
		log.Println("Paired error: " + err.Error())
		fmt.Println("Paired error: " + err.Error())
		return "", err
	}

	if strings.Contains(pairResp, "No Pair key on MCU - Pairing") || len(pairResp) == 0 {
		paired = true
	} else {
		paired = false
		return "", errors.New("did not get expected PAIR response")
	}
	log.Println("PAIR response: " + pairResp)
	log.Println(hex.EncodeToString(result))
	time.Sleep(300 * time.Millisecond)

	//** Start LOAD_PUBLIC
	public_key := gnscrypto.GeneratePubKeyFromPriv(pemLoc)
	fmt.Println("Public key: " + public_key)
	result, err = readResponse(myport, 200, "LOAD_PUBLIC "+public_key, 200, false)
	fmt.Println("Load_PUBLIC response: " + string(result))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		//for the correct load public response should be 0 bytes
		if len(result) > 0 {
			loadedpub = false
			return "", errors.New("did not get expected LOAD_PUBLIC response")
		} else {
			loadedpub = true
		}
	}
	//** Start GET_ID
	result, err = readResponse(myport, 200, "GET_ID", 200, false)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	var uuidStr string
	if len(result) >= 9 {
		result = result[0:9]
		uuidStr = hex.EncodeToString(result)
		loc, _ := time.LoadLocation("UTC")
		now := time.Now().In(loc)
		//fmt.Println("ZONE : ", loc, " Time : ", now) // UTC
		//log.Println("Successfully unlocked: " + uuidStr)
		//f.WriteString(fmt.Sprintln(now, ",", unlocked, ",", paired, ",", loadedpub, ",", public_key, ",", uuidStr))
		fmt.Println(now, ",", unlocked, ",", paired, ",", loadedpub, ",", public_key, ",", uuidStr)
	} else {
		//log.Println("Error unlocking card invalid UUID")
		return "", errors.New("invalid uuid string length")
	}

	return uuidStr, nil
}

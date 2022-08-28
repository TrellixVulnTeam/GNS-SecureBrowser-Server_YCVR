package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	gnscrypto "gnsdeviceserver/crypto"
	gnsserial "gnsdeviceserver/serial"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

var portName string
var IsConnected bool
var myport serial.Port
var pemLoc string
var portInUsed sync.Mutex
var MaxWriteBytes int = 100

var f *os.File

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Init() {
	system_os := runtime.GOOS
	switch system_os {
	case "darwin":
		pemLoc = "/Library/GNS/GNSPrivate.pem"
		log.Println("Using pem at: " + pemLoc)
	default:
		pemLoc = "GNSPrivate.pem"
	}
	IsConnected = false
}

func readResponse(port serial.Port, count int, command string, delay int) ([]byte, error) {
	var n int
	var err error
	sent := 0
	for sent < len([]rune(command)) {
		increment := min(MaxWriteBytes, len([]rune(command))-sent)
		n, err = port.Write([]byte(command[sent : sent+increment]))
		if err != nil {
			return nil, err
		}
		sent += increment
	}
	_, err = port.Write([]byte("\n"))
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
	buff := make([]byte, count)
	var result []byte
	for {
		// Reads up to 100 bytes
		n, err = port.Read(buff)
		if err != nil {
			return nil, err
		}
		if n == 0 {

			break
		}

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
	result, err = readResponse(myport, 200, "UNLOCK "+stSafePassword, 200)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		unlocked = true
	}
	log.Println("UNLOCK Response: " + string(result))
	log.Println("hex")
	log.Println(hex.EncodeToString(result))
	//might want to check for expected unlock response

	time.Sleep(300 * time.Millisecond)

	//** Start PAIR
	result, err = readResponse(myport, 200, "PAIR", 200)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		paired = true
	}
	log.Println("PAIR response: " + string(result))
	log.Println(hex.EncodeToString(result))
	time.Sleep(300 * time.Millisecond)

	//** Start LOAD_PUBLIC
	public_key := gnscrypto.GeneratePubKeyFromPriv(pemLoc)
	fmt.Println("Public key: " + public_key)
	result, err = readResponse(myport, 200, "LOAD_PUBLIC "+public_key, 200)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		//for the correct load public response should be 0 bytes
		if len(result) > 0 {
			loadedpub = false
		} else {
			loadedpub = true
		}
	}
	log.Println("LOAD_PUBLIC response: " + string(result))
	log.Println(hex.EncodeToString(result))

	//** Start GET_ID
	result, err = readResponse(myport, 200, "GET_ID", 200)
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
		f.WriteString(fmt.Sprintln(now, ",", unlocked, ",", paired, ",", loadedpub, ",", public_key, ",", uuidStr))
	} else {
		//log.Println("Error unlocking card invalid UUID")
		return "", errors.New("invalid uuid string length")
	}

	return uuidStr, nil
}

func FindPort() (string, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return "", err
	}
	if len(ports) == 0 {
		return "", nil
	}

	for _, port := range ports {
		if port.IsUSB {
			//fmt.Println(port.Name)
			if strings.Contains(port.Name, "usbmodem") || strings.Contains(port.Name, "COM") {
				fmt.Println("PID: " + port.PID)
				fmt.Println("Product: " + port.Product)
				fmt.Println("VID: " + port.VID)
				var lastChar = port.Name[len(port.Name)-1:]
				if lastChar == "3" {
					return port.Name, nil
				}
			}
		}
	}
	// we did not find a port
	return "", nil

}

func monitorPort() {
	var err error
	var uuid string
	var textWarning bool
	for {

		if !IsConnected {
			time.Sleep(2 * time.Second)
			portName, err = gnsserial.FindPort()
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
					fmt.Println("Connecting to port: " + portName)
					myport, err = serial.Open(portName, mode)

					if err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						IsConnected = true
						myport.SetReadTimeout(120 * time.Millisecond)
						textWarning = true
					}
				}
				uuid, err = UnlockCard()
				if err != nil {
					log.Println(err.Error())
				} else {
					if len(uuid) >= 9 {
						log.Println("Successfully unlocked: " + uuid)
					}
				}
			} else {
				IsConnected = false
				if textWarning {
					log.Println("No port found yet")
					textWarning = false
				}

				//fmt.Println(IsConnected)

			}
		} else {
			//if already connected  let's try to see if connection is alive
			_, error := myport.GetModemStatusBits()
			if error != nil {

				fmt.Println("Lost connection or error on Port. Will try to re-connect")
				IsConnected = false
			}
		}
	}
}

func main() {
	var err error
	log.Println("Starting GNS Provision")
	var firstWrite bool
	//
	if _, err := os.Stat("provision.log"); errors.Is(err, os.ErrNotExist) {
		firstWrite = true
	} else {
		firstWrite = false
	}
	f, err = os.OpenFile("provision.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if firstWrite {
		f.WriteString("TIME, UNLOCKED, PAIRED, LOADED_PUB, PUB_KEY, UUID\n")
	}
	defer f.Close()

	Init()
	done := make(chan bool)
	go monitorPort()
	<-done
}

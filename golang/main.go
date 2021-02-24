package main

import (
	"fmt"
	"strings"
	"time"

	//generate a manifest and compile via rsrc -manifest test.manifest -o rsrc.syso
	//then compile with go build -ldflags="-H windowsgui"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	//has no ability to list open com ports
	//but has a better serial communicatio library
	"github.com/tarm/serial"

	//we only use this for the com port list function
	//it can't set a timeout on serial connections so its garbage
	search "go.bug.st/serial"
)

var teensy *serial.Port
var errPort error
var comNum string
var windowMessage string = "Hello"

func main() {

	var inTE, outTE *walk.TextEdit
	var nitra string

	err := sync()
	if err != nil {
		windowMessage = "You Broke It"
	}

	MainWindow{
		Title:  "B.O.R.I.S.",
		Size:   Size{300, 200},
		Layout: VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true, Text: windowMessage},
				},
			},
			PushButton{
				Text: "UPDATE",
				OnClicked: func() {
					nitra = inTE.Text()
					teensy.Write([]byte(nitra + "\n"))
					outTE.SetText(windowMessage)
				},
			},
		},
	}.Run()

}

func sync() (err error) {

	comPortList, err := search.GetPortsList()
	if err != nil {
		return err
	}

	if len(comPortList) > 0 {
		for p := range comPortList {
			connectSuccessful := connect(comPortList[p])
			fmt.Println(comPortList[p])
			if connectSuccessful {
				break
			}
		}

	} else {
		windowMessage = "0 COM ports detected"
	}

	return err

}

// takes a com port ie COM9 and attempts to PING/ACK.
func connect(comPort string) (success bool) {

	//generate the serial port configuration
	c := &serial.Config{Name: comPort, Baud: 9600, ReadTimeout: time.Millisecond * 500}
	teensy, errPort = serial.OpenPort(c)
	if errPort != nil {
		teensy = nil
		return false
	}

	//attempt to start the handshake process with the teensy
	//all communications must end with newline so the teensy knows we've stopped transmitting
	c1 := make(chan error, 1)
	go func() {
		_, writeErr := teensy.Write([]byte("PING\n"))
		c1 <- writeErr
	}()

	select {
	case errPort = <-c1:
		buff := make([]byte, 32)
		incoming := ""
		for {
			n, err := teensy.Read(buff)
			if err != nil {
				windowMessage = "There was an error reading from teensy"
				return false
			} else if n == 0 {
				//end of file
				break
			}
			//append the characters in the buffer to the message we're trying to recieve
			incoming += string(buff[:n])

			//dont let another device spam us
			if len(incoming) > 5 {
				break
			}
			// If we receive a newline stop reading, this is a good thing
			if strings.Contains(string(buff[:n]), "\n") {
				break
			}
		}
		incoming = strings.TrimSpace(incoming)
		if incoming == "ACK" {
			windowMessage = "Connected on port: " + comPort
			return true
		}

	case <-time.After(100 * time.Millisecond):
		windowMessage = "No ports responded, connection failure"
		return false
	}

	return false
}

func writeNitra(nitra string) {
	if teensy != nil {
		_, err := teensy.Write([]byte(nitra + "\n"))
		if err != nil {
			windowMessage = "There was an error: " + err.Error()
		}
	} else {
		windowMessage = "update sent"
	}

}

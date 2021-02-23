package main

import (
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
			PushButton{
				Text: "Re-Connect",
				OnClicked: func() {
					//resync and reset the status message
					if sync() == nil {
						outTE.SetText(windowMessage)
					}

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
	for p := range comPortList {
		connectErr := connect(comPortList[p])
		if connectErr == nil {
			return connectErr
		}
	}

	return err

}

// takes a com port ie COM9 and attempts to PING/ACK.
func connect(comPort string) (err error) {

	//generate the serial port configuration
	c := &serial.Config{Name: "COM9", Baud: 9600, ReadTimeout: time.Second * 1}
	//attempt to connect to the arduino via serial
	teensy, errPort = serial.OpenPort(c)
	if err != nil {
		windowMessage = "cant open port"
		return errPort
	}

	//attempt to start the handshake process with the teensy
	//all communications must end with newline so the teensy knows we've stopped transmitting
	_, errPort = teensy.Write([]byte("PING\n"))
	if errPort != nil {
		return errPort
	}

	buff := make([]byte, 32)
	incoming := ""
	for {
		n, err := teensy.Read(buff)
		if err != nil {
			windowMessage = "cant read"
			//failure to read
			return err
		} else if n == 0 {
			//end of file
			break
		}
		//append the characters in the buffer to the message
		incoming += string(buff[:n])
		// If we receive a newline stop reading, this is a good thing
		if strings.Contains(string(buff[:n]), "\n") {
			break
		}
	}
	incoming = strings.TrimSpace(incoming)
	if incoming != "ACK" {
		windowMessage = incoming + " - cant open port"
		return err

	}
	//should be nil if we get to here
	return errPort
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

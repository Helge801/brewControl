package main

import (
	"errors"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

var reg = regexp.MustCompile(`\st=(\d+)`)

const interval = time.Second * 10

// RunLoop lets temp loop run as long as it's bookean value is true
// This var should probably have a mutex to protect against race conditions
// but it is only intended to be written too once to start and once to end
// so I'm just going to leave it for now.
var RunLoop = false

// StartMonitor begins monitoring temp from prob and adding DB entries with readings
func StartMonitor() {
	RunLoop = true
	AdLog("Starting recording")
	AdLog("opening gpio")
	e := rpio.Open()
	Fatal(e)
	AdLog("Established connection to GPIO pins")
	pin := rpio.Pin(17)
	pin.Output()
	AdLog("Getting probe path")
	path := getFilePath(pin)
	go runCheckLoop(pin, path)
}

func runCheckLoop(pin rpio.Pin, path string) {
	defer close(pin)
	for RunLoop {

		file, e := ioutil.ReadFile(path)
		if e != nil {
			badLoop(pin, "Cannot read from pobe")
			continue
		}
		match := reg.FindSubmatch(file)
		if len(match) > 1 {
			tempString := (string(match[1][:2]) + "." + string(match[1][2:]))
			temp, e := strconv.ParseFloat(tempString, 32)
			if e != nil {
				badLoop(pin, "Cannot Reat temp from probe output")
				continue
			}
			temp = (temp * 1.8) + 32.0
			temp = math.Floor(temp*10) / 10
			InsertEntry(float32(temp))
			SendEntry(float32(temp))
			if temp < 72.0 {
				pin.High()
			} else {
				pin.Low()
			}
		} else {
			badLoop(pin, "Cannot Read temp from probe output")
			continue
		}
		sleepInterval()
	}
}

func badLoop(pin rpio.Pin, message string) {
	pin.Low()
	NonFatal(errors.New(message))
	InsertEntry(0.0)
	sleepInterval()
}

func getFilePath(pin rpio.Pin) string {
	files, e := ioutil.ReadDir("/sys/bus/w1/devices/")
	if e != nil {
		close(pin)
		Fatal(e)
	}
	fileName := ""
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "28-") {
			fileName = f.Name()
			break
		}
	}
	if fileName == "" {
		close(pin)
		Fatal(errors.New("Could not establish connection with probe"))
	}
	return "/sys/bus/w1/devices/" + fileName + "/w1_slave"
}

func close(pin rpio.Pin) {
	pin.Low()
	rpio.Close()
}

func sleepInterval() {
	time.Sleep(interval)
}

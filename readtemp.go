package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

var reg = regexp.MustCompile(`\st=(\d+)`)

func main() {
	fmt.Println("opening gpio")
	e := rpio.Open()
	err(e)
	pin := rpio.Pin(17)
	pin.Output()
	defer close(pin)

	for true {

		files, e := ioutil.ReadDir("/sys/bus/w1/devices/")
		err(e)
		fileName := ""
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "28-") {
				fileName = f.Name()
				break
			}
		}

		file, e := ioutil.ReadFile("/sys/bus/w1/devices/" + fileName + "/w1_slave")
		err(e)
		match := reg.FindSubmatch(file)
		if len(match) > 1 {
			tempString := (string(match[1][:2]) + "." + string(match[1][2:]))
			temp, e := strconv.ParseFloat(tempString, 32)
			err(e)
			temp = (temp * 1.8) + 32.0
			temp = math.Floor(temp*10) / 10
			fmt.Println(temp)
			if temp > 73.0 {
				pin.High()
			} else {
				pin.Low()
			}
		}

		// pin.Toggle()

		time.Sleep(time.Second * 2)

	}

	// for x := 0; x < 20; x++ {
	//      pin.Toggle()
	//      time.Sleep(time.Second / 5)
	// }

}

func close(pin rpio.Pin) {
	pin.Low()
	rpio.Close()
}

func err(e error) {
	if e != nil {
		panic(fmt.Sprint(e))
	}
}

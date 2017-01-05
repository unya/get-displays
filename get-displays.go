package main

import (
	"fmt"
	"github.com/anoopengineer/edidparser/edid"
	"github.com/jochenvg/go-udev"
	"io/ioutil"
	"log"
	"strings"
)

func printEdidBytes(edid []byte) {
	fmt.Println("EDID dump")
	columnCounter := 0
	for i := 0; i < len(edid); i++ {
		fmt.Printf("%02X ", edid[i])
		if columnCounter == 15 {
			fmt.Println()
			columnCounter = 0
		} else {
			columnCounter++
		}
	}
}

func main() {
	u := udev.Udev{}
	e := u.NewEnumerate()
	e.AddMatchSubsystem("drm")
	e.AddMatchIsInitialized()
	devices, _ := e.Devices()
	for i := range devices {
		device := devices[i]
		if device.SysattrValue("status") == "connected" {
			devicePath := strings.Split(device.Syspath(), "/")
			deviceName := devicePath[len(devicePath)-1]
			// fmt.Println(device.Syspath())
			bs, err := ioutil.ReadFile(device.Syspath() + "/edid")
			if err != nil {
				log.Fatal("Unable to read EDID from ", device.Syspath(), err)
			}
			// printEdidBytes(bs)
			e, err := edid.NewEdid(bs)
			if err != nil {
				log.Fatal("Unable to parse EDID ", err)
			} else {
				fmt.Println(strings.Join([]string{deviceName, e.ManufacturerId, e.MonitorName, e.MonitorSerialNumber}, ":"))
			}
		}
	}
}

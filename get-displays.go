package main

import "fmt"
import "github.com/jochenvg/go-udev"

func main() {
	u := udev.Udev{}
	e := u.NewEnumerate()
	e.AddMatchSubsystem("drm")
	e.AddMatchIsInitialized()
	devices, _ := e.Devices()
	for i := range devices {
		device := devices[i]
		if device.SysattrValue("status") == "connected" {
			fmt.Println(device.Syspath())
		}
	}
}

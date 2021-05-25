package main

import (
	"fmt"
	"github.com/xupengyao/zoom-lib-golang"
)

func main() {
	zoom.APIKey = "BsFHCXL3TmiJkbCPsXIAGw"
	zoom.APISecret = "LHvcmdTGxMdvcNsv2ppm5gghtoGlMXmrvjmK"
	r,e := zoom.GetUserSettings(zoom.GetUserSettingsOpts{EmailOrID: "4quuSr4bSPqY7IiAQL9fGg"})
	fmt.Println(e)
	fmt.Println(r)
}

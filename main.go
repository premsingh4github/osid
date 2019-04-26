package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	fmt.Print("hello world")
	fmt.Print("this is test app")
	//v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	//fmt.Println(v)

	//out, _ := exec.Command("wmic bios get serialnumber").Output() // err ignored for brevity
	//for _, l := range strings.Split(string(out), "\n") {
	//	if strings.Contains(l, "IOPlatformSerialNumber") {
	//		s := strings.Split(l, " ")
	//		fmt.Printf("%s\n", s[len(s)-1])
	//	}
	//}
	//webView := webview.New(webview.Settings{
	//	Title:     "My test web view app",
	//	URL:       "http://google.com",
	//	Width:     1000,
	//	Height:    800,
	//	Resizable: true,
	//	Debug:     true,
	//	ExternalInvokeCallback: nil,
	//})
	//
	//webView.SetFullscreen(true)
	//webView.Run()

	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	//u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

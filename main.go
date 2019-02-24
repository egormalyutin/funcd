package main

import (
	"log"
	"os"

	. "github.com/gvalkov/golang-evdev"
)

func main() {
	if os.Getuid() != 0 {
		log.Print("Funcd must be started from root user.")
		log.Fatal("If you want to run it using \"sudo\", please, use \"su -c\", or audio controls will not work.")
	}

	var err error

	lastBrightness, err = getBrightness()
	if err != nil {
		log.Fatal(err)
	}

	users, err = getUsers()
	if err != nil {
		log.Fatal(err)
	}

	devices, err := findDevices()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	for _, device := range devices {
		go func(device *InputDevice) {
			for {
				func() {
					// Don't panic!
					defer func() {
						if r := recover(); r != nil {
							log.Print("Panic: ", r)
						}
					}()

					for {
						ev, err := device.ReadOne()
						if err != nil {
							done <- true
							log.Print(err)
						}

						binding, ok := bindings[int(ev.Code)]
						if ok {
							err = binding(ev)
							if err != nil {
								log.Print(err)
							}
						}
					}
				}()
			}
		}(device)
	}

	log.Printf("Funcd started")

	for _ = range devices {
		<-done
	}
}

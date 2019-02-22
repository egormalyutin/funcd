package main

import (
	"log"

	. "github.com/gvalkov/golang-evdev"
)

func main() {
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

	log.Printf("Funcd started, listening on %d devices", len(devices))

	for _ = range devices {
		<-done
	}
}

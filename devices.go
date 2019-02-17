package main

import (
	. "github.com/gvalkov/golang-evdev"
)

func findDevices() ([]*InputDevice, error) {
	devices, err := ListInputDevices()
	if err != nil {
		return []*InputDevice{}, err
	}

	results := []*InputDevice{}

	for _, device := range devices {
		for _, cp := range device.Capabilities {
			for _, key := range cp {
				if _, ok := bindings[key.Code]; ok {
					cont := false
					for _, result := range results {
						if device.Fn == result.Fn {
							cont = true
							break
						}
					}

					if !cont {
						results = append(results, device)
					}
				}
			}
		}
	}

	return results, nil
}

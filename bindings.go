package main

import (
	. "github.com/gvalkov/golang-evdev"
)

const brightnessDiff = 0.03

var lastBrightness float64 = 1

var bindings = map[int]func(*InputEvent) error{
	// BRIGHTNESS //
	KEY_SCREENLOCK: func(ev *InputEvent) error {
		kv := NewKeyEvent(ev)
		if kv.State == KeyUp {
			br, err := getBrightness()
			if err != nil {
				return err
			}

			if br == 0 {
				return setBrightness(lastBrightness)
			} else {
				lastBrightness = br
				return setBrightness(0)
			}
		}
		return nil
	},
	KEY_BRIGHTNESSUP: func(ev *InputEvent) error {
		kv := NewKeyEvent(ev)
		if kv.State != KeyUp {
			return incBrightness(brightnessDiff)
		}
		return nil
	},
	KEY_BRIGHTNESSDOWN: func(ev *InputEvent) error {
		kv := NewKeyEvent(ev)
		if kv.State != KeyUp {
			return decBrightness(brightnessDiff)
		}
		return nil
	},
}

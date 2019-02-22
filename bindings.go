package main

import (
	. "github.com/gvalkov/golang-evdev"
)

const brightnessDiff = 0.03

var lastBrightness float64 = 1

func screensaverBinding(ev *InputEvent) error {
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
}

func toggleTouchpadBinding(ev *InputEvent) error {
	kv := NewKeyEvent(ev)
	if kv.State == KeyUp {
		return toggleTouchpad()
	}
	return nil
}

var bindings = map[int]func(*InputEvent) error{
	// BRIGHTNESS //
	KEY_SCREENLOCK:  screensaverBinding,
	KEY_SCREENSAVER: screensaverBinding,
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

	// TOUCHPAD //
	KEY_TOUCHPAD_TOGGLE: toggleTouchpadBinding,
	KEY_F21:             toggleTouchpadBinding,

	// VOLUME //
	KEY_MUTE: func(ev *InputEvent) error {
		kv := NewKeyEvent(ev)
		if kv.State == KeyUp {
			return toggleVolume()
		}
		return nil
	},
	KEY_VOLUMEDOWN: func(ev *InputEvent) error {
		kv := NewKeyEvent(ev)
		if kv.State != KeyUp {
			return decVolume()
		}
		return nil
	},
	KEY_VOLUMEUP: func(ev *InputEvent) error {
		kv := NewKeyEvent(ev)
		if kv.State != KeyUp {
			return incVolume()
		}
		return nil
	},
}

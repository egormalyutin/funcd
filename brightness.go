package main

import (
	// "os"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"path"
	"strconv"
	"strings"
)

// Port of https://github.com/kevva/brightness/blob/master/lib/linux.js

const dir = "/sys/class/backlight"

func getBrightnessH(device string) (int, error) {
	b, err := ioutil.ReadFile(path.Join(dir, device, "brightness"))
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return 0, err
	}

	return n, nil
}

func getMaxBrightnessH(device string) (int, error) {
	b, err := ioutil.ReadFile(path.Join(dir, device, "max_brightness"))
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return 0, err
	}

	return n, nil
}

func setBrightnessH(device string, val int) error {
	return ioutil.WriteFile(path.Join(dir, device, "brightness"), []byte(fmt.Sprint(val)), 0644)
}

func getBacklightH() (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", errors.New("No backlight device found")
	}

	return files[0].Name(), nil
}

func getBrightness() (float64, error) {
	device, err := getBacklightH()
	if err != nil {
		return 0, err
	}

	max, err := getMaxBrightnessH(device)
	if err != nil {
		return 0, err
	}

	current, err := getBrightnessH(device)
	if err != nil {
		return 0, err
	}

	return float64(current) / float64(max), nil
}

func setBrightness(val float64) error {
	if val < 0 {
		val = 0
	} else if val > 1 {
		val = 1
	}

	device, err := getBacklightH()
	if err != nil {
		return err
	}

	max, err := getMaxBrightnessH(device)
	if err != nil {
		return err
	}

	brightness := int(math.Floor(val * float64(max)))

	return setBrightnessH(device, brightness)
}

func withBrightness(f func(float64) float64) error {
	b, err := getBrightness()
	if err != nil {
		return err
	}

	end := f(b)

	return setBrightness(end)
}

func decBrightness(val float64) error {
	return withBrightness(func(c float64) float64 { return c - val })
}

func incBrightness(val float64) error {
	return withBrightness(func(c float64) float64 { return c + val })
}

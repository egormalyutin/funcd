package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	idRegexp1    = regexp.MustCompile("TouchPad\\s*id\\=[0-9]{1,2}")
	idRegexp2    = regexp.MustCompile("[0-9]{1,2}")
	stateRegexp1 = regexp.MustCompile("(?m)Device Enabled.*$")
	stateRegexp2 = regexp.MustCompile("(?m)\\d+$")
)

func getTouchpadID() (int, error) {
	out, err := exec.Command("xinput", "list").Output()
	if err != nil {
		return 0, err
	}

	res1 := idRegexp1.Find(out)
	if res1 == nil {
		return 0, errors.New("Not found touchpad ID")
	}

	res2 := idRegexp2.Find(res1)
	if res2 == nil {
		return 0, errors.New("Not found touchpad ID")
	}

	return strconv.Atoi(strings.TrimSpace(string(res2)))
}

func getTouchpadState(id int) (bool, error) {
	out, err := exec.Command("xinput", "list-props", fmt.Sprint(id)).Output()
	if err != nil {
		return false, err
	}

	res1 := stateRegexp1.Find(out)
	if res1 == nil {
		return false, errors.New("Not found touchpad state")
	}

	res2 := stateRegexp2.Find(res1)
	if res2 == nil {
		return false, errors.New("Not found touchpad state")
	}

	return "1" == strings.TrimSpace(string(res2)), nil
}

func enableTouchpad(id int) error {
	return exec.Command("xinput", "enable", fmt.Sprint(id)).Run()
}

func disableTouchpad(id int) error {
	return exec.Command("xinput", "disable", fmt.Sprint(id)).Run()
}

func toggleTouchpad() error {
	id, err := getTouchpadID()
	if err != nil {
		return err
	}

	s, err := getTouchpadState(id)
	if err != nil {
		return err
	}

	if s {
		return disableTouchpad(id)
	} else {
		return enableTouchpad(id)
	}
}

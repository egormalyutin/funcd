package main

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	idRegexp1 = regexp.MustCompile("TouchPad\\s*id\\=[0-9]{1,2}")
	idRegexp2 = regexp.MustCompile("[0-9]{1,2}")
)

func getID() (int, error) {
	out, err := exec.Command("xinput", "list").Output()
	if err != nil {
		return 0, err
	}

	res1 := idRegexp1.Find(out)
	if res1 == nil {
		return 0, errors.New("Not found touchpad")
	}

	res2 := idRegexp2.Find(res1)
	if res2 == nil {
		return 0, errors.New("Not found touchpad")
	}

	return strconv.Atoi(strings.TrimSpace(string(res2)))
}

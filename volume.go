package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const volumeDiff = "3"

func runUser(command string, args ...string) error {
	if len(users) == 0 {
		return errors.New("Not found any usual user for changing volume")
	}

	user := users[0]
	cmd := exec.Command(command, args...)
	cmd.Env = append(os.Environ(), fmt.Sprint("XDG_RUNTIME_DIR=/run/user/", user.Uid))
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: user.Uid, Gid: user.Gid}
	return cmd.Run()
}

func decVolume() error {
	return runUser("amixer", "set", "Master", volumeDiff+"%-")
}

func incVolume() error {
	return runUser("amixer", "set", "Master", volumeDiff+"%+")
}

func toggleVolume() error {
	return runUser("amixer", "set", "Master", "toggle")
}

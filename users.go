package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.socketloop.com/tutorials/golang-get-all-local-users-and-print-out-their-home-directory-description-and-group-id

type User struct {
	Uid uint32
	Gid uint32
}

var users = []User{}

func getUsers() ([]User, error) {
	result := []User{}

	file, err := os.Open("/etc/passwd")

	if err != nil {
		return result, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if equal := strings.Index(line, "#"); equal < 0 {
			lineSlice := strings.FieldsFunc(line, func(divide rune) bool {
				return divide == ':'
			})

			if len(lineSlice) >= 3 {
				uid, err := strconv.Atoi(lineSlice[2])
				if err != nil {
					return result, err
				}

				if uid >= 1000 && uid != 65534 {
					gid, err := strconv.Atoi(lineSlice[3])
					if err != nil {
						return result, err
					}

					result = append(result, User{Uid: uint32(uid), Gid: uint32(gid)})
				}
			}

		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}

	}

	return result, nil
}

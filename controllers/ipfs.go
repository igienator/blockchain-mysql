package controllers

import (
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func HashIpfs() error {
	sh := shell.NewShell("localhost:5001")
	var rowHash string
	var allHash []string

	users, err := ReturnUsers()
	if err != nil {
		return err
	}
	for _, userString := range users {
		nameHash, err := sh.Add(strings.NewReader(userString.Name))
		if err != nil {
			return err
		}
		surnameHash, err := sh.Add(strings.NewReader(userString.Surname))
		if err != nil {
			return err
		}
		emailHash, err := sh.Add(strings.NewReader(userString.Email))
		if err != nil {
			return err
		}
		descHash, err := sh.Add(strings.NewReader(userString.Description))
		if err != nil {
			return err
		}

		rowHash = strings.Join([]string{nameHash, surnameHash, emailHash, descHash}, ",")
		allHash = append(allHash, rowHash)
	}
	allRows := strings.Join(allHash, " | ")
	cid, err := sh.Add(strings.NewReader(allRows))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)
	if err = sh.Publish("/ipfs/QmUsersHash", cid); err != nil {
		return err
	}

	return nil
}

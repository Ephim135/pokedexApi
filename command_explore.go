package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("to many arguments")
	}
	if len(args) < 1 {
		return errors.New("to few arguments")
	}
	fmt.Println(args[0])
	return nil
}

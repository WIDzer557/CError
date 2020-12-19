package catcher

import (
	"fmt"
	"os"
)

var Stdout = os.Stdout

func Print(err interface{}) {
	if err != nil {
		fmt.Fprintln(Stdout, err)
	}
}

func Fatal(err interface{}) {
	if err != nil {
		fmt.Fprintln(Stdout, err)
		os.Exit(1)
	}
}
package main

import (
	"fmt"
	"github.com/k-takata/go-iscygpty"
	"os"
)

func main() {
	type pair struct {
		f    *os.File
		name string
	}
	var a = []pair{
		{os.Stdin, "Stdin"},
		{os.Stdout, "Stdout"},
		{os.Stderr, "Stderr"},
	}
	for _, v := range a {
		if iscygpty.IsCygwinPty(v.f.Fd()) {
			fmt.Println(v.name, "is Cygwin/MSYS pty.",
				"(" + iscygpty.GetPipeName(v.f.Fd()) + ")")
		} else {
			fmt.Println(v.name, "is Not Cygwin/MSYS pty.")
		}
	}

	if iscygpty.IsCygwinPtyUsed() {
		fmt.Println("Cygwin/MSYS pty is used.")
	} else {
		fmt.Println("Cygwin/MSYS pty is not used.")
	}
}

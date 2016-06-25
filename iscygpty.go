package iscygpty

import (
	"os"
)

// Return true if Cygwin/MSYS pty is used.
func IsCygwinPtyUsed() bool {
	ret := false
	for _, v := range [](*os.File){os.Stdin, os.Stdout, os.Stderr} {
		ret = ret || IsCygwinPty(v.Fd())
	}
	return ret
}

// +build !windows

package iscygpty

func IsCygwinPty(fd uintptr) bool {
	return false
}

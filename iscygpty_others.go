// +build !windows

package iscygpty

// IsCygwinPty returns true if the file descriptor is Cygwin/MSYS pty.
func IsCygwinPty(fd uintptr) bool {
	return false
}

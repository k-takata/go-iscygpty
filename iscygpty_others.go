// +build !windows

package iscygpty

// GetPipeName returns the name of the file descriptor if it is a named pipe
// on Windows. Otherwise it returns an empty string.
func GetPipeName(fd uintptr) string {
	return ""
}

// IsCygwinPty returns true if the file descriptor is a Cygwin/MSYS pty.
func IsCygwinPty(fd uintptr) bool {
	return false
}

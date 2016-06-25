// +build windows

package iscygpty

import (
	"regexp"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var procGetFileInformationByHandleEx = kernel32.NewProc("GetFileInformationByHandleEx")
var procGetFileType = kernel32.NewProc("GetFileType")

const fileNameInfo uintptr = 2
const fileTypePipe = 3

// IsCygwinPty returns true if the file descriptor is Cygwin/MSYS pty.
// Only works on Vista or later. (Always returns false on XP or earlier.)
func IsCygwinPty(fd uintptr) bool {
	// Check if GetFileInformationByHandleEx is available.
	proc := procGetFileInformationByHandleEx
	if proc == nil {
		return false
	}
	err := proc.Find()
	if err != nil {
		procGetFileInformationByHandleEx = nil
		return false
	}

	// Cygwin/msys's pty is a pipe.
	ft, _, e := syscall.Syscall(procGetFileType.Addr(), 1, fd, 0, 0)
	if ft != fileTypePipe || e != 0 {
		return false
	}

	var buf [2 + syscall.MAX_PATH]uint16

	r, _, e := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(),
		4, fd, fileNameInfo, uintptr(unsafe.Pointer(&buf)),
		uintptr(len(buf)*2), 0, 0)
	if r == 0 || e != 0 {
		return false
	}

	l := *(*uint32)(unsafe.Pointer(&buf))
	s := string(utf16.Decode(buf[2 : 2+l/2]))

	// Check the name of the pipe.
	matched, _ := regexp.MatchString(`\\(?:cygwin|msys)-[0-9a-f]{16}-pty[0-9]+-(?:from|to)-master`, s)
	if matched {
		return true
	}
	return false
}

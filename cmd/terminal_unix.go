package cmd

import (
	"syscall"
	"unsafe"
)

var s struct {
	rows uint16
	cols uint16
	posX uint16
	posY uint16
}

func GetTerminalSize() (int, int) {
	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&s)),
	)
	return int(s.cols), int(s.rows)
}

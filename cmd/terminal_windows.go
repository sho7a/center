// +build !linux windows !darwin

package cmd

import (
	"syscall"
	"unsafe"
)

type (
	short     int16
	word      uint16
	smallRect struct {
		Left   short
		Top    short
		Right  short
		Bottom short
	}
	coord struct {
		X short
		Y short
	}
	terminalScreenBufferInfo struct {
		Size              coord
		CursorPosition    coord
		Attributes        word
		Window            smallRect
		MaximumWindowSize coord
	}
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var GetTerminalScreenBufferInfoProc = kernel32.NewProc("GetConsoleScreenBufferInfo")

func GetTerminalSize() (int, int) {
	handle := GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	info := GetTerminalScreenBufferInfo(handle)
	return int(info.Window.Right - info.Window.Left + 1), int(info.Window.Bottom - info.Window.Top + 1)
}

func GetStdHandle(stdHandle int) uintptr {
	handle, _ := syscall.GetStdHandle(stdHandle)
	return uintptr(handle)
}

func GetTerminalScreenBufferInfo(handle uintptr) *terminalScreenBufferInfo {
	var info terminalScreenBufferInfo
	GetTerminalScreenBufferInfoProc.Call(handle, uintptr(unsafe.Pointer(&info)), 0)
	return &info
}

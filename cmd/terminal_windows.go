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

var k32 = syscall.NewLazyDLL("kernel32.dll")
var GetTerminalScreenBufferInfoProc = k32.NewProc("GetConsoleScreenBufferInfo")

func GetTerminalSize() (int, int) {
	h := GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	i, _ := GetTerminalScreenBufferInfo(h)
	return int(i.Window.Right - i.Window.Left + 1), int(i.Window.Bottom - i.Window.Top + 1)
}

func GetStdHandle(stdH int) uintptr {
	h, _ := syscall.GetStdHandle(stdH)
	return uintptr(h)
}

func GetTerminalScreenBufferInfo(h uintptr) *terminalScreenBufferInfo {
	var i terminalScreenBufferInfo
	GetTerminalScreenBufferInfoProc.Call(h, uintptr(unsafe.Pointer(&i)), 0)
	return &i
}

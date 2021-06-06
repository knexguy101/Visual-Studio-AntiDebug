package main

import (
	"fmt"
	"github.com/nvsoft/win"
	"syscall"
	"time"
)

func DisableAttachDebug() error {
	mainDll, err := syscall.LoadLibrary("Dll1.dll")
	if err != nil {
		panic(err)
	}
	defer syscall.FreeLibrary(mainDll)
	eraseHandle, err := syscall.GetProcAddress(mainDll, "ErasePEHeaderFromMemory")
	if err != nil {
		panic(err)
	}

	hInstance := win.GetModuleHandle(nil)
	for range time.Tick(1 * time.Second) {
		_, _, err := syscall.Syscall(eraseHandle, 1, uintptr(hInstance), 0, 0)
		if err.Error() != "The operation completed successfully." {
			fmt.Println(err)
		}
	}
	return nil
}
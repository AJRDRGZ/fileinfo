//go:build windows
// +build windows

package fileinfo

import (
	"fmt"

	"golang.org/x/sys/windows"
)

// IsHidden returns if the windows file is hidden
func IsHidden(filename string) bool {
	pointer, err := windows.UTF16PtrFromString(filename)
	if err != nil {
		panic(fmt.Sprint(filename, err))
	}
	attributes, err := windows.GetFileAttributes(pointer)
	if err != nil {
		panic(fmt.Sprint(filename, err))
	}
	return attributes&windows.FILE_ATTRIBUTE_HIDDEN != 0
}

// GetUserAndGroup returns always empty because the windows os hasn't
// this information
func GetUserAndGroup(infoSys any) (userName, groupName string) {
	return
}

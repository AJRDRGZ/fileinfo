//go:build unix
// +build unix

package fileinfo

import (
	"fmt"
	"os/user"
	"strings"
	"syscall"
)

// IsHidden returns if the unix file is hidden
func IsHidden(filename string) bool {
	return strings.HasPrefix(filename, ".")
}

// GetUserAndGroup returns the user name and group name of an unix file
func GetUserAndGroup(infoSys any) (userName, groupName string) {
	stat, ok := infoSys.(*syscall.Stat_t)
	if !ok {
		return
	}

	if u, err := user.LookupId(fmt.Sprintf("%d", stat.Uid)); err == nil {
		userName = u.Username
	}

	if g, err := user.LookupGroupId(fmt.Sprintf("%d", stat.Gid)); err == nil {
		groupName = g.Name
	}

	return
}

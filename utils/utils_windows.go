package utils

import (
	"io/fs"
	"syscall"
	"time"
)

var LogDirs = [2]string{
	`\Documents\IINACT`,
	`\AppData\Roaming\Advanced Combat Tracker\FFXIVLogs`,
}

func GetCreationTime(fileInfo fs.FileInfo) (time.Time, error) {
	d := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	if d == nil {
		return time.Now(), &FFLogError{}
	}
	ctime := time.Unix(0, int64(d.CreationTime.Nanoseconds()))
	return ctime, nil
}

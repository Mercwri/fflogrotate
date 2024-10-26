package rotate

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/Mercwri/fflogrotate/utils"
)

func Rotate(short_term int, long_term int) {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	for _, logLoc := range utils.LogDirs {
		logDirPath := fmt.Sprintf("%s%s", user.HomeDir, logLoc)
		log.Printf("Scanning Dir: %s", logDirPath)
		logDir := NewLogDir(logDirPath, (time.Hour * time.Duration(short_term)), (time.Hour * time.Duration(long_term)))
		for _, logF := range logDir.LogFiles {
			log.Printf("Found: %s is %s old", logF.FileName, logF.LogAge)
			if logF.Archive {
				if logF.LogAge > logDir.LongTerm {
					overDue := logDir.LongTerm - logF.LogAge
					log.Printf("Archive %s is %s older than the Longterm Rotation Period", logF.FileName, overDue)
					log.Printf("Deleting: %s", logF.FileName)
					err = os.Remove(logF.FilePath)
					if err != nil {
						log.Panic(err)
					}
				}
			}
			if !logF.Archive {
				if logF.LogAge > logDir.ShortTerm {
					log.Printf("Rotating: %s", logF.FileName)
					err = logF.ArchiveFile()
					if err != nil {
						log.Panic(err)
					}
					err = os.Remove(logF.FilePath)
					if err != nil {
						log.Panic(err)
					}
				}
			}
		}
	}
}

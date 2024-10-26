package rotate

import (
	"log"
	"os"
	"time"
)

type FFLogFolder struct {
	Path      string
	ShortTerm time.Duration
	LongTerm  time.Duration
	LogFiles  []FFLogFile
}

func NewLogDir(path string, termshort time.Duration, termlong time.Duration) FFLogFolder {
	var files []FFLogFile
	logFolder, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	logFiles, err := logFolder.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, logF := range logFiles {
		valid, err := ValidateFilename(logF.Name())
		if err != nil {
			log.Println(err)
		}
		if valid {
			newLogFile, err := NewLogFile(path, logF)
			if err != nil {
				log.Println(err)
				continue
			}
			files = append(files, newLogFile)
		}
	}
	return FFLogFolder{
		path,
		termshort,
		termlong,
		files,
	}
}

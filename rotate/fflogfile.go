package rotate

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/Mercwri/fflogrotate/utils"
)

type FFLogFile struct {
	FilePath string
	LogAge   time.Duration
	Archive  bool
	FileName string
}

func NewLogFile(path string, file fs.DirEntry) (FFLogFile, error) {
	newFile := FFLogFile{}
	fullpath := fmt.Sprintf("%s\\%s", path, file.Name())
	fileHandler, err := os.Open(fullpath)
	if err != nil {
		return newFile, err
	}
	fileInfo, err := fileHandler.Stat()
	if err != nil {
		return newFile, err
	}
	defer fileHandler.Close()
	newFile.FileName = file.Name()
	newFile.Archive = false
	newFile.FilePath = fullpath
	ctime, err := utils.GetCreationTime(fileInfo)
	if err != nil {
		log.Panic(err)
	}
	newFile.LogAge = time.Since(ctime)
	err = newFile.IsArchive()
	if err != nil {
		log.Panic(err)
	}
	return newFile, nil
}

func ValidateFilename(filename string) (bool, error) {
	regx := regexp.MustCompile(`Network_\d{5}_\d{8}(\.\w{3}|\.\w{3}\.\w{3})`)
	matchOut := regx.Match([]byte(filename))
	return matchOut, nil
}

func (f *FFLogFile) IsArchive() error {
	bytes, err := os.ReadFile(f.FilePath)
	if err != nil {
		return err
	}
	filetype := http.DetectContentType(bytes)
	if filetype == "application/zip" {
		f.Archive = true
	}
	return nil
}

func (f *FFLogFile) ArchiveFile() error {
	archive, err := os.Create(fmt.Sprintf("%s.zip", f.FilePath))
	if err != nil {
		log.Panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	logFile, err := os.Open(f.FilePath)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	zw, err := zipWriter.Create(f.FileName)
	if err != nil {
		log.Panic(err)
	}
	if _, err := io.Copy(zw, logFile); err != nil {
		log.Panic(err)
	}
	defer zipWriter.Close()
	return err
}

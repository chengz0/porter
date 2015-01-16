package util

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

// write back to file
func WriteConfig(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	file.Write(data)
	return nil
}

// get file info
func GetFileInfo(path string) ([]byte, error) {
	fileinfo, err := os.Stat(path)
	if os.IsExist(err) {
		return nil, err
	}
	ret := FileInfo{
		Name:    fileinfo.Name(),
		Size:    fileinfo.Size(),
		Mode:    fileinfo.Mode(),
		ModTime: fileinfo.ModTime(),
		IsDir:   fileinfo.IsDir(),
	}
	body, _ := json.Marshal(ret)
	return body, nil
}

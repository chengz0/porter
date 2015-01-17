package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

// file exist
func FileExist(path string) bool {
	file, err := os.Stat(path)
	if err != nil || file.IsDir() {
		return false
	}
	return true
}

// read bytes from file
func ReadBytesFromFile(file string) ([]byte, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return bs, nil
}

// read string from file
func ReadStringFromFile(file string) (string, error) {
	str, err := ReadBytesFromFile(file)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

// read non empty string from file
func ReadNoStringFromFile(file string) (string, error) {
	str, err := ReadStringFromFile(file)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(str, "\r\t\n "), nil
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

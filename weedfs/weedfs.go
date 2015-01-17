package weedfs

import (
	"bytes"
	"github.com/deepglint/muses/util/http"
	"github.com/deepglint/weed-fs/go/operation"
	"io/ioutil"
	"log"
	"os"
)

type WeedConfig struct {
	WeedServer string
	TTL        int
	MaxMB      int
}

func (weed *WeedConfig) SubmitFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Can not open file: %s\n", path)
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Can not read file: %s\n", file.Name())
		return "", err
	}
	reader := bytes.NewReader(content)

	// write to weedfs
	part := operation.FilePart{
		Reader: reader,
	}
	parts := make([]operation.FilePart, 1)
	parts[0] = part
	results, err := operation.SubmitFiles(weed.WeedServer, parts, "", "", weed.TTL, weed.MaxMB)
	if err != nil {
		log.Printf("Can not submit file %s to weed.\n", path)
		return "", err
	}
	return results[0].Fid, nil
}

func (weed *WeedConfig) GetFile(fid string) ([]byte, error) {
	url, err := operation.LookupFileId(weed.WeedServer, fid)
	if err != nil {
		log.Printf("Can not find file: %s\n", fid)
		return nil, err
	}
	body, err := http.HTTPGet(url)
	if err != nil {
		log.Printf("Can not download file: %s", fid)
		return nil, err
	}
	return body, nil
}

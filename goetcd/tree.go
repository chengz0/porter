package goetcd

import (
	"encoding/json"
	"fmt"
	"github.com/chengz0/porter/util"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileModel struct {
	Name string
	Path string
	Hash string
	Fid  string
}

// init dir on etcd for module xxx
func InitDir4Module(client *etcd.Client, module string) error {
	_, err := client.CreateDir("/deepglint/"+module, 0)
	return err
}

// list modules in etcd
func ListModules(client *etcd.Client) error {
	resp, err := client.Get("/deepglint", false, false)
	fmt.Println("Modules: ")
	for _, node := range resp.Node.Nodes {
		fmt.Printf("	%s\n", node.Key)
	}
	return err
}

func InitDirTree4Module(client *etcd.Client, module, path string) {
	files := make(map[string]*FileModel)
	err := filepath.Walk(path, func(path string, curf os.FileInfo, err error) error {
		if curf == nil {
			return err
		}
		if curf.IsDir() {
			return nil
		}

		// insert to etcd
		curfile := new(FileModel)
		curfile.Name = curf.Name()
		curfile.Path = path
		//
		fileinfo, _ := util.GetFileInfo(path)
		hash := util.MD5Hash(fileinfo)
		curfile.Hash = hash
		// curfile.Fid =
		//
		files[hash] = curfile

		value, _ := json.Marshal(curfile)
		client.Create("/deepglint/"+module+"/1.0.1-dev/"+path, string(value), 50)
		return nil
	})
	wc, _ := json.Marshal(files)
	util.WriteConfig(".tmp", wc)
	if err != nil {
		log.Printf("filepath.Walk() returned %v\n", err)
	}
}

// add version
func AddIterVersion(client *etcd.Client, module, version, role string) error {
	key := fmt.Sprintf("/deepglint/%s/%s/%s", module, version, role)
	value := fmt.Sprintf("192.168.2.103:5000/%s:%s-%s", module, version, role)
	_, err := client.Create(key, value, 0)
	return err
}

// pass to next phase
func Pass2NextPhase(client *etcd.Client, module, from, to string) error {
	fromkey := fmt.Sprintf("/deepglint/%s", module)
	fromnode, err := client.Get(fromkey, true, false)
	if err != nil {
		return err
	}

	tmpkey := fromnode.Node.Nodes[fromnode.Node.Nodes.Len()-1].Key
	tmpnode, err := client.Get(tmpkey+"/"+from, false, false)
	if err != nil {
		return err
	}
	tokey := tmpkey + "/" + to
	tovalue := strings.TrimRight(tmpnode.Node.Value, from) + to
	_, err = client.Create(tokey, tovalue, 0)
	return err
}

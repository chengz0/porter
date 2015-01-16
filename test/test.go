package main

import (
	// "../goetcd"
	// "github.com/coreos/go-etcd/etcd"
	"encoding/json"
	"log"
	"os"
)

func main() {
	// etcd.InitDirTree("./libra")
	// 	client := etcd.NewClient([]string{"http://192.168.2.100:4001"})

	// 	s := client.GetCluster()
	// 	log.Println(s)

	// 	goetcd.InitDirTree(client, "libra")
	var file os.FileInfo

	file, _ = os.Stat("cmd_t.go")
	log.Println(file)

	body, _ := json.Marshal(file)
	log.Println(string(body))
}

package global

import (
	"github.com/astaxie/beego/config"
	"github.com/chengz0/porter/goetcd"
	"github.com/chengz0/porter/util"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"os"
)

type GlobalConfig struct {
	Auther    string
	Email     string
	Version   string
	Client    *etcd.Client
	WeedAddr  string
	WeedTTL   int
	WeedMaxMB int
}

var INIConfig config.ConfigContainer
var Config GlobalConfig

func InitEnv(configpath string) {
	if !util.FileExist(configpath) {
		log.Fatalf("Wrong File: %s\n", configpath)
	}
	var err error
	Config.Version, err = util.ReadNoStringFromFile("VERSION")
	if err != nil {
		log.Fatalln("Wrong VERSION file.")
	}

	INIConfig, err = config.NewConfig("ini", configpath)
	if err != nil {
		log.Fatalf("configuration file[%s] cannot parse. ", configpath)
		os.Exit(0)
	}
	Config.Auther = INIConfig.String("auther")
	Config.Email = INIConfig.String("email")

	etcdaddr := INIConfig.String("etcd::addr")
	Config.Client = goetcd.NewEtcdClient(etcdaddr)
	Config.WeedAddr = INIConfig.String("weed::addr")
	Config.WeedTTL, _ = INIConfig.Int("weed::ttl")
	Config.WeedMaxMB, _ = INIConfig.Int("weed::maxmb")

	log.Println(Config)
}

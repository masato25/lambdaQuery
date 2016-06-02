package g

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/toolkits/file"
)

type HttpConfig struct {
	Listen string `json:"listen"`
}

type GraphConfig struct {
	ConnTimeout int32             `json:"connTimeout"`
	CallTimeout int32             `json:"callTimeout"`
	MaxConns    int32             `json:"maxConns"`
	MaxIdle     int32             `json:"maxIdle"`
	Replicas    int32             `json:"replicas"`
	Cluster     map[string]string `json:"cluster"`
}

type GraphDBConfig struct {
	Addr string `json:"addr"`
	Idle int    `json:"idle"`
	Max  int    `json:"max"`
}

type GlobalConfig struct {
	Http    *HttpConfig    `json:"http"`
	Graph   *GraphConfig   `json:"graph"`
	GraphDB *GraphDBConfig `json:"graphdb"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("config file %s is nonexistent", cfg)
	}
	configContent, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		return fmt.Errorf("read config file %s fail %s", cfg, err)
	}
	var c GlobalConfig
	err = json.Unmarshal(configContent, &c)
	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("read config file:", cfg, "successfully")
	return nil
}
